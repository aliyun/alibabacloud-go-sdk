// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsForParentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeTags(v bool) *ListAccountsForParentRequest
	GetIncludeTags() *bool
	SetPageNumber(v int32) *ListAccountsForParentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccountsForParentRequest
	GetPageSize() *int32
	SetParentFolderId(v string) *ListAccountsForParentRequest
	GetParentFolderId() *string
	SetQueryKeyword(v string) *ListAccountsForParentRequest
	GetQueryKeyword() *string
	SetTag(v []*ListAccountsForParentRequestTag) *ListAccountsForParentRequest
	GetTag() []*ListAccountsForParentRequestTag
}

type ListAccountsForParentRequest struct {
	// Specifies whether to return tag information. Valid values:
	//
	// - false (default): Does not return tag information.
	//
	// - true: Returns tag information.
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource folder.
	//
	// example:
	//
	// fd-bVaRIG****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The query keyword, such as the member name.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// admin
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The tags used to filter member accounts.
	Tag []*ListAccountsForParentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsForParentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *ListAccountsForParentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccountsForParentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccountsForParentRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *ListAccountsForParentRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
}

func (s *ListAccountsForParentRequest) GetTag() []*ListAccountsForParentRequestTag {
	return s.Tag
}

func (s *ListAccountsForParentRequest) SetIncludeTags(v bool) *ListAccountsForParentRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageNumber(v int32) *ListAccountsForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageSize(v int32) *ListAccountsForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentRequest) SetParentFolderId(v string) *ListAccountsForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListAccountsForParentRequest) SetQueryKeyword(v string) *ListAccountsForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListAccountsForParentRequest) SetTag(v []*ListAccountsForParentRequestTag) *ListAccountsForParentRequest {
	s.Tag = v
	return s
}

func (s *ListAccountsForParentRequest) Validate() error {
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

type ListAccountsForParentRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsForParentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsForParentRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAccountsForParentRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAccountsForParentRequestTag) SetKey(v string) *ListAccountsForParentRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsForParentRequestTag) SetValue(v string) *ListAccountsForParentRequestTag {
	s.Value = &v
	return s
}

func (s *ListAccountsForParentRequestTag) Validate() error {
	return dara.Validate(s)
}
