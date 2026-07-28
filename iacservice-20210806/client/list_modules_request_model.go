// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListModulesRequest
	GetGroupId() *string
	SetKeyword(v string) *ListModulesRequest
	GetKeyword() *string
	SetModuleName(v string) *ListModulesRequest
	GetModuleName() *string
	SetPageNumber(v int32) *ListModulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModulesRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListModulesRequest
	GetProjectId() *string
	SetTag(v []*ListModulesRequestTag) *ListModulesRequest
	GetTag() []*ListModulesRequestTag
}

type ListModulesRequest struct {
	// The group ID.
	//
	// example:
	//
	// g-kw1a50tj8rk7cki2q8bbat
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The search keyword. Fuzzy match is supported for template names.
	//
	// example:
	//
	// key
	Keyword    *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-al1c58tb2lu9oej36kclvf
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The list of template tags.
	Tag []*ListModulesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListModulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModulesRequest) GoString() string {
	return s.String()
}

func (s *ListModulesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListModulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListModulesRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListModulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModulesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModulesRequest) GetTag() []*ListModulesRequestTag {
	return s.Tag
}

func (s *ListModulesRequest) SetGroupId(v string) *ListModulesRequest {
	s.GroupId = &v
	return s
}

func (s *ListModulesRequest) SetKeyword(v string) *ListModulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListModulesRequest) SetModuleName(v string) *ListModulesRequest {
	s.ModuleName = &v
	return s
}

func (s *ListModulesRequest) SetPageNumber(v int32) *ListModulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulesRequest) SetPageSize(v int32) *ListModulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulesRequest) SetProjectId(v string) *ListModulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModulesRequest) SetTag(v []*ListModulesRequestTag) *ListModulesRequest {
	s.Tag = v
	return s
}

func (s *ListModulesRequest) Validate() error {
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

type ListModulesRequestTag struct {
	// The tag key of the template.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the template.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListModulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListModulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListModulesRequestTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListModulesRequestTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListModulesRequestTag) SetTagKey(v string) *ListModulesRequestTag {
	s.TagKey = &v
	return s
}

func (s *ListModulesRequestTag) SetTagValue(v string) *ListModulesRequestTag {
	s.TagValue = &v
	return s
}

func (s *ListModulesRequestTag) Validate() error {
	return dara.Validate(s)
}
