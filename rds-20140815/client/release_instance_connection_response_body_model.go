// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseInstanceConnectionResponseBody
	GetRequestId() *string
}

type ReleaseInstanceConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceConnectionResponseBody) SetRequestId(v string) *ReleaseInstanceConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
