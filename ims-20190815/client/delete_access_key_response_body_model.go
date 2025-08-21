// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessKeyResponseBody
	GetRequestId() *string
}

type DeleteAccessKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessKeyResponseBody) SetRequestId(v string) *DeleteAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
