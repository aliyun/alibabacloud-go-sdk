// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewayAuthConsumerResponseBody
	GetCode() *int32
	SetData(v int64) *AddGatewayAuthConsumerResponseBody
	GetData() *int64
	SetDynamicCode(v string) *AddGatewayAuthConsumerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *AddGatewayAuthConsumerResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *AddGatewayAuthConsumerResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *AddGatewayAuthConsumerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewayAuthConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayAuthConsumerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayAuthConsumerResponseBody
	GetSuccess() *bool
}

type AddGatewayAuthConsumerResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the consumer.
	//
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in the `ErrMessage` parameter.
	//
	// >  If the return value of the `ErrMessage` parameter is `The Value of Input Parameter %s is not valid` and the return value of the `DynamicMessage` parameter is `DtsJobId`, the specified `DtsJobId` parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddGatewayAuthConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthConsumerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayAuthConsumerResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddGatewayAuthConsumerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *AddGatewayAuthConsumerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *AddGatewayAuthConsumerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddGatewayAuthConsumerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayAuthConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayAuthConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayAuthConsumerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayAuthConsumerResponseBody) SetCode(v int32) *AddGatewayAuthConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetData(v int64) *AddGatewayAuthConsumerResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetDynamicCode(v string) *AddGatewayAuthConsumerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetDynamicMessage(v string) *AddGatewayAuthConsumerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetErrorCode(v string) *AddGatewayAuthConsumerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetHttpStatusCode(v int32) *AddGatewayAuthConsumerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetMessage(v string) *AddGatewayAuthConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetRequestId(v string) *AddGatewayAuthConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) SetSuccess(v bool) *AddGatewayAuthConsumerResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayAuthConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
