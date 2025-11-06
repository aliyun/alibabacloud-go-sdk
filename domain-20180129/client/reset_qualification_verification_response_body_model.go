// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetQualificationVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetQualificationVerificationResponseBody
	GetRequestId() *string
}

type ResetQualificationVerificationResponseBody struct {
	// example:
	//
	// D6CB3623-4726-4947-AC2B-2C6E673B447C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetQualificationVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetQualificationVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ResetQualificationVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetQualificationVerificationResponseBody) SetRequestId(v string) *ResetQualificationVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetQualificationVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
