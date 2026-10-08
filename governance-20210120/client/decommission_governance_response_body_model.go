// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecommissionGovernanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DecommissionGovernanceResponseBody
	GetRequestId() *string
}

type DecommissionGovernanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 37C4280D-C0AC-5EDD-B1EF-013808C4A357
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecommissionGovernanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecommissionGovernanceResponseBody) GoString() string {
	return s.String()
}

func (s *DecommissionGovernanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecommissionGovernanceResponseBody) SetRequestId(v string) *DecommissionGovernanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecommissionGovernanceResponseBody) Validate() error {
	return dara.Validate(s)
}
