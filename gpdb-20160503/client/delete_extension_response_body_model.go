// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExtensionResponseBody
	GetRequestId() *string
}

type DeleteExtensionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExtensionResponseBody) SetRequestId(v string) *DeleteExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
