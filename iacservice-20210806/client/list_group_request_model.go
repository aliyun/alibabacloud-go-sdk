// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListGroupRequest
	GetKeyword() *string
	SetPageNumber(v string) *ListGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListGroupRequest
	GetPageSize() *string
	SetProjectId(v string) *ListGroupRequest
	GetProjectId() *string
	SetTag(v []*ListGroupRequestTag) *ListGroupRequest
	GetTag() []*ListGroupRequestTag
}

type ListGroupRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 200
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-14e80de4866bf7ffed0c4072ed9b37
	ProjectId *string                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Tag       []*ListGroupRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupRequest) GoString() string {
	return s.String()
}

func (s *ListGroupRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGroupRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListGroupRequest) GetTag() []*ListGroupRequestTag {
	return s.Tag
}

func (s *ListGroupRequest) SetKeyword(v string) *ListGroupRequest {
	s.Keyword = &v
	return s
}

func (s *ListGroupRequest) SetPageNumber(v string) *ListGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupRequest) SetPageSize(v string) *ListGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupRequest) SetProjectId(v string) *ListGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGroupRequest) SetTag(v []*ListGroupRequestTag) *ListGroupRequest {
	s.Tag = v
	return s
}

func (s *ListGroupRequest) Validate() error {
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

type ListGroupRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListGroupRequestTag) GoString() string {
	return s.String()
}

func (s *ListGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListGroupRequestTag) SetKey(v string) *ListGroupRequestTag {
	s.Key = &v
	return s
}

func (s *ListGroupRequestTag) SetValue(v string) *ListGroupRequestTag {
	s.Value = &v
	return s
}

func (s *ListGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
