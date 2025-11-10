// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersForParentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListFoldersForParentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFoldersForParentRequest
	GetPageSize() *int32
	SetParentFolderId(v string) *ListFoldersForParentRequest
	GetParentFolderId() *string
	SetQueryKeyword(v string) *ListFoldersForParentRequest
	GetQueryKeyword() *string
	SetTag(v []*ListFoldersForParentRequestTag) *ListFoldersForParentRequest
	GetTag() []*ListFoldersForParentRequestTag
}

type ListFoldersForParentRequest struct {
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
	// The ID of the parent folder.
	//
	// If you leave this parameter empty, the information of the first-level subfolders of the Root folder is queried.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The keyword used for the query, such as a folder name.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// rdFolder
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The tags. This parameter specifies a filter condition.
	Tag []*ListFoldersForParentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListFoldersForParentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentRequest) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFoldersForParentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFoldersForParentRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *ListFoldersForParentRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
}

func (s *ListFoldersForParentRequest) GetTag() []*ListFoldersForParentRequestTag {
	return s.Tag
}

func (s *ListFoldersForParentRequest) SetPageNumber(v int32) *ListFoldersForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentRequest) SetPageSize(v int32) *ListFoldersForParentRequest {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentRequest) SetParentFolderId(v string) *ListFoldersForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListFoldersForParentRequest) SetQueryKeyword(v string) *ListFoldersForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListFoldersForParentRequest) SetTag(v []*ListFoldersForParentRequestTag) *ListFoldersForParentRequest {
	s.Tag = v
	return s
}

func (s *ListFoldersForParentRequest) Validate() error {
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

type ListFoldersForParentRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFoldersForParentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentRequestTag) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListFoldersForParentRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListFoldersForParentRequestTag) SetKey(v string) *ListFoldersForParentRequestTag {
	s.Key = &v
	return s
}

func (s *ListFoldersForParentRequestTag) SetValue(v string) *ListFoldersForParentRequestTag {
	s.Value = &v
	return s
}

func (s *ListFoldersForParentRequestTag) Validate() error {
	return dara.Validate(s)
}
