// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGatewayResponseBody
	GetCode() *string
	SetData(v *CreateGatewayResponseBodyData) *CreateGatewayResponseBody
	GetData() *CreateGatewayResponseBodyData
	SetMessage(v string) *CreateGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGatewayResponseBody
	GetRequestId() *string
}

type CreateGatewayResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response payload.
	Data *CreateGatewayResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9CDE3E69-69C2-5402-83AD-ACA80B1AF35B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGatewayResponseBody) GetData() *CreateGatewayResponseBodyData {
	return s.Data
}

func (s *CreateGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayResponseBody) SetCode(v string) *CreateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayResponseBody) SetData(v *CreateGatewayResponseBodyData) *CreateGatewayResponseBody {
	s.Data = v
	return s
}

func (s *CreateGatewayResponseBody) SetMessage(v string) *CreateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGatewayResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s CreateGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGatewayResponseBodyData) SetGatewayId(v string) *CreateGatewayResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
