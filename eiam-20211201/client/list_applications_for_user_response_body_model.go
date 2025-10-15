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
	// The applications that the EIAM account can access.
	Applications []*ListApplicationsForUserResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
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
	// The ID of the application that the EIAM account can access.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Indicates whether the EIAM account has direct permissions on the application. Valid values:
	//
	// 	- true: The EIAM account has direct permissions on the application.
	//
	// 	- false: The EIAM account does not have direct permissions on the application.
	//
	// example:
	//
	// true
	HasDirectAuthorization *bool `json:"HasDirectAuthorization,omitempty" xml:"HasDirectAuthorization,omitempty"`
	// Indicates whether the EIAM account has inherited permissions on the application. Valid values:
	//
	// 	- true: A parent organization or an organization to which the EIAM account belongs has direct permissions on the application.
	//
	// 	- false: A parent organization or an organization to which the EIAM account belongs does not have direct permissions on the application.
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

func (s *ListApplicationsForUserResponseBodyApplications) SetHasDirectAuthorization(v bool) *ListApplicationsForUserResponseBodyApplications {
	s.HasDirectAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) SetHasInheritAuthorization(v bool) *ListApplicationsForUserResponseBodyApplications {
	s.HasInheritAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
