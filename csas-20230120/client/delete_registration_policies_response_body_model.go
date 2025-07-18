// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistrationPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRegistrationPoliciesResponseBody
	GetRequestId() *string
}

type DeleteRegistrationPoliciesResponseBody struct {
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRegistrationPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistrationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistrationPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRegistrationPoliciesResponseBody) SetRequestId(v string) *DeleteRegistrationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRegistrationPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
