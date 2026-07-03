// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeUserBuyStatusResponseBodyData) *DescribeUserBuyStatusResponseBody
	GetData() *DescribeUserBuyStatusResponseBodyData
	SetRequestId(v string) *DescribeUserBuyStatusResponseBody
	GetRequestId() *string
}

type DescribeUserBuyStatusResponseBody struct {
	// The response parameters.
	Data *DescribeUserBuyStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 81D8EC0C-0804-51AD-8C38-17ED0BC74892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserBuyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusResponseBody) GetData() *DescribeUserBuyStatusResponseBodyData {
	return s.Data
}

func (s *DescribeUserBuyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBuyStatusResponseBody) SetData(v *DescribeUserBuyStatusResponseBodyData) *DescribeUserBuyStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserBuyStatusResponseBody) SetRequestId(v string) *DescribeUserBuyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserBuyStatusResponseBodyData struct {
	// Indicates whether the current account can perform operations on threat analysis orders. Valid values:
	//
	// - true: The account can purchase, upgrade, or downgrade threat analysis.
	//
	// - false: The account cannot perform operations on threat analysis orders.
	//
	// example:
	//
	// true
	CanBuy *bool `json:"CanBuy,omitempty" xml:"CanBuy,omitempty"`
	// The purchased capacity of Simple Log Service (SLS) for threat analysis. Unit: GB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The number of days before threat analysis expires.
	//
	// example:
	//
	// 3
	DurationDays *int64 `json:"DurationDays,omitempty" xml:"DurationDays,omitempty"`
	// The expiration time of threat analysis. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1669823999000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud account that purchased threat analysis.
	//
	// example:
	//
	// 123XXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The name of the Alibaba Cloud account that purchased threat analysis.
	//
	// example:
	//
	// sas_account_xxx
	MainUserName *string `json:"MainUserName,omitempty" xml:"MainUserName,omitempty"`
	// The ID of the master account of the resource directory.
	//
	// example:
	//
	// 123XXXXXX
	MasterUserId *int64 `json:"MasterUserId,omitempty" xml:"MasterUserId,omitempty"`
	// The display name of the master account of the resource directory.
	//
	// example:
	//
	// rd_master_xxx
	MasterUserName *string `json:"MasterUserName,omitempty" xml:"MasterUserName,omitempty"`
	// The type of the current order.
	//
	// - 0: The order includes threat analysis traffic and capacity.
	//
	// - 1: The order includes only threat analysis capacity.
	//
	// example:
	//
	// 1
	RdOrder *int32 `json:"RdOrder,omitempty" xml:"RdOrder,omitempty"`
	// The ID of the Security Center instance.
	//
	// example:
	//
	// sas-instance-xxxxx
	SasInstanceId *string `json:"SasInstanceId,omitempty" xml:"SasInstanceId,omitempty"`
	// The ID of the currently logged-on Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The name of the currently logged-on Alibaba Cloud account.
	//
	// example:
	//
	// sas_account_xxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s DescribeUserBuyStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusResponseBodyData) GetCanBuy() *bool {
	return s.CanBuy
}

func (s *DescribeUserBuyStatusResponseBodyData) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DescribeUserBuyStatusResponseBodyData) GetDurationDays() *int64 {
	return s.DurationDays
}

func (s *DescribeUserBuyStatusResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUserBuyStatusResponseBodyData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *DescribeUserBuyStatusResponseBodyData) GetMainUserName() *string {
	return s.MainUserName
}

func (s *DescribeUserBuyStatusResponseBodyData) GetMasterUserId() *int64 {
	return s.MasterUserId
}

func (s *DescribeUserBuyStatusResponseBodyData) GetMasterUserName() *string {
	return s.MasterUserName
}

func (s *DescribeUserBuyStatusResponseBodyData) GetRdOrder() *int32 {
	return s.RdOrder
}

func (s *DescribeUserBuyStatusResponseBodyData) GetSasInstanceId() *string {
	return s.SasInstanceId
}

func (s *DescribeUserBuyStatusResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeUserBuyStatusResponseBodyData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *DescribeUserBuyStatusResponseBodyData) SetCanBuy(v bool) *DescribeUserBuyStatusResponseBodyData {
	s.CanBuy = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetCapacity(v int32) *DescribeUserBuyStatusResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetDurationDays(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.DurationDays = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetEndTime(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMainUserId(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMainUserName(v string) *DescribeUserBuyStatusResponseBodyData {
	s.MainUserName = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMasterUserId(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.MasterUserId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMasterUserName(v string) *DescribeUserBuyStatusResponseBodyData {
	s.MasterUserName = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetRdOrder(v int32) *DescribeUserBuyStatusResponseBodyData {
	s.RdOrder = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetSasInstanceId(v string) *DescribeUserBuyStatusResponseBodyData {
	s.SasInstanceId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetSubUserId(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetSubUserName(v string) *DescribeUserBuyStatusResponseBodyData {
	s.SubUserName = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
