// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpResponseBody
	GetRequestId() *string
}

type DeleteIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpResponseBody) SetRequestId(v string) *DeleteIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpResponseBody) Validate() error {
	return dara.Validate(s)
}
