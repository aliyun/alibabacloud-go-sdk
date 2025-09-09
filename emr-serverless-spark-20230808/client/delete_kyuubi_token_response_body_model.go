// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKyuubiTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKyuubiTokenResponseBody
	GetRequestId() *string
}

type DeleteKyuubiTokenResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteKyuubiTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKyuubiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKyuubiTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKyuubiTokenResponseBody) SetRequestId(v string) *DeleteKyuubiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKyuubiTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
