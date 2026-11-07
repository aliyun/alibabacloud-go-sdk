// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayAuthConsumerResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteGatewayAuthConsumerResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteGatewayAuthConsumerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteGatewayAuthConsumerResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *DeleteGatewayAuthConsumerResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DeleteGatewayAuthConsumerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayAuthConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayAuthConsumerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayAuthConsumerResponseBody
	GetSuccess() *bool
}

type DeleteGatewayAuthConsumerResponseBody struct {
	// The code is used only to synchronize with the success field.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the execution is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The placeholder in the dynamic error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the **%s*	- placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the value of the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
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

func (s DeleteGatewayAuthConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayAuthConsumerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetCode(v int32) *DeleteGatewayAuthConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetData(v bool) *DeleteGatewayAuthConsumerResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetDynamicCode(v string) *DeleteGatewayAuthConsumerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetDynamicMessage(v string) *DeleteGatewayAuthConsumerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetErrorCode(v string) *DeleteGatewayAuthConsumerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayAuthConsumerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetMessage(v string) *DeleteGatewayAuthConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetRequestId(v string) *DeleteGatewayAuthConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) SetSuccess(v bool) *DeleteGatewayAuthConsumerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
