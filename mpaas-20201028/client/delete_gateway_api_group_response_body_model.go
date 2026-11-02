// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteGatewayApiGroupResponseBody
	GetAccessDeniedDetail() *string
	SetErrorField(v string) *DeleteGatewayApiGroupResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *DeleteGatewayApiGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteGatewayApiGroupResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteGatewayApiGroupResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteGatewayApiGroupResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *DeleteGatewayApiGroupResponseBody
	GetSuccess() *bool
}

type DeleteGatewayApiGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ErrorField         *string `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayApiGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteGatewayApiGroupResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *DeleteGatewayApiGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayApiGroupResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteGatewayApiGroupResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteGatewayApiGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayApiGroupResponseBody) SetAccessDeniedDetail(v string) *DeleteGatewayApiGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) SetErrorField(v string) *DeleteGatewayApiGroupResponseBody {
	s.ErrorField = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayApiGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) SetRequestId(v string) *DeleteGatewayApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) SetResultCode(v string) *DeleteGatewayApiGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) SetResultMessage(v string) *DeleteGatewayApiGroupResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) SetSuccess(v bool) *DeleteGatewayApiGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayApiGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
