// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchWorkspaceRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SearchWorkspaceRolesResponseBodyData) *SearchWorkspaceRolesResponseBody
	GetData() []*SearchWorkspaceRolesResponseBodyData
	SetMaxResults(v int32) *SearchWorkspaceRolesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchWorkspaceRolesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchWorkspaceRolesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchWorkspaceRolesResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *SearchWorkspaceRolesResponseBody
	GetTotalCount() *string
}

type SearchWorkspaceRolesResponseBody struct {
	Data []*SearchWorkspaceRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchWorkspaceRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchWorkspaceRolesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceRolesResponseBody) GetData() []*SearchWorkspaceRolesResponseBodyData {
	return s.Data
}

func (s *SearchWorkspaceRolesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchWorkspaceRolesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchWorkspaceRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchWorkspaceRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchWorkspaceRolesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *SearchWorkspaceRolesResponseBody) SetData(v []*SearchWorkspaceRolesResponseBodyData) *SearchWorkspaceRolesResponseBody {
	s.Data = v
	return s
}

func (s *SearchWorkspaceRolesResponseBody) SetMaxResults(v int32) *SearchWorkspaceRolesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBody) SetNextToken(v string) *SearchWorkspaceRolesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBody) SetRequestId(v string) *SearchWorkspaceRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBody) SetSuccess(v bool) *SearchWorkspaceRolesResponseBody {
	s.Success = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBody) SetTotalCount(v string) *SearchWorkspaceRolesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchWorkspaceRolesResponseBodyData struct {
	// example:
	//
	// MANAGER
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// SLBLogDefaultRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// INNER
	RoleSource *string `json:"RoleSource,omitempty" xml:"RoleSource,omitempty"`
	// example:
	//
	// inner
	RoleSourceName *string `json:"RoleSourceName,omitempty" xml:"RoleSourceName,omitempty"`
}

func (s SearchWorkspaceRolesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchWorkspaceRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceRolesResponseBodyData) GetRoleId() *int64 {
	return s.RoleId
}

func (s *SearchWorkspaceRolesResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *SearchWorkspaceRolesResponseBodyData) GetRoleSource() *string {
	return s.RoleSource
}

func (s *SearchWorkspaceRolesResponseBodyData) GetRoleSourceName() *string {
	return s.RoleSourceName
}

func (s *SearchWorkspaceRolesResponseBodyData) SetRoleId(v int64) *SearchWorkspaceRolesResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBodyData) SetRoleName(v string) *SearchWorkspaceRolesResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBodyData) SetRoleSource(v string) *SearchWorkspaceRolesResponseBodyData {
	s.RoleSource = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBodyData) SetRoleSourceName(v string) *SearchWorkspaceRolesResponseBodyData {
	s.RoleSourceName = &v
	return s
}

func (s *SearchWorkspaceRolesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
