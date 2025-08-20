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
	// 868EF07F-D0B2-5043-B092-0C14CD00B65A
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
