// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportGatewayApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ImportGatewayApisResponseBody
	GetAccessDeniedDetail() *string
	SetApiCount(v int32) *ImportGatewayApisResponseBody
	GetApiCount() *int32
	SetDownloadUrl(v string) *ImportGatewayApisResponseBody
	GetDownloadUrl() *string
	SetErrorField(v string) *ImportGatewayApisResponseBody
	GetErrorField() *string
	SetGroupCount(v int32) *ImportGatewayApisResponseBody
	GetGroupCount() *int32
	SetHttpStatusCode(v int32) *ImportGatewayApisResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ImportGatewayApisResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ImportGatewayApisResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ImportGatewayApisResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *ImportGatewayApisResponseBody
	GetSuccess() *bool
}

type ImportGatewayApisResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ApiCount           *int32  `json:"ApiCount,omitempty" xml:"ApiCount,omitempty"`
	DownloadUrl        *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ErrorField         *string `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
	GroupCount         *int32  `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportGatewayApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportGatewayApisResponseBody) GoString() string {
	return s.String()
}

func (s *ImportGatewayApisResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ImportGatewayApisResponseBody) GetApiCount() *int32 {
	return s.ApiCount
}

func (s *ImportGatewayApisResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ImportGatewayApisResponseBody) GetErrorField() *string {
	return s.ErrorField
}

func (s *ImportGatewayApisResponseBody) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *ImportGatewayApisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportGatewayApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportGatewayApisResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ImportGatewayApisResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ImportGatewayApisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportGatewayApisResponseBody) SetAccessDeniedDetail(v string) *ImportGatewayApisResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetApiCount(v int32) *ImportGatewayApisResponseBody {
	s.ApiCount = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetDownloadUrl(v string) *ImportGatewayApisResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetErrorField(v string) *ImportGatewayApisResponseBody {
	s.ErrorField = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetGroupCount(v int32) *ImportGatewayApisResponseBody {
	s.GroupCount = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetHttpStatusCode(v int32) *ImportGatewayApisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetRequestId(v string) *ImportGatewayApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetResultCode(v string) *ImportGatewayApisResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetResultMessage(v string) *ImportGatewayApisResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ImportGatewayApisResponseBody) SetSuccess(v bool) *ImportGatewayApisResponseBody {
	s.Success = &v
	return s
}

func (s *ImportGatewayApisResponseBody) Validate() error {
	return dara.Validate(s)
}
