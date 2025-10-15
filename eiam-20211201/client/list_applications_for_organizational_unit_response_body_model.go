// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsForOrganizationalUnitResponseBodyApplications) *ListApplicationsForOrganizationalUnitResponseBody
	GetApplications() []*ListApplicationsForOrganizationalUnitResponseBodyApplications
	SetRequestId(v string) *ListApplicationsForOrganizationalUnitResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsForOrganizationalUnitResponseBody
	GetTotalCount() *int64
}

type ListApplicationsForOrganizationalUnitResponseBody struct {
	// The applications that the EIAM organization can access.
	Applications []*ListApplicationsForOrganizationalUnitResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
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

func (s ListApplicationsForOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) GetApplications() []*ListApplicationsForOrganizationalUnitResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) SetApplications(v []*ListApplicationsForOrganizationalUnitResponseBodyApplications) *ListApplicationsForOrganizationalUnitResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) SetRequestId(v string) *ListApplicationsForOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) SetTotalCount(v int64) *ListApplicationsForOrganizationalUnitResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) Validate() error {
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

type ListApplicationsForOrganizationalUnitResponseBodyApplications struct {
	// The ID of the application that the EIAM organization can access.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s ListApplicationsForOrganizationalUnitResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForOrganizationalUnitResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForOrganizationalUnitResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
