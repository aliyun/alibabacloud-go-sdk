// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateGatewayApiResponseBody
	GetAccessDeniedDetail() *string
	SetErrorField(v string) *UpdateGatewayApiResponseBody
	GetErrorField() *string
	SetHttpStatusCode(v int32) *UpdateGatewayApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateGatewayApiResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateGatewayApiResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *UpdateGatewayApiResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *UpdateGatewayApiResponseBody
	GetSuccess() *bool
}

type UpdateGatewayApiResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ErrorField         *string `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateGatewayApiResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *UpdateGatewayApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayApiResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateGatewayApiResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateGatewayApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayApiResponseBody) SetAccessDeniedDetail(v string) *UpdateGatewayApiResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) SetErrorField(v string) *UpdateGatewayApiResponseBody {
	s.ErrorField = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) SetRequestId(v string) *UpdateGatewayApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) SetResultCode(v string) *UpdateGatewayApiResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) SetResultMessage(v string) *UpdateGatewayApiResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) SetSuccess(v bool) *UpdateGatewayApiResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayApiResponseBody) Validate() error {
	return dara.Validate(s)
}
