// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelQualificationVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelQualificationVerificationResponseBody
	GetRequestId() *string
}

type CancelQualificationVerificationResponseBody struct {
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelQualificationVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelQualificationVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelQualificationVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelQualificationVerificationResponseBody) SetRequestId(v string) *CancelQualificationVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelQualificationVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
