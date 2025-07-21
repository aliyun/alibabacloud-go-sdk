// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExchangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExchangeResponseBody
	GetRequestId() *string
}

type DeleteExchangeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6961FFB8-6358-4EDC-9E3C-4A0C56CE6***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExchangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExchangeResponseBody) SetRequestId(v string) *DeleteExchangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExchangeResponseBody) Validate() error {
	return dara.Validate(s)
}
