// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleasePublicConnectionResponseBody
	GetRequestId() *string
}

type ReleasePublicConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePublicConnectionResponseBody) SetRequestId(v string) *ReleasePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
