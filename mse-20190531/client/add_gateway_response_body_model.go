// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewayResponseBody
	GetCode() *int32
	SetData(v *AddGatewayResponseBodyData) *AddGatewayResponseBody
	GetData() *AddGatewayResponseBodyData
	SetHttpStatusCode(v int32) *AddGatewayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayResponseBody
	GetSuccess() *bool
}

type AddGatewayResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *AddGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, a success message is returned.
	//
	// - If the request fails, the cause of the failure is returned.
	//
	// example:
	//
	// Succ
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayResponseBody) GetData() *AddGatewayResponseBodyData {
	return s.Data
}

func (s *AddGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayResponseBody) SetCode(v int32) *AddGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayResponseBody) SetData(v *AddGatewayResponseBodyData) *AddGatewayResponseBody {
	s.Data = v
	return s
}

func (s *AddGatewayResponseBody) SetHttpStatusCode(v int32) *AddGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayResponseBody) SetMessage(v string) *AddGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayResponseBody) SetRequestId(v string) *AddGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayResponseBody) SetSuccess(v bool) *AddGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddGatewayResponseBodyData struct {
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14cebbrvec4a5****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s AddGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddGatewayResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayResponseBodyData) SetGatewayUniqueId(v string) *AddGatewayResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
