// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushUnBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushUnBindResponseBodyPushResult) *PushUnBindResponseBody
	GetPushResult() *PushUnBindResponseBodyPushResult
	SetRequestId(v string) *PushUnBindResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushUnBindResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushUnBindResponseBody
	GetResultMessage() *string
}

type PushUnBindResponseBody struct {
	PushResult    *PushUnBindResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushUnBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushUnBindResponseBody) GoString() string {
	return s.String()
}

func (s *PushUnBindResponseBody) GetPushResult() *PushUnBindResponseBodyPushResult {
	return s.PushResult
}

func (s *PushUnBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushUnBindResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushUnBindResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushUnBindResponseBody) SetPushResult(v *PushUnBindResponseBodyPushResult) *PushUnBindResponseBody {
	s.PushResult = v
	return s
}

func (s *PushUnBindResponseBody) SetRequestId(v string) *PushUnBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushUnBindResponseBody) SetResultCode(v string) *PushUnBindResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushUnBindResponseBody) SetResultMessage(v string) *PushUnBindResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushUnBindResponseBody) Validate() error {
	if s.PushResult != nil {
		if err := s.PushResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushUnBindResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushUnBindResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushUnBindResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushUnBindResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushUnBindResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushUnBindResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushUnBindResponseBodyPushResult) SetData(v string) *PushUnBindResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushUnBindResponseBodyPushResult) SetResultMsg(v string) *PushUnBindResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushUnBindResponseBodyPushResult) SetSuccess(v bool) *PushUnBindResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushUnBindResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
