// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsForGroupResponseBodyApplications) *ListApplicationsForGroupResponseBody
	GetApplications() []*ListApplicationsForGroupResponseBodyApplications
	SetRequestId(v string) *ListApplicationsForGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsForGroupResponseBody
	GetTotalCount() *int64
}

type ListApplicationsForGroupResponseBody struct {
	Applications []*ListApplicationsForGroupResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsForGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForGroupResponseBody) GetApplications() []*ListApplicationsForGroupResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsForGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForGroupResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsForGroupResponseBody) SetApplications(v []*ListApplicationsForGroupResponseBodyApplications) *ListApplicationsForGroupResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForGroupResponseBody) SetRequestId(v string) *ListApplicationsForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForGroupResponseBody) SetTotalCount(v int64) *ListApplicationsForGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsForGroupResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForGroupResponseBodyApplications struct {
	// 应用的唯一标识。
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色列表。
	ApplicationRoles []*ListApplicationsForGroupResponseBodyApplicationsApplicationRoles `json:"ApplicationRoles,omitempty" xml:"ApplicationRoles,omitempty" type:"Repeated"`
	// 直接分配给当前用户的权限，视为直接授权。
	//
	// example:
	//
	// true
	HasDirectAuthorization *bool `json:"HasDirectAuthorization,omitempty" xml:"HasDirectAuthorization,omitempty"`
	// 通过用户隶属的组织、组获取的权限，视为继承权限。
	//
	// example:
	//
	// false
	HasInheritAuthorization *bool `json:"HasInheritAuthorization,omitempty" xml:"HasInheritAuthorization,omitempty"`
}

func (s ListApplicationsForGroupResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForGroupResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForGroupResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForGroupResponseBodyApplications) GetApplicationRoles() []*ListApplicationsForGroupResponseBodyApplicationsApplicationRoles {
	return s.ApplicationRoles
}

func (s *ListApplicationsForGroupResponseBodyApplications) GetHasDirectAuthorization() *bool {
	return s.HasDirectAuthorization
}

func (s *ListApplicationsForGroupResponseBodyApplications) GetHasInheritAuthorization() *bool {
	return s.HasInheritAuthorization
}

func (s *ListApplicationsForGroupResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForGroupResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForGroupResponseBodyApplications) SetApplicationRoles(v []*ListApplicationsForGroupResponseBodyApplicationsApplicationRoles) *ListApplicationsForGroupResponseBodyApplications {
	s.ApplicationRoles = v
	return s
}

func (s *ListApplicationsForGroupResponseBodyApplications) SetHasDirectAuthorization(v bool) *ListApplicationsForGroupResponseBodyApplications {
	s.HasDirectAuthorization = &v
	return s
}

func (s *ListApplicationsForGroupResponseBodyApplications) SetHasInheritAuthorization(v bool) *ListApplicationsForGroupResponseBodyApplications {
	s.HasInheritAuthorization = &v
	return s
}

func (s *ListApplicationsForGroupResponseBodyApplications) Validate() error {
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

type ListApplicationsForGroupResponseBodyApplicationsApplicationRoles struct {
	// 应用角色标识。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
}

func (s ListApplicationsForGroupResponseBodyApplicationsApplicationRoles) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForGroupResponseBodyApplicationsApplicationRoles) GoString() string {
	return s.String()
}

func (s *ListApplicationsForGroupResponseBodyApplicationsApplicationRoles) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListApplicationsForGroupResponseBodyApplicationsApplicationRoles) SetApplicationRoleId(v string) *ListApplicationsForGroupResponseBodyApplicationsApplicationRoles {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListApplicationsForGroupResponseBodyApplicationsApplicationRoles) Validate() error {
	return dara.Validate(s)
}
