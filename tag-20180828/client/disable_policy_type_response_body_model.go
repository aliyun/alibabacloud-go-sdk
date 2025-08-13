// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolicyTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisablePolicyTypeResponseBody
	GetRequestId() *string
}

type DisablePolicyTypeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6E27F22C-EDA3-132E-A53F-77DE3BC2343D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisablePolicyTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisablePolicyTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DisablePolicyTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisablePolicyTypeResponseBody) SetRequestId(v string) *DisablePolicyTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisablePolicyTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
