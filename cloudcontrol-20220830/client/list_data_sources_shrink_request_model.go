// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeName(v string) *ListDataSourcesShrinkRequest
	GetAttributeName() *string
	SetFilterShrink(v string) *ListDataSourcesShrinkRequest
	GetFilterShrink() *string
}

type ListDataSourcesShrinkRequest struct {
	// The name of the attribute. Only `RegionId` is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// RegionId
	AttributeName *string `json:"attributeName,omitempty" xml:"attributeName,omitempty"`
	// The filter condition. The value must be a JSON string in the {"key1":"value1"} format.
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) GetAttributeName() *string {
	return s.AttributeName
}

func (s *ListDataSourcesShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListDataSourcesShrinkRequest) SetAttributeName(v string) *ListDataSourcesShrinkRequest {
	s.AttributeName = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetFilterShrink(v string) *ListDataSourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
