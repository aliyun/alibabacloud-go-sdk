// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyCredentialResponseBody
	GetRequestId() *string
}

type VerifyCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D5F0BDFB-A229-5F1D-B790-33709D43****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyCredentialResponseBody) SetRequestId(v string) *VerifyCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
