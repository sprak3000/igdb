package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct TestDummy -add-tags json -w

type TestDummy struct {
	ID              int           `json:"int"`
	BoolValue       bool          `json:"bool_value"`
	CreatedAt       int           `json:"created_at"`
	EnumTest        TestDummyEnum `json:"enum_test"`
	FloatValue      float64       `json:"float_value"`
	Game            int           `json:"game"`
	IntegerArray    []int         `json:"integer_array"`
	IntegerValue    int           `json:"integer_value"`
	Name            string        `json:"name"`
	NewIntegerValue int           `json:"new_integer_value"`
	Private         bool          `json:"private"`
	Slug            string        `json:"slug"`
	StringArray     []string      `json:"string_array"`
	TestDummies     []int         `json:"test_dummies"`
	TestDummy       int           `json:"test_dummy"`
	UpdatedAt       int           `json:"updated_at"`
	URL             string        `json:"url"`
	User            int           `json:"user"`
}

//go:generate stringer -type=TestDummyEnum

type TestDummyEnum int

const (
	TestDummyEnum1 TestDummyEnum = iota + 1
	TestDummyEnum2
)

// TestDummyService handles all the API calls for the IGDB TestDummy endpoint.
type TestDummyService service

// Get returns a single TestDummy identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any TestDummies, an error is returned.
func (ts *TestDummyService) Get(id int, opts ...Option) (*TestDummy, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var dum []*TestDummy

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ts.client.get(ts.end, &dum, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get TestDummy with ID %v", id)
	}

	return dum[0], nil
}

// List returns a list of TestDummies identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a TestDummy is ignored. If none of the IDs
// match a TestDummy, an error is returned.
func (ts *TestDummyService) List(ids []int, opts ...Option) ([]*TestDummy, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var dum []*TestDummy

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ts.client.get(ts.end, &dum, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get TestDummies with IDs %v", ids)
	}

	return dum, nil
}

// Index returns an index of TestDummies based solely on the provided functional
// options used to sort, filter, and paginate the results. If no TestDummies can
// be found using the provided options, an error is returned.
func (ts *TestDummyService) Index(opts ...Option) ([]*TestDummy, error) {
	var dum []*TestDummy

	err := ts.client.get(ts.end, &dum, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of TestDummies")
	}

	return dum, nil
}

// Count returns the number of TestDummies available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which TestDummies to count.
func (ts *TestDummyService) Count(opts ...Option) (int, error) {
	ct, err := ts.client.getCount(ts.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count TestDummies")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB TestDummy object.
func (ts *TestDummyService) Fields() ([]string, error) {
	f, err := ts.client.getFields(ts.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get TestDummy fields")
	}

	return f, nil
}