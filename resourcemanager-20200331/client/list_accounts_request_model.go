// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeTags(v bool) *ListAccountsRequest
	GetIncludeTags() *bool
	SetPageNumber(v int32) *ListAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccountsRequest
	GetPageSize() *int32
	SetTag(v []*ListAccountsRequestTag) *ListAccountsRequest
	GetTag() []*ListAccountsRequestTag
}

type ListAccountsRequest struct {
	// Specifies whether to return the information of tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tag key and value.
	Tag []*ListAccountsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *ListAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccountsRequest) GetTag() []*ListAccountsRequestTag {
	return s.Tag
}

func (s *ListAccountsRequest) SetIncludeTags(v bool) *ListAccountsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsRequest) SetPageNumber(v int32) *ListAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsRequest) SetPageSize(v int32) *ListAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsRequest) SetTag(v []*ListAccountsRequestTag) *ListAccountsRequest {
	s.Tag = v
	return s
}

func (s *ListAccountsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountsRequestTag struct {
	// A tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAccountsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAccountsRequestTag) SetKey(v string) *ListAccountsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsRequestTag) SetValue(v string) *ListAccountsRequestTag {
	s.Value = &v
	return s
}

func (s *ListAccountsRequestTag) Validate() error {
	return dara.Validate(s)
}
