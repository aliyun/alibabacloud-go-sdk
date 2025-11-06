// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyContactFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyContactFieldResponseBody
	GetRequestId() *string
}

type VerifyContactFieldResponseBody struct {
	// example:
	//
	// ABAC3BAC-FCFA-4DAE-B47C-FA4105CB07C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyContactFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyContactFieldResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyContactFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyContactFieldResponseBody) SetRequestId(v string) *VerifyContactFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyContactFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
