// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyAppResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyAppResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyAppResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyAppResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAppResponseBody
	GetSuccess() *bool
	SetUserCode(v string) *ModifyAppResponseBody
	GetUserCode() *string
}

type ModifyAppResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/gcqe6nn722rw1g5/gcqe6nn722rw1g5.diff.zip?Expires=1750651377&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=55f4Rcml1vjydPgiv5c9KRdSGQo%3D
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// OK
	UserCode *string `json:"UserCode,omitempty" xml:"UserCode,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAppResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyAppResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAppResponseBody) GetUserCode() *string {
	return s.UserCode
}

func (s *ModifyAppResponseBody) SetAccessDeniedDetail(v string) *ModifyAppResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyAppResponseBody) SetCode(v string) *ModifyAppResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppResponseBody) SetDynamicCode(v string) *ModifyAppResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyAppResponseBody) SetDynamicMessage(v string) *ModifyAppResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyAppResponseBody) SetHttpStatusCode(v int32) *ModifyAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAppResponseBody) SetMessage(v string) *ModifyAppResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppResponseBody) SetSuccess(v bool) *ModifyAppResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAppResponseBody) SetUserCode(v string) *ModifyAppResponseBody {
	s.UserCode = &v
	return s
}

func (s *ModifyAppResponseBody) Validate() error {
	return dara.Validate(s)
}
