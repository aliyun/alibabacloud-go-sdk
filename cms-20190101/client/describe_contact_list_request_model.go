// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChanelType(v string) *DescribeContactListRequest
	GetChanelType() *string
	SetChanelValue(v string) *DescribeContactListRequest
	GetChanelValue() *string
	SetContactName(v string) *DescribeContactListRequest
	GetContactName() *string
	SetPageNumber(v int32) *DescribeContactListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeContactListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeContactListRequest
	GetRegionId() *string
}

type DescribeContactListRequest struct {
	// The alert notification method. Valid values:
	//
	// <props="china">- SMS: text message.
	//
	// <props="china">- AliIM: TradeManager.
	//
	// - Mail: email.
	//
	// - DingWebHook: DingTalk chatbot.
	//
	// example:
	//
	// Mail
	ChanelType *string `json:"ChanelType,omitempty" xml:"ChanelType,omitempty"`
	// The value of the alert notification method.
	//
	// > This parameter is required only when `ChanelType` is set to `Mail`.
	//
	// example:
	//
	// Alice@example.com
	ChanelValue *string `json:"ChanelValue,omitempty" xml:"ChanelValue,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// Alice
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The current page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContactListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactListRequest) GetChanelType() *string {
	return s.ChanelType
}

func (s *DescribeContactListRequest) GetChanelValue() *string {
	return s.ChanelValue
}

func (s *DescribeContactListRequest) GetContactName() *string {
	return s.ContactName
}

func (s *DescribeContactListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeContactListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContactListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContactListRequest) SetChanelType(v string) *DescribeContactListRequest {
	s.ChanelType = &v
	return s
}

func (s *DescribeContactListRequest) SetChanelValue(v string) *DescribeContactListRequest {
	s.ChanelValue = &v
	return s
}

func (s *DescribeContactListRequest) SetContactName(v string) *DescribeContactListRequest {
	s.ContactName = &v
	return s
}

func (s *DescribeContactListRequest) SetPageNumber(v int32) *DescribeContactListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeContactListRequest) SetPageSize(v int32) *DescribeContactListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContactListRequest) SetRegionId(v string) *DescribeContactListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContactListRequest) Validate() error {
	return dara.Validate(s)
}
