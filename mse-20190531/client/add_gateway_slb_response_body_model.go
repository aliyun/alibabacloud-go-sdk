// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewaySlbResponseBody
	GetCode() *int32
	SetData(v string) *AddGatewaySlbResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *AddGatewaySlbResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewaySlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewaySlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewaySlbResponseBody
	GetSuccess() *bool
}

type AddGatewaySlbResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// lb-uf6dqr4ondqi5w3df7hdf
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, a success message is returned.
	//
	// - If the request fails, the cause of the failure is returned.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddGatewaySlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySlbResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewaySlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewaySlbResponseBody) GetData() *string {
	return s.Data
}

func (s *AddGatewaySlbResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewaySlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewaySlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewaySlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewaySlbResponseBody) SetCode(v int32) *AddGatewaySlbResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewaySlbResponseBody) SetData(v string) *AddGatewaySlbResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewaySlbResponseBody) SetHttpStatusCode(v int32) *AddGatewaySlbResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewaySlbResponseBody) SetMessage(v string) *AddGatewaySlbResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewaySlbResponseBody) SetRequestId(v string) *AddGatewaySlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewaySlbResponseBody) SetSuccess(v bool) *AddGatewaySlbResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewaySlbResponseBody) Validate() error {
	return dara.Validate(s)
}
