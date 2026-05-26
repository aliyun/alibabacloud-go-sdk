// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicySetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolicySetResponseBody
	GetRequestId() *string
}

type DeletePolicySetResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicySetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicySetResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicySetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicySetResponseBody) SetRequestId(v string) *DeletePolicySetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicySetResponseBody) Validate() error {
	return dara.Validate(s)
}
