// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteGatewayApiResponseBody
	GetAccessDeniedDetail() *string
	SetErrorField(v string) *DeleteGatewayApiResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *DeleteGatewayApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteGatewayApiResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteGatewayApiResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteGatewayApiResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *DeleteGatewayApiResponseBody
	GetSuccess() *bool
}

type DeleteGatewayApiResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ErrorField         *string `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteGatewayApiResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *DeleteGatewayApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayApiResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteGatewayApiResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteGatewayApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayApiResponseBody) SetAccessDeniedDetail(v string) *DeleteGatewayApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) SetErrorField(v string) *DeleteGatewayApiResponseBody {
	s.ErrorField = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) SetRequestId(v string) *DeleteGatewayApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) SetResultCode(v string) *DeleteGatewayApiResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) SetResultMessage(v string) *DeleteGatewayApiResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) SetSuccess(v bool) *DeleteGatewayApiResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayApiResponseBody) Validate() error {
	return dara.Validate(s)
}
