// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAccessKeyResponseBody
	GetRequestId() *string
}

type UpdateAccessKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccessKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccessKeyResponseBody) SetRequestId(v string) *UpdateAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccessKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
