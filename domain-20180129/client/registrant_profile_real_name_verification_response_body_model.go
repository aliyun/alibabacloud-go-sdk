// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistrantProfileRealNameVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RegistrantProfileRealNameVerificationResponseBody
	GetRequestId() *string
}

type RegistrantProfileRealNameVerificationResponseBody struct {
	// example:
	//
	// 51D584A2-0CCD-4336-AD7D-1AD4C67B5545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegistrantProfileRealNameVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegistrantProfileRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *RegistrantProfileRealNameVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegistrantProfileRealNameVerificationResponseBody) SetRequestId(v string) *RegistrantProfileRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
