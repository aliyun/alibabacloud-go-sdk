// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMultipleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushMultipleResponseBodyPushResult) *PushMultipleResponseBody
	GetPushResult() *PushMultipleResponseBodyPushResult
	SetRequestId(v string) *PushMultipleResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushMultipleResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushMultipleResponseBody
	GetResultMessage() *string
}

type PushMultipleResponseBody struct {
	PushResult    *PushMultipleResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushMultipleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleResponseBody) GoString() string {
	return s.String()
}

func (s *PushMultipleResponseBody) GetPushResult() *PushMultipleResponseBodyPushResult {
	return s.PushResult
}

func (s *PushMultipleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushMultipleResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushMultipleResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushMultipleResponseBody) SetPushResult(v *PushMultipleResponseBodyPushResult) *PushMultipleResponseBody {
	s.PushResult = v
	return s
}

func (s *PushMultipleResponseBody) SetRequestId(v string) *PushMultipleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMultipleResponseBody) SetResultCode(v string) *PushMultipleResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushMultipleResponseBody) SetResultMessage(v string) *PushMultipleResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushMultipleResponseBody) Validate() error {
	return dara.Validate(s)
}

type PushMultipleResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushMultipleResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushMultipleResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushMultipleResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushMultipleResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushMultipleResponseBodyPushResult) SetData(v string) *PushMultipleResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushMultipleResponseBodyPushResult) SetResultMsg(v string) *PushMultipleResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushMultipleResponseBodyPushResult) SetSuccess(v bool) *PushMultipleResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushMultipleResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
