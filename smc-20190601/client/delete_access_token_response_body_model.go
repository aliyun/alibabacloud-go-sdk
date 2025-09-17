// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessTokenResponseBody
	GetRequestId() *string
}

type DeleteAccessTokenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DB4A7EA2-6FDA-5655-B067-854532FB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessTokenResponseBody) SetRequestId(v string) *DeleteAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
