// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDomainVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelDomainVerificationResponseBody
	GetRequestId() *string
}

type CancelDomainVerificationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AC0AF67-D303-4EB9-B20E-B4D4B2C3F97B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDomainVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDomainVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDomainVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDomainVerificationResponseBody) SetRequestId(v string) *CancelDomainVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDomainVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
