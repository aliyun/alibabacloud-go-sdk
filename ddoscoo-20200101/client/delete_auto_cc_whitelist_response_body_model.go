// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoCcWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAutoCcWhitelistResponseBody
	GetRequestId() *string
}

type DeleteAutoCcWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoCcWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoCcWhitelistResponseBody) SetRequestId(v string) *DeleteAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoCcWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
