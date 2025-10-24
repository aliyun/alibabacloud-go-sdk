// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPersonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListCustomPersonsRequest
	GetCategoryId() *string
	SetOwnerAccount(v string) *ListCustomPersonsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCustomPersonsRequest
	GetOwnerId() *int64
	SetPersonId(v string) *ListCustomPersonsRequest
	GetPersonId() *string
	SetResourceOwnerAccount(v string) *ListCustomPersonsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCustomPersonsRequest
	GetResourceOwnerId() *int64
}

type ListCustomPersonsRequest struct {
	// The ID of the figure library about which you want to query information. The ID is used to uniquely identify a custom figure library. Make sure that the ID is unique. If you do not specify this parameter, the operation returns all the custom figure libraries. The ID can be up to 120 characters in length and is not case-sensitive.
	//
	// > You cannot specify the ID of the system figure library for this parameter.
	//
	// example:
	//
	// CategoryId-****
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the figure about which you want to query information. The ID is used to uniquely identify a figure. Make sure that the ID is unique. If you do not specify this parameter, the operation returns the information about all the figures in the specified figure library.
	//
	// example:
	//
	// PersonId-****
	PersonId             *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListCustomPersonsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListCustomPersonsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCustomPersonsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCustomPersonsRequest) GetPersonId() *string {
	return s.PersonId
}

func (s *ListCustomPersonsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCustomPersonsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCustomPersonsRequest) SetCategoryId(v string) *ListCustomPersonsRequest {
	s.CategoryId = &v
	return s
}

func (s *ListCustomPersonsRequest) SetOwnerAccount(v string) *ListCustomPersonsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCustomPersonsRequest) SetOwnerId(v int64) *ListCustomPersonsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCustomPersonsRequest) SetPersonId(v string) *ListCustomPersonsRequest {
	s.PersonId = &v
	return s
}

func (s *ListCustomPersonsRequest) SetResourceOwnerAccount(v string) *ListCustomPersonsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCustomPersonsRequest) SetResourceOwnerId(v int64) *ListCustomPersonsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCustomPersonsRequest) Validate() error {
	return dara.Validate(s)
}
