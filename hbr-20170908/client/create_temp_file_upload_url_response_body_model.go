// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *CreateTempFileUploadUrlResponseBody
	GetBucketName() *string
	SetCode(v string) *CreateTempFileUploadUrlResponseBody
	GetCode() *string
	SetEndpoint(v string) *CreateTempFileUploadUrlResponseBody
	GetEndpoint() *string
	SetExpireTime(v int64) *CreateTempFileUploadUrlResponseBody
	GetExpireTime() *int64
	SetMessage(v string) *CreateTempFileUploadUrlResponseBody
	GetMessage() *string
	SetOssAccessKeyId(v string) *CreateTempFileUploadUrlResponseBody
	GetOssAccessKeyId() *string
	SetPolicy(v string) *CreateTempFileUploadUrlResponseBody
	GetPolicy() *string
	SetRequestId(v string) *CreateTempFileUploadUrlResponseBody
	GetRequestId() *string
	SetSignature(v string) *CreateTempFileUploadUrlResponseBody
	GetSignature() *string
	SetSuccess(v bool) *CreateTempFileUploadUrlResponseBody
	GetSuccess() *bool
	SetTempFileKey(v string) *CreateTempFileUploadUrlResponseBody
	GetTempFileKey() *string
}

type CreateTempFileUploadUrlResponseBody struct {
	// The name of the OSS bucket to which the file is uploaded.
	//
	// example:
	//
	// hbr-temp-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The endpoint that is used to upload the file to OSS.
	//
	// example:
	//
	// oss-cn-shenzhen.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The expiration time of the signature that is used to upload the file to OSS. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1654326678
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The AccessKey ID that is used to upload the file to OSS.
	//
	// example:
	//
	// LTAI****Up
	OssAccessKeyId *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	// The policy that is used to upload the file to OSS.
	//
	// example:
	//
	// eyJleH****V19
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F4EEB401-DD21-588D-AE3B-1E835C7655E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature that is used to upload the file to OSS.
	//
	// example:
	//
	// RmhI****0A=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The key that is used to upload the file to OSS.
	//
	// example:
	//
	// temp/1440155109798732/upload/2022-07-29/49bed34c-b430-4e7e-89b1-4be2b734f95c/iaclone.diff
	TempFileKey *string `json:"TempFileKey,omitempty" xml:"TempFileKey,omitempty"`
}

func (s CreateTempFileUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTempFileUploadUrlResponseBody) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateTempFileUploadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTempFileUploadUrlResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateTempFileUploadUrlResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *CreateTempFileUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTempFileUploadUrlResponseBody) GetOssAccessKeyId() *string {
	return s.OssAccessKeyId
}

func (s *CreateTempFileUploadUrlResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *CreateTempFileUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTempFileUploadUrlResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *CreateTempFileUploadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTempFileUploadUrlResponseBody) GetTempFileKey() *string {
	return s.TempFileKey
}

func (s *CreateTempFileUploadUrlResponseBody) SetBucketName(v string) *CreateTempFileUploadUrlResponseBody {
	s.BucketName = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetCode(v string) *CreateTempFileUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetEndpoint(v string) *CreateTempFileUploadUrlResponseBody {
	s.Endpoint = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetExpireTime(v int64) *CreateTempFileUploadUrlResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetMessage(v string) *CreateTempFileUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetOssAccessKeyId(v string) *CreateTempFileUploadUrlResponseBody {
	s.OssAccessKeyId = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetPolicy(v string) *CreateTempFileUploadUrlResponseBody {
	s.Policy = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetRequestId(v string) *CreateTempFileUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetSignature(v string) *CreateTempFileUploadUrlResponseBody {
	s.Signature = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetSuccess(v bool) *CreateTempFileUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetTempFileKey(v string) *CreateTempFileUploadUrlResponseBody {
	s.TempFileKey = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
