// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageContactVerificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *ListMessageContactVerificationsRequest
	GetContactId() *string
	SetPageNumber(v int32) *ListMessageContactVerificationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMessageContactVerificationsRequest
	GetPageSize() *int32
}

type ListMessageContactVerificationsRequest struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
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

func (s ListMessageContactVerificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactVerificationsRequest) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListMessageContactVerificationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMessageContactVerificationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageContactVerificationsRequest) SetContactId(v string) *ListMessageContactVerificationsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactVerificationsRequest) SetPageNumber(v int32) *ListMessageContactVerificationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactVerificationsRequest) SetPageSize(v int32) *ListMessageContactVerificationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageContactVerificationsRequest) Validate() error {
	return dara.Validate(s)
}
