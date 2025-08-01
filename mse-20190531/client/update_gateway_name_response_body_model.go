// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayNameResponseBody
	GetCode() *int32
	SetData(v string) *UpdateGatewayNameResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateGatewayNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayNameResponseBody
	GetSuccess() *bool
}

type UpdateGatewayNameResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 8C95711F-E702-5395-BFAA-21BA946CDE47
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

func (s UpdateGatewayNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayNameResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateGatewayNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayNameResponseBody) SetCode(v int32) *UpdateGatewayNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetData(v string) *UpdateGatewayNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetMessage(v string) *UpdateGatewayNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetRequestId(v string) *UpdateGatewayNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetSuccess(v bool) *UpdateGatewayNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) Validate() error {
	return dara.Validate(s)
}
