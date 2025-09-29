// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialGetResultIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialGetResultIntlResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialGetResultIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialGetResultIntlResponseBody
	GetRequestId() *string
	SetResult(v *CredentialGetResultIntlResponseBodyResult) *CredentialGetResultIntlResponseBody
	GetResult() *CredentialGetResultIntlResponseBodyResult
}

type CredentialGetResultIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CredentialGetResultIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CredentialGetResultIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialGetResultIntlResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialGetResultIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialGetResultIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialGetResultIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialGetResultIntlResponseBody) GetResult() *CredentialGetResultIntlResponseBodyResult {
	return s.Result
}

func (s *CredentialGetResultIntlResponseBody) SetCode(v string) *CredentialGetResultIntlResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialGetResultIntlResponseBody) SetMessage(v string) *CredentialGetResultIntlResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialGetResultIntlResponseBody) SetRequestId(v string) *CredentialGetResultIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialGetResultIntlResponseBody) SetResult(v *CredentialGetResultIntlResponseBodyResult) *CredentialGetResultIntlResponseBody {
	s.Result = v
	return s
}

func (s *CredentialGetResultIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type CredentialGetResultIntlResponseBodyResult struct {
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
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s CredentialGetResultIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialGetResultIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CredentialGetResultIntlResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *CredentialGetResultIntlResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CredentialGetResultIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *CredentialGetResultIntlResponseBodyResult) SetExtIdInfo(v string) *CredentialGetResultIntlResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CredentialGetResultIntlResponseBodyResult) SetStatus(v string) *CredentialGetResultIntlResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CredentialGetResultIntlResponseBodyResult) SetSubCode(v string) *CredentialGetResultIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CredentialGetResultIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
