// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v string) *DescribeContactsRequest
	GetContactIds() *string
	SetContactName(v string) *DescribeContactsRequest
	GetContactName() *string
	SetEmail(v string) *DescribeContactsRequest
	GetEmail() *string
	SetPage(v int64) *DescribeContactsRequest
	GetPage() *int64
	SetPhone(v string) *DescribeContactsRequest
	GetPhone() *string
	SetRegionId(v string) *DescribeContactsRequest
	GetRegionId() *string
	SetSize(v int64) *DescribeContactsRequest
	GetSize() *int64
	SetVerbose(v string) *DescribeContactsRequest
	GetVerbose() *string
}

type DescribeContactsRequest struct {
	// The ID of the alert contact that you want to query. Separate multiple contact IDs with spaces.
	//
	// example:
	//
	// 123,321
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// John Doe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1381111*****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of alert contacts to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Specifies whether to return redundant information.
	//
	// example:
	//
	// true
	Verbose *string `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s DescribeContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactsRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *DescribeContactsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *DescribeContactsRequest) GetEmail() *string {
	return s.Email
}

func (s *DescribeContactsRequest) GetPage() *int64 {
	return s.Page
}

func (s *DescribeContactsRequest) GetPhone() *string {
	return s.Phone
}

func (s *DescribeContactsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContactsRequest) GetSize() *int64 {
	return s.Size
}

func (s *DescribeContactsRequest) GetVerbose() *string {
	return s.Verbose
}

func (s *DescribeContactsRequest) SetContactIds(v string) *DescribeContactsRequest {
	s.ContactIds = &v
	return s
}

func (s *DescribeContactsRequest) SetContactName(v string) *DescribeContactsRequest {
	s.ContactName = &v
	return s
}

func (s *DescribeContactsRequest) SetEmail(v string) *DescribeContactsRequest {
	s.Email = &v
	return s
}

func (s *DescribeContactsRequest) SetPage(v int64) *DescribeContactsRequest {
	s.Page = &v
	return s
}

func (s *DescribeContactsRequest) SetPhone(v string) *DescribeContactsRequest {
	s.Phone = &v
	return s
}

func (s *DescribeContactsRequest) SetRegionId(v string) *DescribeContactsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContactsRequest) SetSize(v int64) *DescribeContactsRequest {
	s.Size = &v
	return s
}

func (s *DescribeContactsRequest) SetVerbose(v string) *DescribeContactsRequest {
	s.Verbose = &v
	return s
}

func (s *DescribeContactsRequest) Validate() error {
	return dara.Validate(s)
}
