// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseInstanceResponseBody
	GetRequestId() *string
}

type ReleaseInstanceResponseBody struct {
	// example:
	//
	// 06FFAF5F-CD3E-4886-A849-AAB40DFF6515
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
