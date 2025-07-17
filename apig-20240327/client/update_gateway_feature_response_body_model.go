// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayFeatureResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGatewayFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayFeatureResponseBody
	GetRequestId() *string
}

type UpdateGatewayFeatureResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 043360DA-ED3B-5386-9B7A-D94DECF99A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFeatureResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayFeatureResponseBody) SetCode(v string) *UpdateGatewayFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayFeatureResponseBody) SetMessage(v string) *UpdateGatewayFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayFeatureResponseBody) SetRequestId(v string) *UpdateGatewayFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
