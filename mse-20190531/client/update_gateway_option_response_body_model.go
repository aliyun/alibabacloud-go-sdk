// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayOptionResponseBody
	GetCode() *int32
	SetData(v *GatewayOption) *UpdateGatewayOptionResponseBody
	GetData() *GatewayOption
	SetHttpStatusCode(v int32) *UpdateGatewayOptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayOptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayOptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayOptionResponseBody
	GetSuccess() *bool
}

type UpdateGatewayOptionResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {\\"LogConfigDetails\\": {\\"LogEnabled\\": True}, \\"TraceDetails\\": {\\"Sample\\": 17, \\"TraceEnabled\\": True}}
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
	// 	- If the request fails, an error message is returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 124D02EB-DBDD-534D-A701-B4A95D3D****
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

func (s UpdateGatewayOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayOptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayOptionResponseBody) GetData() *GatewayOption {
	return s.Data
}

func (s *UpdateGatewayOptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayOptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayOptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayOptionResponseBody) SetCode(v int32) *UpdateGatewayOptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetData(v *GatewayOption) *UpdateGatewayOptionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayOptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetMessage(v string) *UpdateGatewayOptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetRequestId(v string) *UpdateGatewayOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetSuccess(v bool) *UpdateGatewayOptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) Validate() error {
	return dara.Validate(s)
}
