// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttestorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAttestorResponseBody
	GetRequestId() *string
}

type DeleteAttestorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD380235-A0B8-540D-A0D5-D6288446****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAttestorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttestorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAttestorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAttestorResponseBody) SetRequestId(v string) *DeleteAttestorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAttestorResponseBody) Validate() error {
	return dara.Validate(s)
}
