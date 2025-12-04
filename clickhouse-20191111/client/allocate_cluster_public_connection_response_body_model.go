// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateClusterPublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateClusterPublicConnectionResponseBody
	GetRequestId() *string
}

type AllocateClusterPublicConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateClusterPublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateClusterPublicConnectionResponseBody) SetRequestId(v string) *AllocateClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateClusterPublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
