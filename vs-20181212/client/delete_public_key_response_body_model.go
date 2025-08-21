// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePublicKeyResponseBody
	GetRequestId() *string
}

type DeletePublicKeyResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePublicKeyResponseBody) SetRequestId(v string) *DeletePublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
