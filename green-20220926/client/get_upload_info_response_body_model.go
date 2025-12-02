// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *GetUploadInfoResponseBody
	GetAccessId() *string
	SetCode(v int32) *GetUploadInfoResponseBody
	GetCode() *int32
	SetExpire(v int64) *GetUploadInfoResponseBody
	GetExpire() *int64
	SetFolder(v string) *GetUploadInfoResponseBody
	GetFolder() *string
	SetHost(v string) *GetUploadInfoResponseBody
	GetHost() *string
	SetHttpStatusCode(v int32) *GetUploadInfoResponseBody
	GetHttpStatusCode() *int32
	SetKey(v string) *GetUploadInfoResponseBody
	GetKey() *string
	SetMsg(v string) *GetUploadInfoResponseBody
	GetMsg() *string
	SetName(v string) *GetUploadInfoResponseBody
	GetName() *string
	SetPolicy(v string) *GetUploadInfoResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GetUploadInfoResponseBody
	GetRequestId() *string
	SetSignature(v string) *GetUploadInfoResponseBody
	GetSignature() *string
	SetSuccess(v bool) *GetUploadInfoResponseBody
	GetSuccess() *bool
}

type GetUploadInfoResponseBody struct {
	// Upload authorization ID.
	//
	// example:
	//
	// LTAI5t9HM*****EXQmw3DVH
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// Error code, consistent with HTTP status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// In seconds.
	//
	// example:
	//
	// 900
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// Folder name.
	//
	// example:
	//
	// image/upload/xxx
	Folder *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	// Upload host.
	//
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Key used for uploading files.
	//
	// example:
	//
	// image/upload/xxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Further description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Used for front-end image upload.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS upload file Policy.
	//
	// example:
	//
	// xxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Upload signature information.
	//
	// example:
	//
	// iyu7VHblYj+mEF9p46cdGOlNPAw=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *GetUploadInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetUploadInfoResponseBody) GetExpire() *int64 {
	return s.Expire
}

func (s *GetUploadInfoResponseBody) GetFolder() *string {
	return s.Folder
}

func (s *GetUploadInfoResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetUploadInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUploadInfoResponseBody) GetKey() *string {
	return s.Key
}

func (s *GetUploadInfoResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetUploadInfoResponseBody) GetName() *string {
	return s.Name
}

func (s *GetUploadInfoResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadInfoResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *GetUploadInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUploadInfoResponseBody) SetAccessId(v string) *GetUploadInfoResponseBody {
	s.AccessId = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetCode(v int32) *GetUploadInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetExpire(v int64) *GetUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetFolder(v string) *GetUploadInfoResponseBody {
	s.Folder = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetHost(v string) *GetUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetHttpStatusCode(v int32) *GetUploadInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetKey(v string) *GetUploadInfoResponseBody {
	s.Key = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetMsg(v string) *GetUploadInfoResponseBody {
	s.Msg = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetName(v string) *GetUploadInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetPolicy(v string) *GetUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetRequestId(v string) *GetUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetSignature(v string) *GetUploadInfoResponseBody {
	s.Signature = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetSuccess(v bool) *GetUploadInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetUploadInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
