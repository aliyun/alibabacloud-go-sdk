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
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUserBuyStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type GetUserBuyStatusResponseBodyData struct {
	// Bid。
	//
	// example:
	//
	// 26842
	Bid *int64 `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// True
	Buy *bool `json:"Buy,omitempty" xml:"Buy,omitempty"`
	// example:
	//
	// False
	Indebt *bool   `json:"Indebt,omitempty" xml:"Indebt,omitempty"`
	Tag    *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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

func (s *GetUserBuyStatusResponseBodyData) SetTag(v string) *GetUserBuyStatusResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
