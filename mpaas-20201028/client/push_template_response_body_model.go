// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushTemplateResponseBodyPushResult) *PushTemplateResponseBody
	GetPushResult() *PushTemplateResponseBodyPushResult
	SetRequestId(v string) *PushTemplateResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushTemplateResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushTemplateResponseBody
	GetResultMessage() *string
}

type PushTemplateResponseBody struct {
	PushResult    *PushTemplateResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *PushTemplateResponseBody) GetPushResult() *PushTemplateResponseBodyPushResult {
	return s.PushResult
}

func (s *PushTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushTemplateResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushTemplateResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushTemplateResponseBody) SetPushResult(v *PushTemplateResponseBodyPushResult) *PushTemplateResponseBody {
	s.PushResult = v
	return s
}

func (s *PushTemplateResponseBody) SetRequestId(v string) *PushTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushTemplateResponseBody) SetResultCode(v string) *PushTemplateResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushTemplateResponseBody) SetResultMessage(v string) *PushTemplateResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type PushTemplateResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushTemplateResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushTemplateResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushTemplateResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushTemplateResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushTemplateResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushTemplateResponseBodyPushResult) SetData(v string) *PushTemplateResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushTemplateResponseBodyPushResult) SetResultMsg(v string) *PushTemplateResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushTemplateResponseBodyPushResult) SetSuccess(v bool) *PushTemplateResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushTemplateResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
