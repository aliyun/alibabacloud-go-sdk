// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppResponseBody
	GetAccessDeniedDetail() *string
	SetAppId(v string) *CreateAppResponseBody
	GetAppId() *string
	SetAppName(v string) *CreateAppResponseBody
	GetAppName() *string
	SetCode(v string) *CreateAppResponseBody
	GetCode() *string
	SetDynamicCode(v string) *CreateAppResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAppResponseBody
	GetSuccess() *bool
	SetUserCode(v string) *CreateAppResponseBody
	GetUserCode() *string
}

type CreateAppResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// app-bd5e3533
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// TestApp_g5t
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/j6if7e3w217z31q/j6if7e3w217z31q.sql.zip?Expires=1753331032&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=vOXj4E1%2FCqncWcDtu3UxxuOcyh0%3D
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E604ABBF-FD0F-5080-BE2B-BCF674A9E941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// OK
	UserCode *string `json:"UserCode,omitempty" xml:"UserCode,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAppResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAppResponseBody) GetUserCode() *string {
	return s.UserCode
}

func (s *CreateAppResponseBody) SetAccessDeniedDetail(v string) *CreateAppResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppResponseBody) SetAppId(v string) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetAppName(v string) *CreateAppResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppResponseBody) SetCode(v string) *CreateAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResponseBody) SetDynamicCode(v string) *CreateAppResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppResponseBody) SetDynamicMessage(v string) *CreateAppResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppResponseBody) SetHttpStatusCode(v int32) *CreateAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAppResponseBody) SetMessage(v string) *CreateAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetSuccess(v bool) *CreateAppResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAppResponseBody) SetUserCode(v string) *CreateAppResponseBody {
	s.UserCode = &v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	return dara.Validate(s)
}
