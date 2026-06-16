// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsForUserResponseBodyApplications) *ListApplicationsForUserResponseBody
	GetApplications() []*ListApplicationsForUserResponseBodyApplications
	SetRequestId(v string) *ListApplicationsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsForUserResponseBody
	GetTotalCount() *int64
}

type ListApplicationsForUserResponseBody struct {
	// The list of applications that the account is authorized to access.
	Applications []*ListApplicationsForUserResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponseBody) GetApplications() []*ListApplicationsForUserResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsForUserResponseBody) SetApplications(v []*ListApplicationsForUserResponseBodyApplications) *ListApplicationsForUserResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForUserResponseBody) SetRequestId(v string) *ListApplicationsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForUserResponseBody) SetTotalCount(v int64) *ListApplicationsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsForUserResponseBody) Validate() error {
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

type ListApplicationsForUserResponseBodyApplications struct {
	// The application ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The list of application roles.
	ApplicationRoles []*ListApplicationsForUserResponseBodyApplicationsApplicationRoles `json:"ApplicationRoles,omitempty" xml:"ApplicationRoles,omitempty" type:"Repeated"`
	// Indicates whether a direct authorization exists. Valid values:
	//
	// - true: A direct authorization record exists between the application and the account.
	//
	// - false: No direct authorization record exists between the application and the account.
	//
	// example:
	//
	// true
	HasDirectAuthorization *bool `json:"HasDirectAuthorization,omitempty" xml:"HasDirectAuthorization,omitempty"`
	// Indicates whether an inherited authorization exists. Valid values:
	//
	// - true: A direct authorization record exists between the application and a parent organization or a group to which the account belongs.
	//
	// - false: No direct authorization record exists between the application and any of the parent organizations or groups to which the account belongs.
	//
	// example:
	//
	// false
	HasInheritAuthorization *bool `json:"HasInheritAuthorization,omitempty" xml:"HasInheritAuthorization,omitempty"`
}

func (s ListApplicationsForUserResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForUserResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForUserResponseBodyApplications) GetApplicationRoles() []*ListApplicationsForUserResponseBodyApplicationsApplicationRoles {
	return s.ApplicationRoles
}

func (s *ListApplicationsForUserResponseBodyApplications) GetHasDirectAuthorization() *bool {
	return s.HasDirectAuthorization
}

func (s *ListApplicationsForUserResponseBodyApplications) GetHasInheritAuthorization() *bool {
	return s.HasInheritAuthorization
}

func (s *ListApplicationsForUserResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForUserResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) SetApplicationRoles(v []*ListApplicationsForUserResponseBodyApplicationsApplicationRoles) *ListApplicationsForUserResponseBodyApplications {
	s.ApplicationRoles = v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) SetHasDirectAuthorization(v bool) *ListApplicationsForUserResponseBodyApplications {
	s.HasDirectAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) SetHasInheritAuthorization(v bool) *ListApplicationsForUserResponseBodyApplications {
	s.HasInheritAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) Validate() error {
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

type ListApplicationsForUserResponseBodyApplicationsApplicationRoles struct {
	// The ID of the application role.
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// Indicates whether the role is directly assigned to the user.
	//
	// example:
	//
	// true
	HasDirectAuthorization *bool `json:"HasDirectAuthorization,omitempty" xml:"HasDirectAuthorization,omitempty"`
	// Indicates whether the role is inherited from an organization or a group to which the user belongs.
	//
	// example:
	//
	// false
	HasInheritAuthorization *bool `json:"HasInheritAuthorization,omitempty" xml:"HasInheritAuthorization,omitempty"`
}

func (s ListApplicationsForUserResponseBodyApplicationsApplicationRoles) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForUserResponseBodyApplicationsApplicationRoles) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) GetHasDirectAuthorization() *bool {
	return s.HasDirectAuthorization
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) GetHasInheritAuthorization() *bool {
	return s.HasInheritAuthorization
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) SetApplicationRoleId(v string) *ListApplicationsForUserResponseBodyApplicationsApplicationRoles {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) SetHasDirectAuthorization(v bool) *ListApplicationsForUserResponseBodyApplicationsApplicationRoles {
	s.HasDirectAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) SetHasInheritAuthorization(v bool) *ListApplicationsForUserResponseBodyApplicationsApplicationRoles {
	s.HasInheritAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplicationsApplicationRoles) Validate() error {
	return dara.Validate(s)
}
