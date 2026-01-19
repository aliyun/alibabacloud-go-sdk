// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPrimaryClientPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPrimaryClientPublicKeyResponseBody
	GetRequestId() *string
}

type SetPrimaryClientPublicKeyResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPrimaryClientPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPrimaryClientPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *SetPrimaryClientPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPrimaryClientPublicKeyResponseBody) SetRequestId(v string) *SetPrimaryClientPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPrimaryClientPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
