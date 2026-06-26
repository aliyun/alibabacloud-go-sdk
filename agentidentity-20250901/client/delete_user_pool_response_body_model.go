// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserPoolResponseBody
	GetRequestId() *string
}

type DeleteUserPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserPoolResponseBody) SetRequestId(v string) *DeleteUserPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
