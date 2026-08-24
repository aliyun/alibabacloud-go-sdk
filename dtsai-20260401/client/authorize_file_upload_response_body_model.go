// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *AuthorizeFileUploadResponseBody
	GetAccessKeyId() *string
	SetBucket(v string) *AuthorizeFileUploadResponseBody
	GetBucket() *string
	SetEncodedPolicy(v string) *AuthorizeFileUploadResponseBody
	GetEncodedPolicy() *string
	SetEndpoint(v string) *AuthorizeFileUploadResponseBody
	GetEndpoint() *string
	SetErrorCode(v string) *AuthorizeFileUploadResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AuthorizeFileUploadResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *AuthorizeFileUploadResponseBody
	GetHttpStatusCode() *int32
	SetObjectKey(v string) *AuthorizeFileUploadResponseBody
	GetObjectKey() *string
	SetRequestId(v string) *AuthorizeFileUploadResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *AuthorizeFileUploadResponseBody
	GetSecurityToken() *string
	SetSignature(v string) *AuthorizeFileUploadResponseBody
	GetSignature() *string
	SetSuccess(v bool) *AuthorizeFileUploadResponseBody
	GetSuccess() *bool
}

type AuthorizeFileUploadResponseBody struct {
	// The temporary AccessKey ID used for OSS PostObject.
	//
	// example:
	//
	// STS.NV5xxx
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The destination OSS bucket.
	//
	// example:
	//
	// dts-ai-upload-cn-beijing7500163e0eae09
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The Base64-encoded PostObject policy, which includes the ObjectKey and file size limits.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjo...
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	// OSS Endpoint
	//
	// example:
	//
	// oss-cn-beijing.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter [Query] is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The business-level HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The object key that must be used as-is for the upload.
	//
	// example:
	//
	// 0a1b_123456_0123456789abcdef0123456789abcdef
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A30D0930-xxxx-xxxx-xxxx-C2C661CC8B58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Security Token Service (STS) token.
	//
	// example:
	//
	// CAISxxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The policy signature.
	//
	// example:
	//
	// masked-signature
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeFileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadResponseBody) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AuthorizeFileUploadResponseBody) GetBucket() *string {
	return s.Bucket
}

func (s *AuthorizeFileUploadResponseBody) GetEncodedPolicy() *string {
	return s.EncodedPolicy
}

func (s *AuthorizeFileUploadResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *AuthorizeFileUploadResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AuthorizeFileUploadResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AuthorizeFileUploadResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AuthorizeFileUploadResponseBody) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *AuthorizeFileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeFileUploadResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AuthorizeFileUploadResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *AuthorizeFileUploadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeFileUploadResponseBody) SetAccessKeyId(v string) *AuthorizeFileUploadResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetBucket(v string) *AuthorizeFileUploadResponseBody {
	s.Bucket = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetEncodedPolicy(v string) *AuthorizeFileUploadResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetEndpoint(v string) *AuthorizeFileUploadResponseBody {
	s.Endpoint = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetErrorCode(v string) *AuthorizeFileUploadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetErrorMessage(v string) *AuthorizeFileUploadResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetHttpStatusCode(v int32) *AuthorizeFileUploadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetObjectKey(v string) *AuthorizeFileUploadResponseBody {
	s.ObjectKey = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetRequestId(v string) *AuthorizeFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetSecurityToken(v string) *AuthorizeFileUploadResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetSignature(v string) *AuthorizeFileUploadResponseBody {
	s.Signature = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetSuccess(v bool) *AuthorizeFileUploadResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
