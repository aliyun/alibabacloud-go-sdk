// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDfpExtProcServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayDfpExtProcServiceResponseBody
	GetCode() *string
	SetData(v *GetGatewayDfpExtProcServiceResponseBodyData) *GetGatewayDfpExtProcServiceResponseBody
	GetData() *GetGatewayDfpExtProcServiceResponseBodyData
	SetMessage(v string) *GetGatewayDfpExtProcServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayDfpExtProcServiceResponseBody
	GetRequestId() *string
}

type GetGatewayDfpExtProcServiceResponseBody struct {
	// The result code of the request. The value is Ok if the request is successful.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The binding information of the dynamic forwarding ExtProc service for the gateway.
	//
	// example:
	//
	// {"bound":true,"serviceId":"svc-cr8k0udlhtgt80sh7png","servicePort":9002}
	Data *GetGatewayDfpExtProcServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The result message of the request. If the request failed, the message contains the error message of the fault.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used for troubleshooting.
	//
	// example:
	//
	// F7B8A1C2-1234-5678-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayDfpExtProcServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDfpExtProcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayDfpExtProcServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayDfpExtProcServiceResponseBody) GetData() *GetGatewayDfpExtProcServiceResponseBodyData {
	return s.Data
}

func (s *GetGatewayDfpExtProcServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayDfpExtProcServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayDfpExtProcServiceResponseBody) SetCode(v string) *GetGatewayDfpExtProcServiceResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBody) SetData(v *GetGatewayDfpExtProcServiceResponseBodyData) *GetGatewayDfpExtProcServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBody) SetMessage(v string) *GetGatewayDfpExtProcServiceResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBody) SetRequestId(v string) *GetGatewayDfpExtProcServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayDfpExtProcServiceResponseBodyData struct {
	// Indicates whether the gateway is bound to a dynamic forwarding ExtProc service.
	//
	// example:
	//
	// true
	Bound *bool `json:"bound,omitempty" xml:"bound,omitempty"`
	// The ID of the bound ExtProc gRPC service. This parameter is not returned if no service is bound.
	//
	// example:
	//
	// svc-cr8k0udlhtgt80sh7png
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The gRPC port of the bound service. This parameter is not returned if no service is bound.
	//
	// example:
	//
	// 9002
	ServicePort *int32 `json:"servicePort,omitempty" xml:"servicePort,omitempty"`
}

func (s GetGatewayDfpExtProcServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDfpExtProcServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) GetBound() *bool {
	return s.Bound
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) SetBound(v bool) *GetGatewayDfpExtProcServiceResponseBodyData {
	s.Bound = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) SetServiceId(v string) *GetGatewayDfpExtProcServiceResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) SetServicePort(v int32) *GetGatewayDfpExtProcServiceResponseBodyData {
	s.ServicePort = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
