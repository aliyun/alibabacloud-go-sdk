// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypeRegistrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *ListResourceTypeRegistrationsRequest
	GetEntityType() *string
	SetPageNumber(v int32) *ListResourceTypeRegistrationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceTypeRegistrationsRequest
	GetPageSize() *int32
	SetRegistrationId(v string) *ListResourceTypeRegistrationsRequest
	GetRegistrationId() *string
	SetResourceType(v string) *ListResourceTypeRegistrationsRequest
	GetResourceType() *string
	SetStatus(v string) *ListResourceTypeRegistrationsRequest
	GetStatus() *string
}

type ListResourceTypeRegistrationsRequest struct {
	// The entity type. Set the value to Module.
	//
	// example:
	//
	// Module
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the registration record.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The resource type. The resource type can contain letters, digits, colons (:), and asterisks (\\*). You can use an asterisk (\\*) to perform a fuzzy match.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The registration state. Valid values:
	//
	// 	- IN_PROGRESS
	//
	// 	- COMPLETE
	//
	// 	- FAILED
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceTypeRegistrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceTypeRegistrationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceTypeRegistrationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceTypeRegistrationsRequest) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *ListResourceTypeRegistrationsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTypeRegistrationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceTypeRegistrationsRequest) SetEntityType(v string) *ListResourceTypeRegistrationsRequest {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetPageNumber(v int32) *ListResourceTypeRegistrationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetPageSize(v int32) *ListResourceTypeRegistrationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetRegistrationId(v string) *ListResourceTypeRegistrationsRequest {
	s.RegistrationId = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetResourceType(v string) *ListResourceTypeRegistrationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetStatus(v string) *ListResourceTypeRegistrationsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) Validate() error {
	return dara.Validate(s)
}
