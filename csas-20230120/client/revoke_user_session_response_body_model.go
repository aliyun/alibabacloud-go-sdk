// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeUserSessionResponseBody
	GetRequestId() *string
}

type RevokeUserSessionResponseBody struct {
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeUserSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserSessionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeUserSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeUserSessionResponseBody) SetRequestId(v string) *RevokeUserSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeUserSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
