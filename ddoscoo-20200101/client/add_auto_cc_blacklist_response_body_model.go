// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAutoCcBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAutoCcBlacklistResponseBody
	GetRequestId() *string
}

type AddAutoCcBlacklistResponseBody struct {
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAutoCcBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddAutoCcBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAutoCcBlacklistResponseBody) SetRequestId(v string) *AddAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAutoCcBlacklistResponseBody) Validate() error {
	return dara.Validate(s)
}
