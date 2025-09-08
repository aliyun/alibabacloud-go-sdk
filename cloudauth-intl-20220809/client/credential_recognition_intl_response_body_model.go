// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialRecognitionIntlResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialRecognitionIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialRecognitionIntlResponseBody
	GetRequestId() *string
	SetResult(v *CredentialRecognitionIntlResponseBodyResult) *CredentialRecognitionIntlResponseBody
	GetResult() *CredentialRecognitionIntlResponseBodyResult
}

type CredentialRecognitionIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CredentialRecognitionIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CredentialRecognitionIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialRecognitionIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialRecognitionIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialRecognitionIntlResponseBody) GetResult() *CredentialRecognitionIntlResponseBodyResult {
	return s.Result
}

func (s *CredentialRecognitionIntlResponseBody) SetCode(v string) *CredentialRecognitionIntlResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialRecognitionIntlResponseBody) SetMessage(v string) *CredentialRecognitionIntlResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialRecognitionIntlResponseBody) SetRequestId(v string) *CredentialRecognitionIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialRecognitionIntlResponseBody) SetResult(v *CredentialRecognitionIntlResponseBodyResult) *CredentialRecognitionIntlResponseBody {
	s.Result = v
	return s
}

func (s *CredentialRecognitionIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type CredentialRecognitionIntlResponseBodyResult struct {
	// example:
	//
	// {
	//
	//   "address": "",
	//
	//   "name":""
	//
	// }
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// example:
	//
	// S
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CredentialRecognitionIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *CredentialRecognitionIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *CredentialRecognitionIntlResponseBodyResult) GetSuccess() *string {
	return s.Success
}

func (s *CredentialRecognitionIntlResponseBodyResult) SetExtIdInfo(v string) *CredentialRecognitionIntlResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CredentialRecognitionIntlResponseBodyResult) SetSubCode(v string) *CredentialRecognitionIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CredentialRecognitionIntlResponseBodyResult) SetSuccess(v string) *CredentialRecognitionIntlResponseBodyResult {
	s.Success = &v
	return s
}

func (s *CredentialRecognitionIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
