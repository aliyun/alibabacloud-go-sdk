// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeActivationCodesResponseBodyItems) *DescribeActivationCodesResponseBody
	GetItems() []*DescribeActivationCodesResponseBodyItems
	SetPageNumber(v int32) *DescribeActivationCodesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeActivationCodesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeActivationCodesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeActivationCodesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeActivationCodesResponseBody struct {
	// The queried activation codes.
	Items []*DescribeActivationCodesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65D7ACE6-4A61-4B6E-B357-8CB24A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeActivationCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodesResponseBody) GetItems() []*DescribeActivationCodesResponseBodyItems {
	return s.Items
}

func (s *DescribeActivationCodesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActivationCodesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeActivationCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActivationCodesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeActivationCodesResponseBody) SetItems(v []*DescribeActivationCodesResponseBodyItems) *DescribeActivationCodesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeActivationCodesResponseBody) SetPageNumber(v int32) *DescribeActivationCodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActivationCodesResponseBody) SetPageRecordCount(v int32) *DescribeActivationCodesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeActivationCodesResponseBody) SetRequestId(v string) *DescribeActivationCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActivationCodesResponseBody) SetTotalRecordCount(v int32) *DescribeActivationCodesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeActivationCodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeActivationCodesResponseBodyItems struct {
	// The time when the activation code takes effect.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	ActivateAt *string `json:"ActivateAt,omitempty" xml:"ActivateAt,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// testCode
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the activation code expires.
	//
	// example:
	//
	// 2054-10-16 16:46:20
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// The time when the activation code was generated.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the activation code was updated.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The activation code ID.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The media access control (MAC) address used in the generation of the activation code.
	//
	// example:
	//
	// 12:34:56:78:98:00
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The name of the activation code.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the database.
	//
	// example:
	//
	// 1234567890123456
	SystemIdentifier *string `json:"SystemIdentifier,omitempty" xml:"SystemIdentifier,omitempty"`
}

func (s DescribeActivationCodesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodesResponseBodyItems) GetActivateAt() *string {
	return s.ActivateAt
}

func (s *DescribeActivationCodesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeActivationCodesResponseBodyItems) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *DescribeActivationCodesResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeActivationCodesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeActivationCodesResponseBodyItems) GetId() *int32 {
	return s.Id
}

func (s *DescribeActivationCodesResponseBodyItems) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeActivationCodesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeActivationCodesResponseBodyItems) GetSystemIdentifier() *string {
	return s.SystemIdentifier
}

func (s *DescribeActivationCodesResponseBodyItems) SetActivateAt(v string) *DescribeActivationCodesResponseBodyItems {
	s.ActivateAt = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetDescription(v string) *DescribeActivationCodesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetExpireAt(v string) *DescribeActivationCodesResponseBodyItems {
	s.ExpireAt = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetGmtCreated(v string) *DescribeActivationCodesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetGmtModified(v string) *DescribeActivationCodesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetId(v int32) *DescribeActivationCodesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetMacAddress(v string) *DescribeActivationCodesResponseBodyItems {
	s.MacAddress = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetName(v string) *DescribeActivationCodesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) SetSystemIdentifier(v string) *DescribeActivationCodesResponseBodyItems {
	s.SystemIdentifier = &v
	return s
}

func (s *DescribeActivationCodesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
