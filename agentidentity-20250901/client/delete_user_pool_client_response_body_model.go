// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPoolClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserPoolClientResponseBody
	GetRequestId() *string
}

type DeleteUserPoolClientResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPoolClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPoolClientResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPoolClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserPoolClientResponseBody) SetRequestId(v string) *DeleteUserPoolClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserPoolClientResponseBody) Validate() error {
	return dara.Validate(s)
}
