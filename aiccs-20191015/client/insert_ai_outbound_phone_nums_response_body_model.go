// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertAiOutboundPhoneNumsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsertAiOutboundPhoneNumsResponseBody
	GetCode() *string
	SetData(v *InsertAiOutboundPhoneNumsResponseBodyData) *InsertAiOutboundPhoneNumsResponseBody
	GetData() *InsertAiOutboundPhoneNumsResponseBodyData
	SetMessage(v string) *InsertAiOutboundPhoneNumsResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertAiOutboundPhoneNumsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsertAiOutboundPhoneNumsResponseBody
	GetSuccess() *bool
}

type InsertAiOutboundPhoneNumsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Import result.
	Data *InsertAiOutboundPhoneNumsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
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

func (s InsertAiOutboundPhoneNumsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsResponseBody) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsertAiOutboundPhoneNumsResponseBody) GetData() *InsertAiOutboundPhoneNumsResponseBodyData {
	return s.Data
}

func (s *InsertAiOutboundPhoneNumsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertAiOutboundPhoneNumsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertAiOutboundPhoneNumsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsertAiOutboundPhoneNumsResponseBody) SetCode(v string) *InsertAiOutboundPhoneNumsResponseBody {
	s.Code = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBody) SetData(v *InsertAiOutboundPhoneNumsResponseBodyData) *InsertAiOutboundPhoneNumsResponseBody {
	s.Data = v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBody) SetMessage(v string) *InsertAiOutboundPhoneNumsResponseBody {
	s.Message = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBody) SetRequestId(v string) *InsertAiOutboundPhoneNumsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBody) SetSuccess(v bool) *InsertAiOutboundPhoneNumsResponseBody {
	s.Success = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertAiOutboundPhoneNumsResponseBodyData struct {
	// Details of failed numbers.
	FailInfo []*InsertAiOutboundPhoneNumsResponseBodyDataFailInfo `json:"FailInfo,omitempty" xml:"FailInfo,omitempty" type:"Repeated"`
	// Number of successfully imported entries.
	//
	// example:
	//
	// 7
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// Total number of imported entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s InsertAiOutboundPhoneNumsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) GetFailInfo() []*InsertAiOutboundPhoneNumsResponseBodyDataFailInfo {
	return s.FailInfo
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) SetFailInfo(v []*InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) *InsertAiOutboundPhoneNumsResponseBodyData {
	s.FailInfo = v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) SetSuccessCount(v int32) *InsertAiOutboundPhoneNumsResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) SetTotalCount(v int32) *InsertAiOutboundPhoneNumsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBodyData) Validate() error {
	if s.FailInfo != nil {
		for _, item := range s.FailInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsertAiOutboundPhoneNumsResponseBodyDataFailInfo struct {
	// Custom business information.
	//
	// example:
	//
	// xxxx
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	// Description of the failure reason.
	//
	// example:
	//
	// 号码格式异常
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The callee number for outbound calls.
	//
	// example:
	//
	// 150****0000
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) GetBizData() *string {
	return s.BizData
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) GetMsg() *string {
	return s.Msg
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) SetBizData(v string) *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo {
	s.BizData = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) SetMsg(v string) *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo {
	s.Msg = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) SetPhoneNum(v string) *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo {
	s.PhoneNum = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponseBodyDataFailInfo) Validate() error {
	return dara.Validate(s)
}
