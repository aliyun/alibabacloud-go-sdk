// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushBindResponseBodyPushResult) *PushBindResponseBody
	GetPushResult() *PushBindResponseBodyPushResult
	SetRequestId(v string) *PushBindResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushBindResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushBindResponseBody
	GetResultMessage() *string
}

type PushBindResponseBody struct {
	PushResult    *PushBindResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                         `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushBindResponseBody) GoString() string {
	return s.String()
}

func (s *PushBindResponseBody) GetPushResult() *PushBindResponseBodyPushResult {
	return s.PushResult
}

func (s *PushBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushBindResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushBindResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushBindResponseBody) SetPushResult(v *PushBindResponseBodyPushResult) *PushBindResponseBody {
	s.PushResult = v
	return s
}

func (s *PushBindResponseBody) SetRequestId(v string) *PushBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushBindResponseBody) SetResultCode(v string) *PushBindResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushBindResponseBody) SetResultMessage(v string) *PushBindResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushBindResponseBody) Validate() error {
	return dara.Validate(s)
}

type PushBindResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushBindResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushBindResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushBindResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushBindResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushBindResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushBindResponseBodyPushResult) SetData(v string) *PushBindResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushBindResponseBodyPushResult) SetResultMsg(v string) *PushBindResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushBindResponseBodyPushResult) SetSuccess(v bool) *PushBindResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushBindResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
