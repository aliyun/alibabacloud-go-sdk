// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskBizDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiOutboundTaskBizDataResponseBody
	GetCode() *string
	SetData(v *GetAiOutboundTaskBizDataResponseBodyData) *GetAiOutboundTaskBizDataResponseBody
	GetData() *GetAiOutboundTaskBizDataResponseBodyData
	SetMessage(v string) *GetAiOutboundTaskBizDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiOutboundTaskBizDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiOutboundTaskBizDataResponseBody
	GetSuccess() *bool
}

type GetAiOutboundTaskBizDataResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business information associated with this call.
	Data *GetAiOutboundTaskBizDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// EE338D98-9BD3-4413-B165
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

func (s GetAiOutboundTaskBizDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskBizDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskBizDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiOutboundTaskBizDataResponseBody) GetData() *GetAiOutboundTaskBizDataResponseBodyData {
	return s.Data
}

func (s *GetAiOutboundTaskBizDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiOutboundTaskBizDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiOutboundTaskBizDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiOutboundTaskBizDataResponseBody) SetCode(v string) *GetAiOutboundTaskBizDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBody) SetData(v *GetAiOutboundTaskBizDataResponseBodyData) *GetAiOutboundTaskBizDataResponseBody {
	s.Data = v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBody) SetMessage(v string) *GetAiOutboundTaskBizDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBody) SetRequestId(v string) *GetAiOutboundTaskBizDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBody) SetSuccess(v bool) *GetAiOutboundTaskBizDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskBizDataResponseBodyData struct {
	// Custom business information.
	//
	// example:
	//
	// {"customer":123}
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	// The Activity ID associated with this outbound call.
	//
	// example:
	//
	// 123
	CaseId *int64 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// The outbound phone number.
	//
	// example:
	//
	// 158****0000
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAiOutboundTaskBizDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskBizDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) GetBizData() *string {
	return s.BizData
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) GetCaseId() *int64 {
	return s.CaseId
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) SetBizData(v string) *GetAiOutboundTaskBizDataResponseBodyData {
	s.BizData = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) SetCaseId(v int64) *GetAiOutboundTaskBizDataResponseBodyData {
	s.CaseId = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) SetPhoneNum(v string) *GetAiOutboundTaskBizDataResponseBodyData {
	s.PhoneNum = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) SetTaskId(v int64) *GetAiOutboundTaskBizDataResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskBizDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
