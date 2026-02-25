// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateGatewayResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateGatewayResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateGatewayResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateGatewayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayResponseBody
	GetSuccess() *bool
}

type UpdateGatewayResponseBody struct {
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGatewayResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateGatewayResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayResponseBody) SetData(v bool) *UpdateGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetErrCode(v string) *UpdateGatewayResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetErrMessage(v string) *UpdateGatewayResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetRequestId(v string) *UpdateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetSuccess(v bool) *UpdateGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
