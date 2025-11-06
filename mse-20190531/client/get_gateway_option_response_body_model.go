// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayOptionResponseBody
	GetCode() *int32
	SetData(v *GatewayOption) *GetGatewayOptionResponseBody
	GetData() *GatewayOption
	SetHttpStatusCode(v int32) *GetGatewayOptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayOptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayOptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayOptionResponseBody
	GetSuccess() *bool
}

type GetGatewayOptionResponseBody struct {
	// The status code returned. The value 200 indicates that the request is successfully processed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed configurations of the gateway.
	//
	// 	- **TraceDetails**: the sampling description of Managed Service for OpenTelemetry. Content: TraceEnabled indicates whether Managed Service for OpenTelemetry is activated. Sample indicates the sampling rate of Managed Service for OpenTelemetry.
	//
	// 	- **LogConfigDetails**: the description of Simple Log Service. Content: LogEnabled indicates whether Simple Log Service is activated. ProjectName indicates the Simple Log Service project to which logs are delivered. LogStoreName indicates the name of the Logstore.
	//
	// 	- **EnableHardwareAcceleration**: indicates whether hardware acceleration is enabled.
	//
	// 	- **DisableHttp2Alpn**: indicates whether the HTTP/2 protocol is disabled.
	//
	// 	- **EnableWaf**: indicates whether Web Application Firewall (WAF) is enabled.
	//
	// example:
	//
	// {\\"LogConfigDetails\\": {\\"LogEnabled\\": False}, \\"TraceDetails\\": {\\"TraceEnabled\\": False}}
	Data *GatewayOption `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned, such as the "TaskId not found" message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C92F4A4D-A2FD-593E-839E-F3D4DFD2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayOptionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayOptionResponseBody) GetData() *GatewayOption {
	return s.Data
}

func (s *GetGatewayOptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayOptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayOptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayOptionResponseBody) SetCode(v int32) *GetGatewayOptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetData(v *GatewayOption) *GetGatewayOptionResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayOptionResponseBody) SetHttpStatusCode(v int32) *GetGatewayOptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetMessage(v string) *GetGatewayOptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetRequestId(v string) *GetGatewayOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetSuccess(v bool) *GetGatewayOptionResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayOptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
