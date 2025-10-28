// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypeRegistrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceTypeRegistrationsResponseBody
	GetPageNumber() *int32
	SetRegistrations(v []*ListResourceTypeRegistrationsResponseBodyRegistrations) *ListResourceTypeRegistrationsResponseBody
	GetRegistrations() []*ListResourceTypeRegistrationsResponseBodyRegistrations
	SetRequestId(v string) *ListResourceTypeRegistrationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourceTypeRegistrationsResponseBody
	GetTotalCount() *int32
}

type ListResourceTypeRegistrationsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The registration records of the resource.
	Registrations []*ListResourceTypeRegistrationsResponseBodyRegistrations `json:"Registrations,omitempty" xml:"Registrations,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of registration records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceTypeRegistrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceTypeRegistrationsResponseBody) GetRegistrations() []*ListResourceTypeRegistrationsResponseBodyRegistrations {
	return s.Registrations
}

func (s *ListResourceTypeRegistrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTypeRegistrationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceTypeRegistrationsResponseBody) SetPageNumber(v int32) *ListResourceTypeRegistrationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) SetRegistrations(v []*ListResourceTypeRegistrationsResponseBodyRegistrations) *ListResourceTypeRegistrationsResponseBody {
	s.Registrations = v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) SetRequestId(v string) *ListResourceTypeRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) SetTotalCount(v int32) *ListResourceTypeRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) Validate() error {
	if s.Registrations != nil {
		for _, item := range s.Registrations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceTypeRegistrationsResponseBodyRegistrations struct {
	// The creation time. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2023-03-02T07:28:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The entity type. Only Module may be returned.
	//
	// example:
	//
	// Module
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the registration record.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The registration state. Valid values:
	//
	// 	- IN_PROGRESS: The registration is in progress.
	//
	// 	- COMPLETE: The registration is successful.
	//
	// 	- FAILED: The registration failed.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason for the state.
	//
	// example:
	//
	// Module is created successfully
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The version ID.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListResourceTypeRegistrationsResponseBodyRegistrations) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeRegistrationsResponseBodyRegistrations) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetStatus() *string {
	return s.Status
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) GetVersionId() *string {
	return s.VersionId
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetCreateTime(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.CreateTime = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetEntityType(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetRegistrationId(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetResourceType(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetStatus(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.Status = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetStatusReason(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.StatusReason = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetVersionId(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.VersionId = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) Validate() error {
	return dara.Validate(s)
}
