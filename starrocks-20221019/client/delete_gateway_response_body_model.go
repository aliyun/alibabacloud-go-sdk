// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteGatewayResponseBody
	GetData() *bool
	SetErrCode(v string) *DeleteGatewayResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteGatewayResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteGatewayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayResponseBody
	GetSuccess() *bool
}

type DeleteGatewayResponseBody struct {
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
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGatewayResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteGatewayResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayResponseBody) SetData(v bool) *DeleteGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetErrCode(v string) *DeleteGatewayResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetErrMessage(v string) *DeleteGatewayResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetSuccess(v bool) *DeleteGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
