// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioNoteUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAudioNoteUploadUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetAudioNoteUploadUrlResponseBody
	GetCode() *string
	SetData(v *GetAudioNoteUploadUrlResponseBodyData) *GetAudioNoteUploadUrlResponseBody
	GetData() *GetAudioNoteUploadUrlResponseBodyData
	SetMessage(v string) *GetAudioNoteUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAudioNoteUploadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAudioNoteUploadUrlResponseBody
	GetSuccess() *bool
}

type GetAudioNoteUploadUrlResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAudioNoteUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code description.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE339D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAudioNoteUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAudioNoteUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioNoteUploadUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAudioNoteUploadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAudioNoteUploadUrlResponseBody) GetData() *GetAudioNoteUploadUrlResponseBodyData {
	return s.Data
}

func (s *GetAudioNoteUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAudioNoteUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAudioNoteUploadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAudioNoteUploadUrlResponseBody) SetAccessDeniedDetail(v string) *GetAudioNoteUploadUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBody) SetCode(v string) *GetAudioNoteUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBody) SetData(v *GetAudioNoteUploadUrlResponseBodyData) *GetAudioNoteUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBody) SetMessage(v string) *GetAudioNoteUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBody) SetRequestId(v string) *GetAudioNoteUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBody) SetSuccess(v bool) *GetAudioNoteUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAudioNoteUploadUrlResponseBodyData struct {
	// The AccessKey ID used for signing.
	//
	// example:
	//
	// STS.NYMxfDw3GkXfvEmZHXXXXXX
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The expiration time of the authorization.
	//
	// example:
	//
	// 2026-06-13T12:19:13Z
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The storage path of the OSS file.
	//
	// example:
	//
	// audio/file/sample1
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The host address.
	//
	// example:
	//
	// //alicom-voice-ai-agent-xxxx.oss-cn-xxxx.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The key of the OSS file.
	//
	// example:
	//
	// audio/file/sample1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The maximum file size.
	//
	// example:
	//
	// 2560000
	MaxFileSize *int64 `json:"MaxFileSize,omitempty" xml:"MaxFileSize,omitempty"`
	// The HTTP method used for upload.
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The upload policy.
	//
	// example:
	//
	// text
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The authorization licensing key.
	//
	// example:
	//
	// text
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The signature of the temporary upload credential, used to verify legitimacy during upload.
	//
	// example:
	//
	// oss sig text
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The upload URL.
	//
	// example:
	//
	// //alicom-voice-ai-agent-xxxx.oss-cn-xxxx.aliyuncs.com
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s GetAudioNoteUploadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAudioNoteUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetMaxFileSize() *int64 {
	return s.MaxFileSize
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetMethod() *string {
	return s.Method
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetAudioNoteUploadUrlResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetAccessKeyId(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetExpire(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetFilePath(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetHost(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetKey(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetMaxFileSize(v int64) *GetAudioNoteUploadUrlResponseBodyData {
	s.MaxFileSize = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetMethod(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.Method = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetPolicy(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetSecurityToken(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetSignature(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) SetUploadUrl(v string) *GetAudioNoteUploadUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *GetAudioNoteUploadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
