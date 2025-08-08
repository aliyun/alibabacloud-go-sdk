// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPushSchedulerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelPushSchedulerResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CancelPushSchedulerResponseBody
	GetResultCode() *string
	SetResultContent(v string) *CancelPushSchedulerResponseBody
	GetResultContent() *string
	SetResultMessage(v string) *CancelPushSchedulerResponseBody
	GetResultMessage() *string
}

type CancelPushSchedulerResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CancelPushSchedulerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPushSchedulerResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPushSchedulerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPushSchedulerResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CancelPushSchedulerResponseBody) GetResultContent() *string {
	return s.ResultContent
}

func (s *CancelPushSchedulerResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CancelPushSchedulerResponseBody) SetRequestId(v string) *CancelPushSchedulerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPushSchedulerResponseBody) SetResultCode(v string) *CancelPushSchedulerResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CancelPushSchedulerResponseBody) SetResultContent(v string) *CancelPushSchedulerResponseBody {
	s.ResultContent = &v
	return s
}

func (s *CancelPushSchedulerResponseBody) SetResultMessage(v string) *CancelPushSchedulerResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CancelPushSchedulerResponseBody) Validate() error {
	return dara.Validate(s)
}
