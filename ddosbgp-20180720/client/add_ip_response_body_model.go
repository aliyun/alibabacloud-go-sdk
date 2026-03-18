// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddIpResponseBody
	GetRequestId() *string
}

type AddIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIpResponseBody) SetRequestId(v string) *AddIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIpResponseBody) Validate() error {
	return dara.Validate(s)
}
