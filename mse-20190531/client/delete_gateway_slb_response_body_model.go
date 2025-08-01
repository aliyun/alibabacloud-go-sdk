// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewaySlbResponseBody
	GetCode() *int32
	SetData(v string) *DeleteGatewaySlbResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteGatewaySlbResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewaySlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewaySlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewaySlbResponseBody
	GetSuccess() *bool
}

type DeleteGatewaySlbResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deletion result.
	//
	// example:
	//
	// 28
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
	// 9297B27D-D932-5E9F-93B9-99D6C5F3A879
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

func (s DeleteGatewaySlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySlbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewaySlbResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteGatewaySlbResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewaySlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewaySlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewaySlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewaySlbResponseBody) SetCode(v int32) *DeleteGatewaySlbResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewaySlbResponseBody) SetData(v string) *DeleteGatewaySlbResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewaySlbResponseBody) SetHttpStatusCode(v int32) *DeleteGatewaySlbResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewaySlbResponseBody) SetMessage(v string) *DeleteGatewaySlbResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewaySlbResponseBody) SetRequestId(v string) *DeleteGatewaySlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewaySlbResponseBody) SetSuccess(v bool) *DeleteGatewaySlbResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewaySlbResponseBody) Validate() error {
	return dara.Validate(s)
}
