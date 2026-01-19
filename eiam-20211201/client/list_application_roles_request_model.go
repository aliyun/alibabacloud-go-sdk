// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListApplicationRolesRequest
	GetApplicationId() *string
	SetFilter(v []*ListApplicationRolesRequestFilter) *ListApplicationRolesRequest
	GetFilter() []*ListApplicationRolesRequestFilter
	SetInstanceId(v string) *ListApplicationRolesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListApplicationRolesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationRolesRequest
	GetNextToken() *string
}

type ListApplicationRolesRequest struct {
	// 应用ID
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string                              `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Filter        []*ListApplicationRolesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListApplicationRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRolesRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRolesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationRolesRequest) GetFilter() []*ListApplicationRolesRequestFilter {
	return s.Filter
}

func (s *ListApplicationRolesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationRolesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationRolesRequest) SetApplicationId(v string) *ListApplicationRolesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationRolesRequest) SetFilter(v []*ListApplicationRolesRequestFilter) *ListApplicationRolesRequest {
	s.Filter = v
	return s
}

func (s *ListApplicationRolesRequest) SetInstanceId(v string) *ListApplicationRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationRolesRequest) SetMaxResults(v int32) *ListApplicationRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationRolesRequest) SetNextToken(v string) *ListApplicationRolesRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationRolesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationRolesRequestFilter struct {
	// example:
	//
	// ApplicationRoleNameStartsWith
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListApplicationRolesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRolesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListApplicationRolesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListApplicationRolesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListApplicationRolesRequestFilter) SetName(v string) *ListApplicationRolesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListApplicationRolesRequestFilter) SetValue(v []*string) *ListApplicationRolesRequestFilter {
	s.Value = v
	return s
}

func (s *ListApplicationRolesRequestFilter) Validate() error {
	return dara.Validate(s)
}
