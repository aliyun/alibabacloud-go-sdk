// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayServiceCheckResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayServiceCheckResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayServiceCheckResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayServiceCheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayServiceCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayServiceCheckResponseBody
	GetSuccess() *bool
}

type UpdateGatewayServiceCheckResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The service ID of the operation.
	//
	// example:
	//
	// 12
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 62CBFCB8-DDC6-588C-BF1B-88828AF*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayServiceCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceCheckResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceCheckResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayServiceCheckResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayServiceCheckResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayServiceCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayServiceCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayServiceCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayServiceCheckResponseBody) SetCode(v int32) *UpdateGatewayServiceCheckResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponseBody) SetData(v int64) *UpdateGatewayServiceCheckResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayServiceCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponseBody) SetMessage(v string) *UpdateGatewayServiceCheckResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponseBody) SetRequestId(v string) *UpdateGatewayServiceCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponseBody) SetSuccess(v bool) *UpdateGatewayServiceCheckResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
