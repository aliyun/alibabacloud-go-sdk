// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBitcodeToMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadBitcodeToMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UploadBitcodeToMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *UploadBitcodeToMsaResponseBodyResultContent) *UploadBitcodeToMsaResponseBody
	GetResultContent() *UploadBitcodeToMsaResponseBodyResultContent
	SetResultMessage(v string) *UploadBitcodeToMsaResponseBody
	GetResultMessage() *string
}

type UploadBitcodeToMsaResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *UploadBitcodeToMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// SYSTEM_ERROR
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UploadBitcodeToMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadBitcodeToMsaResponseBody) GoString() string {
	return s.String()
}

func (s *UploadBitcodeToMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadBitcodeToMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UploadBitcodeToMsaResponseBody) GetResultContent() *UploadBitcodeToMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *UploadBitcodeToMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UploadBitcodeToMsaResponseBody) SetRequestId(v string) *UploadBitcodeToMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBody) SetResultCode(v string) *UploadBitcodeToMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBody) SetResultContent(v *UploadBitcodeToMsaResponseBodyResultContent) *UploadBitcodeToMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *UploadBitcodeToMsaResponseBody) SetResultMessage(v string) *UploadBitcodeToMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadBitcodeToMsaResponseBodyResultContent struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1234
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Normal
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadBitcodeToMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s UploadBitcodeToMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) SetCode(v string) *UploadBitcodeToMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) SetData(v string) *UploadBitcodeToMsaResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) SetMessage(v string) *UploadBitcodeToMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) SetSuccess(v bool) *UploadBitcodeToMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *UploadBitcodeToMsaResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
