// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseClusterPublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseClusterPublicConnectionResponseBody
	GetRequestId() *string
}

type ReleaseClusterPublicConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A94B6C02-7BD4-5D67-9776-3AC8317E8DD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseClusterPublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseClusterPublicConnectionResponseBody) SetRequestId(v string) *ReleaseClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
