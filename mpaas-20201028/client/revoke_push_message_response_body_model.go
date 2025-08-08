// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePushMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *RevokePushMessageResponseBodyPushResult) *RevokePushMessageResponseBody
	GetPushResult() *RevokePushMessageResponseBodyPushResult
	SetRequestId(v string) *RevokePushMessageResponseBody
	GetRequestId() *string
	SetResultCode(v string) *RevokePushMessageResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *RevokePushMessageResponseBody
	GetResultMessage() *string
}

type RevokePushMessageResponseBody struct {
	PushResult    *RevokePushMessageResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s RevokePushMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokePushMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RevokePushMessageResponseBody) GetPushResult() *RevokePushMessageResponseBodyPushResult {
	return s.PushResult
}

func (s *RevokePushMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokePushMessageResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *RevokePushMessageResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *RevokePushMessageResponseBody) SetPushResult(v *RevokePushMessageResponseBodyPushResult) *RevokePushMessageResponseBody {
	s.PushResult = v
	return s
}

func (s *RevokePushMessageResponseBody) SetRequestId(v string) *RevokePushMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokePushMessageResponseBody) SetResultCode(v string) *RevokePushMessageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *RevokePushMessageResponseBody) SetResultMessage(v string) *RevokePushMessageResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *RevokePushMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

type RevokePushMessageResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokePushMessageResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s RevokePushMessageResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *RevokePushMessageResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *RevokePushMessageResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *RevokePushMessageResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *RevokePushMessageResponseBodyPushResult) SetData(v string) *RevokePushMessageResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *RevokePushMessageResponseBodyPushResult) SetResultMsg(v string) *RevokePushMessageResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *RevokePushMessageResponseBodyPushResult) SetSuccess(v bool) *RevokePushMessageResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *RevokePushMessageResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
