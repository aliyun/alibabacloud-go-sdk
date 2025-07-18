// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOtpConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOtpConfigResponseBody
	GetRequestId() *string
}

type DeleteOtpConfigResponseBody struct {
	// example:
	//
	// 54A4055A-343D-583E-9EAC-D12231148A68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOtpConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOtpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOtpConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOtpConfigResponseBody) SetRequestId(v string) *DeleteOtpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOtpConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
