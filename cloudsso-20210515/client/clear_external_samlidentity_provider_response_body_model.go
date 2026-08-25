// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearExternalSAMLIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearExternalSAMLIdentityProviderResponseBody
	GetRequestId() *string
}

type ClearExternalSAMLIdentityProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 96D1E5FF-0301-5636-8D33-071E033CFB82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearExternalSAMLIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearExternalSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *ClearExternalSAMLIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearExternalSAMLIdentityProviderResponseBody) SetRequestId(v string) *ClearExternalSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearExternalSAMLIdentityProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
