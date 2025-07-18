// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAcceleratePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEnterpriseAcceleratePolicyResponseBody
	GetRequestId() *string
}

type CreateEnterpriseAcceleratePolicyResponseBody struct {
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnterpriseAcceleratePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAcceleratePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAcceleratePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnterpriseAcceleratePolicyResponseBody) SetRequestId(v string) *CreateEnterpriseAcceleratePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
