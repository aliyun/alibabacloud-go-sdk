// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePushTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *RevokePushTaskResponseBodyPushResult) *RevokePushTaskResponseBody
	GetPushResult() *RevokePushTaskResponseBodyPushResult
	SetRequestId(v string) *RevokePushTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *RevokePushTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *RevokePushTaskResponseBody
	GetResultMessage() *string
}

type RevokePushTaskResponseBody struct {
	PushResult    *RevokePushTaskResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                               `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                               `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s RevokePushTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokePushTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RevokePushTaskResponseBody) GetPushResult() *RevokePushTaskResponseBodyPushResult {
	return s.PushResult
}

func (s *RevokePushTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokePushTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *RevokePushTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *RevokePushTaskResponseBody) SetPushResult(v *RevokePushTaskResponseBodyPushResult) *RevokePushTaskResponseBody {
	s.PushResult = v
	return s
}

func (s *RevokePushTaskResponseBody) SetRequestId(v string) *RevokePushTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokePushTaskResponseBody) SetResultCode(v string) *RevokePushTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *RevokePushTaskResponseBody) SetResultMessage(v string) *RevokePushTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *RevokePushTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type RevokePushTaskResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokePushTaskResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s RevokePushTaskResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *RevokePushTaskResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *RevokePushTaskResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *RevokePushTaskResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *RevokePushTaskResponseBodyPushResult) SetData(v string) *RevokePushTaskResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *RevokePushTaskResponseBodyPushResult) SetResultMsg(v string) *RevokePushTaskResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *RevokePushTaskResponseBodyPushResult) SetSuccess(v bool) *RevokePushTaskResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *RevokePushTaskResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
