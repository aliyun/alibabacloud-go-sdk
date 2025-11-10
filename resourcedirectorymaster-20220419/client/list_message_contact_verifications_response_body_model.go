// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageContactVerificationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactVerifications(v []*ListMessageContactVerificationsResponseBodyContactVerifications) *ListMessageContactVerificationsResponseBody
	GetContactVerifications() []*ListMessageContactVerificationsResponseBodyContactVerifications
	SetPageNumber(v int32) *ListMessageContactVerificationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMessageContactVerificationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMessageContactVerificationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMessageContactVerificationsResponseBody
	GetTotalCount() *int32
}

type ListMessageContactVerificationsResponseBody struct {
	// The record for the contact to be verified.
	ContactVerifications []*ListMessageContactVerificationsResponseBodyContactVerifications `json:"ContactVerifications,omitempty" xml:"ContactVerifications,omitempty" type:"Repeated"`
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
	// The request ID.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 48
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessageContactVerificationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactVerificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsResponseBody) GetContactVerifications() []*ListMessageContactVerificationsResponseBodyContactVerifications {
	return s.ContactVerifications
}

func (s *ListMessageContactVerificationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMessageContactVerificationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageContactVerificationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageContactVerificationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMessageContactVerificationsResponseBody) SetContactVerifications(v []*ListMessageContactVerificationsResponseBodyContactVerifications) *ListMessageContactVerificationsResponseBody {
	s.ContactVerifications = v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetPageNumber(v int32) *ListMessageContactVerificationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetPageSize(v int32) *ListMessageContactVerificationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetRequestId(v string) *ListMessageContactVerificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) SetTotalCount(v int32) *ListMessageContactVerificationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBody) Validate() error {
	if s.ContactVerifications != nil {
		for _, item := range s.ContactVerifications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessageContactVerificationsResponseBodyContactVerifications struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The object that is used for verification. Valid values:
	//
	// - Mobile phone number
	//
	// - Email address
	//
	// example:
	//
	// someone***@example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListMessageContactVerificationsResponseBodyContactVerifications) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactVerificationsResponseBodyContactVerifications) GoString() string {
	return s.String()
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) GetContactId() *string {
	return s.ContactId
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) GetTarget() *string {
	return s.Target
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) SetContactId(v string) *ListMessageContactVerificationsResponseBodyContactVerifications {
	s.ContactId = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) SetTarget(v string) *ListMessageContactVerificationsResponseBodyContactVerifications {
	s.Target = &v
	return s
}

func (s *ListMessageContactVerificationsResponseBodyContactVerifications) Validate() error {
	return dara.Validate(s)
}
