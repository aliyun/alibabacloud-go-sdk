// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthConsumerResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayAuthConsumerResourceResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteGatewayAuthConsumerResourceResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteGatewayAuthConsumerResourceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteGatewayAuthConsumerResourceResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *DeleteGatewayAuthConsumerResourceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DeleteGatewayAuthConsumerResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayAuthConsumerResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayAuthConsumerResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayAuthConsumerResourceResponseBody
	GetSuccess() *bool
}

type DeleteGatewayAuthConsumerResourceResponseBody struct {
	// The code is used only to synchronize with the success status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation succeeded.
	//
	// - false: The operation failed.
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
	// The dynamic error message used to replace **%s*	- in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3369AD10-F1A6-4E6F-B99E-20F51826****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request succeeded.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayAuthConsumerResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthConsumerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetCode(v int32) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetData(v bool) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetDynamicCode(v string) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetDynamicMessage(v string) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetErrorCode(v string) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetMessage(v string) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetRequestId(v string) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) SetSuccess(v bool) *DeleteGatewayAuthConsumerResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
