// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGatewayRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryGatewayRegionResponseBody
	GetCode() *int32
	SetData(v []*string) *QueryGatewayRegionResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *QueryGatewayRegionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryGatewayRegionResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryGatewayRegionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryGatewayRegionResponseBody
	GetSuccess() *bool
}

type QueryGatewayRegionResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 9e78a671-4b9b-4dd4-99c1-0b9da87d3dec
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

func (s QueryGatewayRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGatewayRegionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGatewayRegionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryGatewayRegionResponseBody) GetData() []*string {
	return s.Data
}

func (s *QueryGatewayRegionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryGatewayRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGatewayRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGatewayRegionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGatewayRegionResponseBody) SetCode(v int32) *QueryGatewayRegionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetData(v []*string) *QueryGatewayRegionResponseBody {
	s.Data = v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetHttpStatusCode(v int32) *QueryGatewayRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetMessage(v string) *QueryGatewayRegionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetRequestId(v string) *QueryGatewayRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetSuccess(v bool) *QueryGatewayRegionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
