// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResult(v *PushReportResponseBodyPushResult) *PushReportResponseBody
	GetPushResult() *PushReportResponseBodyPushResult
	SetRequestId(v string) *PushReportResponseBody
	GetRequestId() *string
	SetResultCode(v string) *PushReportResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *PushReportResponseBody
	GetResultMessage() *string
}

type PushReportResponseBody struct {
	PushResult    *PushReportResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Struct"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushReportResponseBody) GoString() string {
	return s.String()
}

func (s *PushReportResponseBody) GetPushResult() *PushReportResponseBodyPushResult {
	return s.PushResult
}

func (s *PushReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushReportResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *PushReportResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *PushReportResponseBody) SetPushResult(v *PushReportResponseBodyPushResult) *PushReportResponseBody {
	s.PushResult = v
	return s
}

func (s *PushReportResponseBody) SetRequestId(v string) *PushReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushReportResponseBody) SetResultCode(v string) *PushReportResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushReportResponseBody) SetResultMessage(v string) *PushReportResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *PushReportResponseBody) Validate() error {
	if s.PushResult != nil {
		if err := s.PushResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushReportResponseBodyPushResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushReportResponseBodyPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushReportResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *PushReportResponseBodyPushResult) GetData() *string {
	return s.Data
}

func (s *PushReportResponseBodyPushResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *PushReportResponseBodyPushResult) GetSuccess() *bool {
	return s.Success
}

func (s *PushReportResponseBodyPushResult) SetData(v string) *PushReportResponseBodyPushResult {
	s.Data = &v
	return s
}

func (s *PushReportResponseBodyPushResult) SetResultMsg(v string) *PushReportResponseBodyPushResult {
	s.ResultMsg = &v
	return s
}

func (s *PushReportResponseBodyPushResult) SetSuccess(v bool) *PushReportResponseBodyPushResult {
	s.Success = &v
	return s
}

func (s *PushReportResponseBodyPushResult) Validate() error {
	return dara.Validate(s)
}
