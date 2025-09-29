// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialSubmitIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialSubmitIntlResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialSubmitIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialSubmitIntlResponseBody
	GetRequestId() *string
	SetResult(v *CredentialSubmitIntlResponseBodyResult) *CredentialSubmitIntlResponseBody
	GetResult() *CredentialSubmitIntlResponseBodyResult
}

type CredentialSubmitIntlResponseBody struct {
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CredentialSubmitIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CredentialSubmitIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialSubmitIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialSubmitIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialSubmitIntlResponseBody) GetResult() *CredentialSubmitIntlResponseBodyResult {
	return s.Result
}

func (s *CredentialSubmitIntlResponseBody) SetCode(v string) *CredentialSubmitIntlResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialSubmitIntlResponseBody) SetMessage(v string) *CredentialSubmitIntlResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialSubmitIntlResponseBody) SetRequestId(v string) *CredentialSubmitIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialSubmitIntlResponseBody) SetResult(v *CredentialSubmitIntlResponseBodyResult) *CredentialSubmitIntlResponseBody {
	s.Result = v
	return s
}

func (s *CredentialSubmitIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type CredentialSubmitIntlResponseBodyResult struct {
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CredentialSubmitIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *CredentialSubmitIntlResponseBodyResult) SetTransactionId(v string) *CredentialSubmitIntlResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *CredentialSubmitIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
