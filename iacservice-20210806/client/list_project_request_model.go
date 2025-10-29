// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListProjectRequest
	GetKeyword() *string
	SetPageNumber(v string) *ListProjectRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListProjectRequest
	GetPageSize() *string
	SetTag(v []*ListProjectRequestTag) *ListProjectRequest
	GetTag() []*ListProjectRequestTag
}

type ListProjectRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Tag      []*ListProjectRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListProjectRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListProjectRequest) GetTag() []*ListProjectRequestTag {
	return s.Tag
}

func (s *ListProjectRequest) SetKeyword(v string) *ListProjectRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectRequest) SetPageNumber(v string) *ListProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRequest) SetPageSize(v string) *ListProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRequest) SetTag(v []*ListProjectRequestTag) *ListProjectRequest {
	s.Tag = v
	return s
}

func (s *ListProjectRequest) Validate() error {
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

type ListProjectRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRequestTag) GoString() string {
	return s.String()
}

func (s *ListProjectRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListProjectRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListProjectRequestTag) SetKey(v string) *ListProjectRequestTag {
	s.Key = &v
	return s
}

func (s *ListProjectRequestTag) SetValue(v string) *ListProjectRequestTag {
	s.Value = &v
	return s
}

func (s *ListProjectRequestTag) Validate() error {
	return dara.Validate(s)
}
