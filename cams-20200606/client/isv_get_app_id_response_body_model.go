// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvGetAppIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *IsvGetAppIdResponseBody
	GetAccessDeniedDetail() *string
	SetAppId(v string) *IsvGetAppIdResponseBody
	GetAppId() *string
	SetCode(v string) *IsvGetAppIdResponseBody
	GetCode() *string
	SetConfigId(v string) *IsvGetAppIdResponseBody
	GetConfigId() *string
	SetMessage(v string) *IsvGetAppIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *IsvGetAppIdResponseBody
	GetRequestId() *string
}

type IsvGetAppIdResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 23hr3v
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the configuration item.
	//
	// example:
	//
	// 28972951817****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IsvGetAppIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsvGetAppIdResponseBody) GoString() string {
	return s.String()
}

func (s *IsvGetAppIdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *IsvGetAppIdResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *IsvGetAppIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *IsvGetAppIdResponseBody) GetConfigId() *string {
	return s.ConfigId
}

func (s *IsvGetAppIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IsvGetAppIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsvGetAppIdResponseBody) SetAccessDeniedDetail(v string) *IsvGetAppIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetAppId(v string) *IsvGetAppIdResponseBody {
	s.AppId = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetCode(v string) *IsvGetAppIdResponseBody {
	s.Code = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetConfigId(v string) *IsvGetAppIdResponseBody {
	s.ConfigId = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetMessage(v string) *IsvGetAppIdResponseBody {
	s.Message = &v
	return s
}

func (s *IsvGetAppIdResponseBody) SetRequestId(v string) *IsvGetAppIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsvGetAppIdResponseBody) Validate() error {
	return dara.Validate(s)
}
