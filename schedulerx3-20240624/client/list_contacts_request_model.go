// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactName(v string) *ListContactsRequest
	GetContactName() *string
	SetPageNum(v int32) *ListContactsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListContactsRequest
	GetPageSize() *int32
}

type ListContactsRequest struct {
	// The name of the alert contact.
	//
	// example:
	//
	// Tom
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactsRequest) GoString() string {
	return s.String()
}

func (s *ListContactsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *ListContactsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListContactsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContactsRequest) SetContactName(v string) *ListContactsRequest {
	s.ContactName = &v
	return s
}

func (s *ListContactsRequest) SetPageNum(v int32) *ListContactsRequest {
	s.PageNum = &v
	return s
}

func (s *ListContactsRequest) SetPageSize(v int32) *ListContactsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactsRequest) Validate() error {
	return dara.Validate(s)
}
