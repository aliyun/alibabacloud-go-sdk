// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBuyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetUserBuyStatusResponseBody
	GetCode() *int32
	SetData(v *GetUserBuyStatusResponseBodyData) *GetUserBuyStatusResponseBody
	GetData() *GetUserBuyStatusResponseBodyData
	SetMsg(v string) *GetUserBuyStatusResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetUserBuyStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserBuyStatusResponseBody
	GetSuccess() *bool
}

type GetUserBuyStatusResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetUserBuyStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID assigned by the backend to uniquely identify the request. This ID can be used to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The success flag.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserBuyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserBuyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetUserBuyStatusResponseBody) GetData() *GetUserBuyStatusResponseBodyData {
	return s.Data
}

func (s *GetUserBuyStatusResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetUserBuyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserBuyStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserBuyStatusResponseBody) SetCode(v int32) *GetUserBuyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetData(v *GetUserBuyStatusResponseBodyData) *GetUserBuyStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetMsg(v string) *GetUserBuyStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetRequestId(v string) *GetUserBuyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetSuccess(v bool) *GetUserBuyStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserBuyStatusResponseBodyData struct {
	// Bid。
	//
	// example:
	//
	// 26842
	Bid *int64 `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// Indicates whether the commodity is activated on Alibaba Cloud.
	//
	// example:
	//
	// True
	Buy *bool `json:"Buy,omitempty" xml:"Buy,omitempty"`
	// Indicates whether there is an overdue payment.
	//
	// example:
	//
	// False
	Indebt *bool `json:"Indebt,omitempty" xml:"Indebt,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The label.
	//
	// example:
	//
	// bailian
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetUserBuyStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserBuyStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusResponseBodyData) GetBid() *int64 {
	return s.Bid
}

func (s *GetUserBuyStatusResponseBodyData) GetBuy() *bool {
	return s.Buy
}

func (s *GetUserBuyStatusResponseBodyData) GetIndebt() *bool {
	return s.Indebt
}

func (s *GetUserBuyStatusResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserBuyStatusResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *GetUserBuyStatusResponseBodyData) SetBid(v int64) *GetUserBuyStatusResponseBodyData {
	s.Bid = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetBuy(v bool) *GetUserBuyStatusResponseBodyData {
	s.Buy = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetIndebt(v bool) *GetUserBuyStatusResponseBodyData {
	s.Indebt = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetInstanceId(v string) *GetUserBuyStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetTag(v string) *GetUserBuyStatusResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
