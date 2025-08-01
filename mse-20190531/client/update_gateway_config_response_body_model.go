// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGatewayConfigResponseBody
	GetRequestId() *string
}

type UpdateGatewayConfigResponseBody struct {
	// example:
	//
	// AF21683A-29C7-4853-AC0F-B5ADEE4D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayConfigResponseBody) SetRequestId(v string) *UpdateGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
