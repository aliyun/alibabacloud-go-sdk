// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AddGatewayResponseBody
	GetData() *bool
	SetErrCode(v string) *AddGatewayResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddGatewayResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AddGatewayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayResponseBody
	GetSuccess() *bool
}

type AddGatewayResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
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
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddGatewayResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddGatewayResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayResponseBody) SetData(v bool) *AddGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewayResponseBody) SetErrCode(v string) *AddGatewayResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddGatewayResponseBody) SetErrMessage(v string) *AddGatewayResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddGatewayResponseBody) SetHttpStatusCode(v int32) *AddGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayResponseBody) SetRequestId(v string) *AddGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayResponseBody) SetSuccess(v bool) *AddGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
