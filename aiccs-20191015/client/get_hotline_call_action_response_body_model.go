// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineCallActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineCallActionResponseBody
	GetCode() *string
	SetData(v *GetHotlineCallActionResponseBodyData) *GetHotlineCallActionResponseBody
	GetData() *GetHotlineCallActionResponseBodyData
	SetMessage(v string) *GetHotlineCallActionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineCallActionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotlineCallActionResponseBody
	GetSuccess() *bool
}

type GetHotlineCallActionResponseBody struct {
	// Status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *GetHotlineCallActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE339D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineCallActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineCallActionResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineCallActionResponseBody) GetData() *GetHotlineCallActionResponseBodyData {
	return s.Data
}

func (s *GetHotlineCallActionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineCallActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineCallActionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotlineCallActionResponseBody) SetCode(v string) *GetHotlineCallActionResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetData(v *GetHotlineCallActionResponseBodyData) *GetHotlineCallActionResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetMessage(v string) *GetHotlineCallActionResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetRequestId(v string) *GetHotlineCallActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) SetSuccess(v bool) *GetHotlineCallActionResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineCallActionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotlineCallActionResponseBodyData struct {
	// Customer ID.
	//
	// example:
	//
	// 8999****
	ActionId *int64 `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// 2367****
	BuId *int64 `json:"BuId,omitempty" xml:"BuId,omitempty"`
	// Callout ID.
	//
	// example:
	//
	// 1122****
	CalloutId *int64 `json:"CalloutId,omitempty" xml:"CalloutId,omitempty"`
	// Call name.
	//
	// example:
	//
	// 王XX
	CalloutName *string `json:"CalloutName,omitempty" xml:"CalloutName,omitempty"`
	// Ticket ID.
	//
	// example:
	//
	// 1138902****
	CaseId *int64 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// Channel ID.
	//
	// example:
	//
	// 2377****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Channel Type. Valid values:
	//
	// - **1**: Hotline.
	//
	// - **2**: Online.
	//
	// example:
	//
	// 2
	ChannelType *int64 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Department ID.
	//
	// example:
	//
	// 1223****
	DepId *int64 `json:"DepId,omitempty" xml:"DepId,omitempty"`
	// Indicates whether the call is transferred.
	//
	// example:
	//
	// true
	IsTransfer *string `json:"IsTransfer,omitempty" xml:"IsTransfer,omitempty"`
	// Membership ID.
	//
	// example:
	//
	// 7856876****
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// Membership List.
	//
	// example:
	//
	// 8900****
	MemberList *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
	// Membership name.
	//
	// example:
	//
	// 匿名会员
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// 1332****
	ServicerId *int64 `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	// Agent name.
	//
	// example:
	//
	// XX测试
	ServicerName *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	// Sub-touch ID.
	//
	// example:
	//
	// 3423****
	SubTouchId *int64 `json:"SubTouchId,omitempty" xml:"SubTouchId,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 12345****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Touch ID.
	//
	// example:
	//
	// 2235****
	TouchId *int64 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
}

func (s GetHotlineCallActionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineCallActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionResponseBodyData) GetActionId() *int64 {
	return s.ActionId
}

func (s *GetHotlineCallActionResponseBodyData) GetBuId() *int64 {
	return s.BuId
}

func (s *GetHotlineCallActionResponseBodyData) GetCalloutId() *int64 {
	return s.CalloutId
}

func (s *GetHotlineCallActionResponseBodyData) GetCalloutName() *string {
	return s.CalloutName
}

func (s *GetHotlineCallActionResponseBodyData) GetCaseId() *int64 {
	return s.CaseId
}

func (s *GetHotlineCallActionResponseBodyData) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetHotlineCallActionResponseBodyData) GetChannelType() *int64 {
	return s.ChannelType
}

func (s *GetHotlineCallActionResponseBodyData) GetDepId() *int64 {
	return s.DepId
}

func (s *GetHotlineCallActionResponseBodyData) GetIsTransfer() *string {
	return s.IsTransfer
}

func (s *GetHotlineCallActionResponseBodyData) GetMemberId() *int64 {
	return s.MemberId
}

func (s *GetHotlineCallActionResponseBodyData) GetMemberList() *string {
	return s.MemberList
}

func (s *GetHotlineCallActionResponseBodyData) GetMemberName() *string {
	return s.MemberName
}

func (s *GetHotlineCallActionResponseBodyData) GetServicerId() *int64 {
	return s.ServicerId
}

func (s *GetHotlineCallActionResponseBodyData) GetServicerName() *string {
	return s.ServicerName
}

func (s *GetHotlineCallActionResponseBodyData) GetSubTouchId() *int64 {
	return s.SubTouchId
}

func (s *GetHotlineCallActionResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetHotlineCallActionResponseBodyData) GetTouchId() *int64 {
	return s.TouchId
}

func (s *GetHotlineCallActionResponseBodyData) SetActionId(v int64) *GetHotlineCallActionResponseBodyData {
	s.ActionId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetBuId(v int64) *GetHotlineCallActionResponseBodyData {
	s.BuId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetCalloutId(v int64) *GetHotlineCallActionResponseBodyData {
	s.CalloutId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetCalloutName(v string) *GetHotlineCallActionResponseBodyData {
	s.CalloutName = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetCaseId(v int64) *GetHotlineCallActionResponseBodyData {
	s.CaseId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetChannelId(v string) *GetHotlineCallActionResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetChannelType(v int64) *GetHotlineCallActionResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetDepId(v int64) *GetHotlineCallActionResponseBodyData {
	s.DepId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetIsTransfer(v string) *GetHotlineCallActionResponseBodyData {
	s.IsTransfer = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetMemberId(v int64) *GetHotlineCallActionResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetMemberList(v string) *GetHotlineCallActionResponseBodyData {
	s.MemberList = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetMemberName(v string) *GetHotlineCallActionResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetServicerId(v int64) *GetHotlineCallActionResponseBodyData {
	s.ServicerId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetServicerName(v string) *GetHotlineCallActionResponseBodyData {
	s.ServicerName = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetSubTouchId(v int64) *GetHotlineCallActionResponseBodyData {
	s.SubTouchId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetTaskId(v int64) *GetHotlineCallActionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) SetTouchId(v int64) *GetHotlineCallActionResponseBodyData {
	s.TouchId = &v
	return s
}

func (s *GetHotlineCallActionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
