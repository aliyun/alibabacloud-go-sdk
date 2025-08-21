// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRenderingInstanceGatewayResponseBody
	GetRequestId() *string
}

type DeleteRenderingInstanceGatewayResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRenderingInstanceGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRenderingInstanceGatewayResponseBody) SetRequestId(v string) *DeleteRenderingInstanceGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRenderingInstanceGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
