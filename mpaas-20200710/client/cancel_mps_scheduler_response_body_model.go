// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMpsSchedulerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelMpsSchedulerResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CancelMpsSchedulerResponseBody
	GetResultCode() *string
	SetResultContent(v string) *CancelMpsSchedulerResponseBody
	GetResultContent() *string
	SetResultMessage(v string) *CancelMpsSchedulerResponseBody
	GetResultMessage() *string
}

type CancelMpsSchedulerResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CancelMpsSchedulerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelMpsSchedulerResponseBody) GoString() string {
	return s.String()
}

func (s *CancelMpsSchedulerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelMpsSchedulerResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CancelMpsSchedulerResponseBody) GetResultContent() *string {
	return s.ResultContent
}

func (s *CancelMpsSchedulerResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CancelMpsSchedulerResponseBody) SetRequestId(v string) *CancelMpsSchedulerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelMpsSchedulerResponseBody) SetResultCode(v string) *CancelMpsSchedulerResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CancelMpsSchedulerResponseBody) SetResultContent(v string) *CancelMpsSchedulerResponseBody {
	s.ResultContent = &v
	return s
}

func (s *CancelMpsSchedulerResponseBody) SetResultMessage(v string) *CancelMpsSchedulerResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CancelMpsSchedulerResponseBody) Validate() error {
	return dara.Validate(s)
}
