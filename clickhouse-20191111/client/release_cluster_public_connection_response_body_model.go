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
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
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
