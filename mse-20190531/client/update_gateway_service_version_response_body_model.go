// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayServiceVersionResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayServiceVersionResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayServiceVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayServiceVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayServiceVersionResponseBody
	GetSuccess() *bool
}

type UpdateGatewayServiceVersionResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// 614
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
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6580AA6-E285-58D2-B00B-12C051B3B7BF
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

func (s UpdateGatewayServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayServiceVersionResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayServiceVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayServiceVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayServiceVersionResponseBody) SetCode(v int32) *UpdateGatewayServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetData(v int64) *UpdateGatewayServiceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayServiceVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetMessage(v string) *UpdateGatewayServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetRequestId(v string) *UpdateGatewayServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetSuccess(v bool) *UpdateGatewayServiceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
