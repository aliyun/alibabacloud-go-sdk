// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListProductsRequest
	GetFilter() *string
}

type ListProductsRequest struct {
	// example:
	//
	// ECS
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListProductsRequest) SetFilter(v string) *ListProductsRequest {
	s.Filter = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
