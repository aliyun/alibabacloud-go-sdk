// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRenderingInstanceGatewayResponseBody
	GetRequestId() *string
}

type CreateRenderingInstanceGatewayResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRenderingInstanceGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRenderingInstanceGatewayResponseBody) SetRequestId(v string) *CreateRenderingInstanceGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRenderingInstanceGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
