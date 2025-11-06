// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBroadcastResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushBroadcastResponseBodyPushResult) *PushBroadcastResponseBody
	GetPushResult() *PushBroadcastResponseBodyPushResult
	SetRequestId(v string) *PushBroadcastResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushBroadcastResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushBroadcastResponseBody
	GetResultMessage() *string
}

type PushBroadcastResponseBody struct {
	PushResult    *PushBroadcastResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushBroadcastResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushBroadcastResponseBody) GoString() string {
	return s.String()
}

func (s *PushBroadcastResponseBody) GetPushResult() *PushBroadcastResponseBodyPushResult {
	return s.PushResult
}

func (s *PushBroadcastResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushBroadcastResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushBroadcastResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushBroadcastResponseBody) SetPushResult(v *PushBroadcastResponseBodyPushResult) *PushBroadcastResponseBody {
	s.PushResult = v
	return s
}

func (s *PushBroadcastResponseBody) SetRequestId(v string) *PushBroadcastResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushBroadcastResponseBody) SetResultCode(v string) *PushBroadcastResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushBroadcastResponseBody) SetResultMessage(v string) *PushBroadcastResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushBroadcastResponseBody) Validate() error {
	if s.PushResult != nil {
		if err := s.PushResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushBroadcastResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushBroadcastResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushBroadcastResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushBroadcastResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushBroadcastResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushBroadcastResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushBroadcastResponseBodyPushResult) SetData(v string) *PushBroadcastResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushBroadcastResponseBodyPushResult) SetResultMsg(v string) *PushBroadcastResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushBroadcastResponseBodyPushResult) SetSuccess(v bool) *PushBroadcastResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushBroadcastResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
