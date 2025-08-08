// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMgsApirestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveMgsApirestResponseBody
	GetRequestId() *string
	SetResultCode(v string) *SaveMgsApirestResponseBody
	GetResultCode() *string
	SetResultContent(v *SaveMgsApirestResponseBodyResultContent) *SaveMgsApirestResponseBody
	GetResultContent() *SaveMgsApirestResponseBodyResultContent
	SetResultMessage(v string) *SaveMgsApirestResponseBody
	GetResultMessage() *string
}

type SaveMgsApirestResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *SaveMgsApirestResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s SaveMgsApirestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveMgsApirestResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMgsApirestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveMgsApirestResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *SaveMgsApirestResponseBody) GetResultContent() *SaveMgsApirestResponseBodyResultContent {
	return s.ResultContent
}

func (s *SaveMgsApirestResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *SaveMgsApirestResponseBody) SetRequestId(v string) *SaveMgsApirestResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMgsApirestResponseBody) SetResultCode(v string) *SaveMgsApirestResponseBody {
	s.ResultCode = &v
	return s
}

func (s *SaveMgsApirestResponseBody) SetResultContent(v *SaveMgsApirestResponseBodyResultContent) *SaveMgsApirestResponseBody {
	s.ResultContent = v
	return s
}

func (s *SaveMgsApirestResponseBody) SetResultMessage(v string) *SaveMgsApirestResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *SaveMgsApirestResponseBody) Validate() error {
	return dara.Validate(s)
}

type SaveMgsApirestResponseBodyResultContent struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Value        *bool   `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveMgsApirestResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s SaveMgsApirestResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *SaveMgsApirestResponseBodyResultContent) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveMgsApirestResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *SaveMgsApirestResponseBodyResultContent) GetValue() *bool {
	return s.Value
}

func (s *SaveMgsApirestResponseBodyResultContent) SetErrorMessage(v string) *SaveMgsApirestResponseBodyResultContent {
	s.ErrorMessage = &v
	return s
}

func (s *SaveMgsApirestResponseBodyResultContent) SetSuccess(v bool) *SaveMgsApirestResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *SaveMgsApirestResponseBodyResultContent) SetValue(v bool) *SaveMgsApirestResponseBodyResultContent {
	s.Value = &v
	return s
}

func (s *SaveMgsApirestResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
