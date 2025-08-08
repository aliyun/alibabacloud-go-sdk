// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMsaDiffResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunMsaDiffResponseBody
	GetRequestId() *string
	SetResultCode(v string) *RunMsaDiffResponseBody
	GetResultCode() *string
	SetResultContent(v *RunMsaDiffResponseBodyResultContent) *RunMsaDiffResponseBody
	GetResultContent() *RunMsaDiffResponseBodyResultContent
	SetResultMessage(v string) *RunMsaDiffResponseBody
	GetResultMessage() *string
}

type RunMsaDiffResponseBody struct {
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *RunMsaDiffResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s RunMsaDiffResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunMsaDiffResponseBody) GoString() string {
	return s.String()
}

func (s *RunMsaDiffResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunMsaDiffResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *RunMsaDiffResponseBody) GetResultContent() *RunMsaDiffResponseBodyResultContent {
	return s.ResultContent
}

func (s *RunMsaDiffResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *RunMsaDiffResponseBody) SetRequestId(v string) *RunMsaDiffResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunMsaDiffResponseBody) SetResultCode(v string) *RunMsaDiffResponseBody {
	s.ResultCode = &v
	return s
}

func (s *RunMsaDiffResponseBody) SetResultContent(v *RunMsaDiffResponseBodyResultContent) *RunMsaDiffResponseBody {
	s.ResultContent = v
	return s
}

func (s *RunMsaDiffResponseBody) SetResultMessage(v string) *RunMsaDiffResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *RunMsaDiffResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunMsaDiffResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunMsaDiffResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s RunMsaDiffResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *RunMsaDiffResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *RunMsaDiffResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *RunMsaDiffResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *RunMsaDiffResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *RunMsaDiffResponseBodyResultContent) SetCode(v string) *RunMsaDiffResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *RunMsaDiffResponseBodyResultContent) SetData(v string) *RunMsaDiffResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *RunMsaDiffResponseBodyResultContent) SetMessage(v string) *RunMsaDiffResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *RunMsaDiffResponseBodyResultContent) SetSuccess(v bool) *RunMsaDiffResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *RunMsaDiffResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
