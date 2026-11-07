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
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The gateway operation details.
	//
	// - **TraceDetails**: The XTrace sampling description. It includes TraceEnabled, which indicates whether XTrace is enabled, and Sample, which indicates the XTrace sampling rate.
	//
	// - **LogConfigDetails**: The log delivery service description. It includes LogEnabled, which indicates whether log delivery to Simple Log Service (SLS) is enabled, ProjectName, which indicates the destination project for log delivery, and LogStoreName, which indicates the Logstore name.
	//
	// - **EnableHardwareAcceleration**: Indicates whether hardware acceleration is enabled.
	//
	// - **DisableHttp2Alpn**: Indicates whether HTTP/2 is disabled.
	//
	// - **EnableWaf**: Indicates whether Web Application Firewall (WAF) is enabled.
	//
	// example:
	//
	// {\\"LogConfigDetails\\": {\\"LogEnabled\\": False}, \\"TraceDetails\\": {\\"TraceEnabled\\": False}}
	Data *GatewayOption `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// - If the request fails, the cause of the failure is returned, such as TaskId not found.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C92F4A4D-A2FD-593E-839E-F3D4DFD2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
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
