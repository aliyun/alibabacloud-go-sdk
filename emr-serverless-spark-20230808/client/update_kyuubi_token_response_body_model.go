// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKyuubiTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKyuubiTokenResponseBody
	GetRequestId() *string
}

type UpdateKyuubiTokenResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateKyuubiTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKyuubiTokenResponseBody) SetRequestId(v string) *UpdateKyuubiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKyuubiTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
