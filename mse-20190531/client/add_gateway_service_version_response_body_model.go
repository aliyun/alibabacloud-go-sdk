// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewayServiceVersionResponseBody
	GetCode() *int32
	SetData(v int64) *AddGatewayServiceVersionResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddGatewayServiceVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewayServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayServiceVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayServiceVersionResponseBody
	GetSuccess() *bool
}

type AddGatewayServiceVersionResponseBody struct {
	// The status code returned. The value 200 indicates that the request is successfully processed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the gateway service.
	//
	// example:
	//
	// 330
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 403
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
	// BA4046B6-CFC6-583C-B608-DD75011A590F
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

func (s AddGatewayServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayServiceVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayServiceVersionResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddGatewayServiceVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayServiceVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayServiceVersionResponseBody) SetCode(v int32) *AddGatewayServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayServiceVersionResponseBody) SetData(v int64) *AddGatewayServiceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewayServiceVersionResponseBody) SetHttpStatusCode(v int32) *AddGatewayServiceVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayServiceVersionResponseBody) SetMessage(v string) *AddGatewayServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayServiceVersionResponseBody) SetRequestId(v string) *AddGatewayServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayServiceVersionResponseBody) SetSuccess(v bool) *AddGatewayServiceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
