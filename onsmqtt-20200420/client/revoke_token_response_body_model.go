// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeTokenResponseBody
	GetRequestId() *string
}

type RevokeTokenResponseBody struct {
	// The request ID. This parameter is a common parameter.
	//
	// example:
	//
	// 833EDFCB-C447-4CE3-B21F-3A4C2D1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeTokenResponseBody) SetRequestId(v string) *RevokeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
