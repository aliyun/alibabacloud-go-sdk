// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnterpriseAcceleratePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEnterpriseAcceleratePolicyResponseBody
	GetRequestId() *string
}

type ModifyEnterpriseAcceleratePolicyResponseBody struct {
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEnterpriseAcceleratePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnterpriseAcceleratePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEnterpriseAcceleratePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEnterpriseAcceleratePolicyResponseBody) SetRequestId(v string) *ModifyEnterpriseAcceleratePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
