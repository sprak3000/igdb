package igdb

import (
	"strconv"
	"strings"
)

// PulseGroup type
type PulseGroup struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       URL    `json:"url"`
	CreatedAt int    `json:"created_at"` //unix epoch
	UpdatedAt int    `json:"updated_at"` //unix epoch
	Tags      []Tag  `json:"tags"`
	Pulses    []int  `json:"pulses"`
	Game      []int  `json:"game"`
}

// GetPulseGroup gets IGDB information for a pulse group identified by its unique IGDB ID.
func (c *Client) GetPulseGroup(id int, opts ...OptionFunc) (*PulseGroup, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "pulse_groups/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var p []PulseGroup

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPulseGroups gets IGDB information for a list of pulse groups identified by their
// unique IGDB IDs.
func (c *Client) GetPulseGroups(ids []int, opts ...OptionFunc) ([]*PulseGroup, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToString(ids)
	url := c.rootURL + "pulse_groups/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var p []*PulseGroup

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPulseGroups searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPulseGroups(qry string, opts ...OptionFunc) ([]*PulseGroup, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "pulse_groups/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var p []*PulseGroup

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
