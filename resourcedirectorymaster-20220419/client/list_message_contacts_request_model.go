// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *ListMessageContactsRequest
	GetContactId() *string
	SetMember(v string) *ListMessageContactsRequest
	GetMember() *string
	SetPageNumber(v int32) *ListMessageContactsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMessageContactsRequest
	GetPageSize() *int32
}

type ListMessageContactsRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The ID of the object to which the contact is bound. Valid values:
	//
	// 	- ID of the resource directory
	//
	// 	- ID of the folder
	//
	// 	- ID of the member
	//
	// example:
	//
	// fd-ZDNPiT****
	Member *string `json:"Member,omitempty" xml:"Member,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMessageContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactsRequest) GoString() string {
	return s.String()
}

func (s *ListMessageContactsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListMessageContactsRequest) GetMember() *string {
	return s.Member
}

func (s *ListMessageContactsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMessageContactsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageContactsRequest) SetContactId(v string) *ListMessageContactsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactsRequest) SetMember(v string) *ListMessageContactsRequest {
	s.Member = &v
	return s
}

func (s *ListMessageContactsRequest) SetPageNumber(v int32) *ListMessageContactsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactsRequest) SetPageSize(v int32) *ListMessageContactsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageContactsRequest) Validate() error {
	return dara.Validate(s)
}
