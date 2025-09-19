// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBinarySecurityPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateBinarySecurityPolicyResponseBody
	GetRequestId() *string
}

type CreateBinarySecurityPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89AD16CC-97EE-50F3-9B12-9E28E5C8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBinarySecurityPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBinarySecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBinarySecurityPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBinarySecurityPolicyResponseBody) SetRequestId(v string) *CreateBinarySecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBinarySecurityPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
