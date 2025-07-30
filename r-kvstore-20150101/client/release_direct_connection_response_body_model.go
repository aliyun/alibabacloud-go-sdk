// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDirectConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseDirectConnectionResponseBody
	GetRequestId() *string
}

type ReleaseDirectConnectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseDirectConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDirectConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseDirectConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseDirectConnectionResponseBody) SetRequestId(v string) *ReleaseDirectConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseDirectConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
