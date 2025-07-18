// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEnterpriseAcceleratePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableEnterpriseAcceleratePolicyResponseBody
	GetRequestId() *string
}

type DisableEnterpriseAcceleratePolicyResponseBody struct {
	// example:
	//
	// E4C3E4CA-87CC-5EF6-91DD-D400A812EB43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableEnterpriseAcceleratePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableEnterpriseAcceleratePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableEnterpriseAcceleratePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableEnterpriseAcceleratePolicyResponseBody) SetRequestId(v string) *DisableEnterpriseAcceleratePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableEnterpriseAcceleratePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
