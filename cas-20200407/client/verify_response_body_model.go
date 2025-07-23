// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyResponseBody
	GetRequestId() *string
	SetSignatureValid(v bool) *VerifyResponseBody
	GetSignatureValid() *bool
}

type VerifyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1ed33293-2e48-6b14-861e-538e28e408eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the signature is valid. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SignatureValid *bool `json:"SignatureValid,omitempty" xml:"SignatureValid,omitempty"`
}

func (s VerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyResponseBody) GetSignatureValid() *bool {
	return s.SignatureValid
}

func (s *VerifyResponseBody) SetRequestId(v string) *VerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyResponseBody) SetSignatureValid(v bool) *VerifyResponseBody {
	s.SignatureValid = &v
	return s
}

func (s *VerifyResponseBody) Validate() error {
	return dara.Validate(s)
}
