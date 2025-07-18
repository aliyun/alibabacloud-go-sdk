// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseAcceleratePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnterpriseAcceleratePolicyResponseBody
	GetRequestId() *string
}

type DeleteEnterpriseAcceleratePolicyResponseBody struct {
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnterpriseAcceleratePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseAcceleratePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseAcceleratePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnterpriseAcceleratePolicyResponseBody) SetRequestId(v string) *DeleteEnterpriseAcceleratePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseAcceleratePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
