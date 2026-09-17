// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialSubmitIntlV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialSubmitIntlV2ResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialSubmitIntlV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialSubmitIntlV2ResponseBody
	GetRequestId() *string
	SetResult(v *CredentialSubmitIntlV2ResponseBodyResult) *CredentialSubmitIntlV2ResponseBody
	GetResult() *CredentialSubmitIntlV2ResponseBodyResult
}

type CredentialSubmitIntlV2ResponseBody struct {
	// The return code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F971622-38C0-5F56-B2EC-315367979B4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *CredentialSubmitIntlV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CredentialSubmitIntlV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialSubmitIntlV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialSubmitIntlV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialSubmitIntlV2ResponseBody) GetResult() *CredentialSubmitIntlV2ResponseBodyResult {
	return s.Result
}

func (s *CredentialSubmitIntlV2ResponseBody) SetCode(v string) *CredentialSubmitIntlV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialSubmitIntlV2ResponseBody) SetMessage(v string) *CredentialSubmitIntlV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialSubmitIntlV2ResponseBody) SetRequestId(v string) *CredentialSubmitIntlV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialSubmitIntlV2ResponseBody) SetResult(v *CredentialSubmitIntlV2ResponseBodyResult) *CredentialSubmitIntlV2ResponseBody {
	s.Result = v
	return s
}

func (s *CredentialSubmitIntlV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialSubmitIntlV2ResponseBodyResult struct {
	// The unique identifier of the verification request.
	//
	// example:
	//
	// hk573be80f944d95ac812e0*******a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CredentialSubmitIntlV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlV2ResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *CredentialSubmitIntlV2ResponseBodyResult) SetTransactionId(v string) *CredentialSubmitIntlV2ResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *CredentialSubmitIntlV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
