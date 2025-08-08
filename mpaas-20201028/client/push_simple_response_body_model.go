// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushSimpleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushSimpleResponseBodyPushResult) *PushSimpleResponseBody
	GetPushResult() *PushSimpleResponseBodyPushResult
	SetRequestId(v string) *PushSimpleResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushSimpleResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushSimpleResponseBody
	GetResultMessage() *string
}

type PushSimpleResponseBody struct {
	PushResult    *PushSimpleResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushSimpleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushSimpleResponseBody) GoString() string {
	return s.String()
}

func (s *PushSimpleResponseBody) GetPushResult() *PushSimpleResponseBodyPushResult {
	return s.PushResult
}

func (s *PushSimpleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushSimpleResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushSimpleResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushSimpleResponseBody) SetPushResult(v *PushSimpleResponseBodyPushResult) *PushSimpleResponseBody {
	s.PushResult = v
	return s
}

func (s *PushSimpleResponseBody) SetRequestId(v string) *PushSimpleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushSimpleResponseBody) SetResultCode(v string) *PushSimpleResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushSimpleResponseBody) SetResultMessage(v string) *PushSimpleResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushSimpleResponseBody) Validate() error {
	return dara.Validate(s)
}

type PushSimpleResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushSimpleResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushSimpleResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushSimpleResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushSimpleResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushSimpleResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushSimpleResponseBodyPushResult) SetData(v string) *PushSimpleResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushSimpleResponseBodyPushResult) SetResultMsg(v string) *PushSimpleResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushSimpleResponseBodyPushResult) SetSuccess(v bool) *PushSimpleResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushSimpleResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
