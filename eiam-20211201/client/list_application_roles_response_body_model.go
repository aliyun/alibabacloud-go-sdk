// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationRoles(v []*ListApplicationRolesResponseBodyApplicationRoles) *ListApplicationRolesResponseBody
	GetApplicationRoles() []*ListApplicationRolesResponseBodyApplicationRoles
	SetMaxResults(v int32) *ListApplicationRolesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationRolesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationRolesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationRolesResponseBody
	GetTotalCount() *int32
}

type ListApplicationRolesResponseBody struct {
	ApplicationRoles []*ListApplicationRolesResponseBodyApplicationRoles `json:"ApplicationRoles,omitempty" xml:"ApplicationRoles,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationRolesResponseBody) GetApplicationRoles() []*ListApplicationRolesResponseBodyApplicationRoles {
	return s.ApplicationRoles
}

func (s *ListApplicationRolesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationRolesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationRolesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationRolesResponseBody) SetApplicationRoles(v []*ListApplicationRolesResponseBodyApplicationRoles) *ListApplicationRolesResponseBody {
	s.ApplicationRoles = v
	return s
}

func (s *ListApplicationRolesResponseBody) SetMaxResults(v int32) *ListApplicationRolesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationRolesResponseBody) SetNextToken(v string) *ListApplicationRolesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationRolesResponseBody) SetRequestId(v string) *ListApplicationRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationRolesResponseBody) SetTotalCount(v int32) *ListApplicationRolesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationRolesResponseBody) Validate() error {
	if s.ApplicationRoles != nil {
		for _, item := range s.ApplicationRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationRolesResponseBodyApplicationRoles struct {
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色的唯一标识
	//
	// example:
	//
	// approle_01kh2vuo8v9splv8maak1d22rxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// 应用角色名称
	//
	// example:
	//
	// Admin Role
	ApplicationRoleName *string `json:"ApplicationRoleName,omitempty" xml:"ApplicationRoleName,omitempty"`
	// example:
	//
	// admin_role
	ApplicationRoleValue *string `json:"ApplicationRoleValue,omitempty" xml:"ApplicationRoleValue,omitempty"`
	// example:
	//
	// Admin Role Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationRolesResponseBodyApplicationRoles) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRolesResponseBodyApplicationRoles) GoString() string {
	return s.String()
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) GetApplicationRoleName() *string {
	return s.ApplicationRoleName
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) GetApplicationRoleValue() *string {
	return s.ApplicationRoleValue
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) SetApplicationId(v string) *ListApplicationRolesResponseBodyApplicationRoles {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) SetApplicationRoleId(v string) *ListApplicationRolesResponseBodyApplicationRoles {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) SetApplicationRoleName(v string) *ListApplicationRolesResponseBodyApplicationRoles {
	s.ApplicationRoleName = &v
	return s
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) SetApplicationRoleValue(v string) *ListApplicationRolesResponseBodyApplicationRoles {
	s.ApplicationRoleValue = &v
	return s
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) SetDescription(v string) *ListApplicationRolesResponseBodyApplicationRoles {
	s.Description = &v
	return s
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) SetInstanceId(v string) *ListApplicationRolesResponseBodyApplicationRoles {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationRolesResponseBodyApplicationRoles) Validate() error {
	return dara.Validate(s)
}
