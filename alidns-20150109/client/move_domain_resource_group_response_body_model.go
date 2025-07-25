// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveDomainResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveDomainResourceGroupResponseBody
	GetRequestId() *string
}

type MoveDomainResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveDomainResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveDomainResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveDomainResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveDomainResourceGroupResponseBody) SetRequestId(v string) *MoveDomainResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveDomainResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
