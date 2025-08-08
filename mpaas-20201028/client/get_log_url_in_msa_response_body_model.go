// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogUrlInMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLogUrlInMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetLogUrlInMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *GetLogUrlInMsaResponseBodyResultContent) *GetLogUrlInMsaResponseBody
	GetResultContent() *GetLogUrlInMsaResponseBodyResultContent
	SetResultMessage(v string) *GetLogUrlInMsaResponseBody
	GetResultMessage() *string
}

type GetLogUrlInMsaResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetLogUrlInMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetLogUrlInMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogUrlInMsaResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogUrlInMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogUrlInMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetLogUrlInMsaResponseBody) GetResultContent() *GetLogUrlInMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetLogUrlInMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetLogUrlInMsaResponseBody) SetRequestId(v string) *GetLogUrlInMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogUrlInMsaResponseBody) SetResultCode(v string) *GetLogUrlInMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetLogUrlInMsaResponseBody) SetResultContent(v *GetLogUrlInMsaResponseBodyResultContent) *GetLogUrlInMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetLogUrlInMsaResponseBody) SetResultMessage(v string) *GetLogUrlInMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetLogUrlInMsaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLogUrlInMsaResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogUrlInMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetLogUrlInMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetLogUrlInMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *GetLogUrlInMsaResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *GetLogUrlInMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *GetLogUrlInMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *GetLogUrlInMsaResponseBodyResultContent) SetCode(v string) *GetLogUrlInMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *GetLogUrlInMsaResponseBodyResultContent) SetData(v string) *GetLogUrlInMsaResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *GetLogUrlInMsaResponseBodyResultContent) SetMessage(v string) *GetLogUrlInMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *GetLogUrlInMsaResponseBodyResultContent) SetSuccess(v bool) *GetLogUrlInMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *GetLogUrlInMsaResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
