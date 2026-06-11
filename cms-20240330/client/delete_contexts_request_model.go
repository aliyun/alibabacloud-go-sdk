// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextIds(v string) *DeleteContextsRequest
	GetContextIds() *string
	SetFilter(v string) *DeleteContextsRequest
	GetFilter() *string
}

type DeleteContextsRequest struct {
	// A comma-separated list of context IDs.
	//
	// example:
	//
	// mem_long_01,mem_long_02
	ContextIds *string `json:"contextIds,omitempty" xml:"contextIds,omitempty"`
	// The filter condition, specified as a JSON string in the query. The syntax is the same as the `filter` parameter of the `SearchContext` operation.
	//
	// example:
	//
	// {"userId":"u-10001"}
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s DeleteContextsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextsRequest) GetContextIds() *string {
	return s.ContextIds
}

func (s *DeleteContextsRequest) GetFilter() *string {
	return s.Filter
}

func (s *DeleteContextsRequest) SetContextIds(v string) *DeleteContextsRequest {
	s.ContextIds = &v
	return s
}

func (s *DeleteContextsRequest) SetFilter(v string) *DeleteContextsRequest {
	s.Filter = &v
	return s
}

func (s *DeleteContextsRequest) Validate() error {
	return dara.Validate(s)
}
