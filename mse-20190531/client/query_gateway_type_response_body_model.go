// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGatewayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryGatewayTypeResponseBody
	GetCode() *int32
	SetData(v []*string) *QueryGatewayTypeResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *QueryGatewayTypeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryGatewayTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryGatewayTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryGatewayTypeResponseBody
	GetSuccess() *bool
}

type QueryGatewayTypeResponseBody struct {
	// The status code returned. The value 200 indicates that the request was successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the gateway type. The data type of this parameter is List.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned. If the request is successful, a success message is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 52BA6DA6-A702-4362-A32F-DFF79655****
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

func (s QueryGatewayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGatewayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGatewayTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryGatewayTypeResponseBody) GetData() []*string {
	return s.Data
}

func (s *QueryGatewayTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryGatewayTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGatewayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGatewayTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGatewayTypeResponseBody) SetCode(v int32) *QueryGatewayTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetData(v []*string) *QueryGatewayTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetHttpStatusCode(v int32) *QueryGatewayTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetMessage(v string) *QueryGatewayTypeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetRequestId(v string) *QueryGatewayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetSuccess(v bool) *QueryGatewayTypeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
