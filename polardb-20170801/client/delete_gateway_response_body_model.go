// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGatewayResponseBody
	GetRequestId() *string
}

type DeleteGatewayResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
