// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestUploadTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *RequestUploadTokenResponseBody
	GetArgs() map[string]interface{}
	SetErrorCode(v int32) *RequestUploadTokenResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *RequestUploadTokenResponseBody
	GetMessage() *string
	SetModel(v *RequestUploadTokenResponseBodyModel) *RequestUploadTokenResponseBody
	GetModel() *RequestUploadTokenResponseBodyModel
	SetRequestId(v string) *RequestUploadTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RequestUploadTokenResponseBody
	GetSuccess() *bool
}

type RequestUploadTokenResponseBody struct {
	// Args
	Args map[string]interface{} `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *RequestUploadTokenResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// AB8AB5EC-9636-421D-AE7C-BB227DFC95B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestUploadTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RequestUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RequestUploadTokenResponseBody) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *RequestUploadTokenResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *RequestUploadTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RequestUploadTokenResponseBody) GetModel() *RequestUploadTokenResponseBodyModel {
	return s.Model
}

func (s *RequestUploadTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RequestUploadTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RequestUploadTokenResponseBody) SetArgs(v map[string]interface{}) *RequestUploadTokenResponseBody {
	s.Args = v
	return s
}

func (s *RequestUploadTokenResponseBody) SetErrorCode(v int32) *RequestUploadTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestUploadTokenResponseBody) SetMessage(v string) *RequestUploadTokenResponseBody {
	s.Message = &v
	return s
}

func (s *RequestUploadTokenResponseBody) SetModel(v *RequestUploadTokenResponseBodyModel) *RequestUploadTokenResponseBody {
	s.Model = v
	return s
}

func (s *RequestUploadTokenResponseBody) SetRequestId(v string) *RequestUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestUploadTokenResponseBody) SetSuccess(v bool) *RequestUploadTokenResponseBody {
	s.Success = &v
	return s
}

func (s *RequestUploadTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type RequestUploadTokenResponseBodyModel struct {
	// OSS AccessKeyId
	//
	// example:
	//
	// STS.NXEGHKdjkdnINNgLiDE
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// OSS AccessKeySecret
	//
	// example:
	//
	// ikKgkNDGedInGEIngL
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// Szi9v92mHNikdknfe
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// symbolic-prod
	UploadBucket *string `json:"UploadBucket,omitempty" xml:"UploadBucket,omitempty"`
	// example:
	//
	// /335374903@iphoneos/
	UploadDir *string `json:"UploadDir,omitempty" xml:"UploadDir,omitempty"`
}

func (s RequestUploadTokenResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s RequestUploadTokenResponseBodyModel) GoString() string {
	return s.String()
}

func (s *RequestUploadTokenResponseBodyModel) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *RequestUploadTokenResponseBodyModel) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *RequestUploadTokenResponseBodyModel) GetEndpoint() *string {
	return s.Endpoint
}

func (s *RequestUploadTokenResponseBodyModel) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RequestUploadTokenResponseBodyModel) GetUploadBucket() *string {
	return s.UploadBucket
}

func (s *RequestUploadTokenResponseBodyModel) GetUploadDir() *string {
	return s.UploadDir
}

func (s *RequestUploadTokenResponseBodyModel) SetAccessKeyId(v string) *RequestUploadTokenResponseBodyModel {
	s.AccessKeyId = &v
	return s
}

func (s *RequestUploadTokenResponseBodyModel) SetAccessKeySecret(v string) *RequestUploadTokenResponseBodyModel {
	s.AccessKeySecret = &v
	return s
}

func (s *RequestUploadTokenResponseBodyModel) SetEndpoint(v string) *RequestUploadTokenResponseBodyModel {
	s.Endpoint = &v
	return s
}

func (s *RequestUploadTokenResponseBodyModel) SetSecurityToken(v string) *RequestUploadTokenResponseBodyModel {
	s.SecurityToken = &v
	return s
}

func (s *RequestUploadTokenResponseBodyModel) SetUploadBucket(v string) *RequestUploadTokenResponseBodyModel {
	s.UploadBucket = &v
	return s
}

func (s *RequestUploadTokenResponseBodyModel) SetUploadDir(v string) *RequestUploadTokenResponseBodyModel {
	s.UploadDir = &v
	return s
}

func (s *RequestUploadTokenResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
