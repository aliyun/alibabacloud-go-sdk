// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappUploadAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatappUploadAuthorizationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatappUploadAuthorizationResponseBody
	GetCode() *string
	SetData(v *GetChatappUploadAuthorizationResponseBodyData) *GetChatappUploadAuthorizationResponseBody
	GetData() *GetChatappUploadAuthorizationResponseBodyData
	SetMessage(v string) *GetChatappUploadAuthorizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatappUploadAuthorizationResponseBody
	GetRequestId() *string
}

type GetChatappUploadAuthorizationResponseBody struct {
	// Access denied for detailed information.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
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
	// The returned data.
	Data *GetChatappUploadAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappUploadAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappUploadAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatappUploadAuthorizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappUploadAuthorizationResponseBody) GetData() *GetChatappUploadAuthorizationResponseBodyData {
	return s.Data
}

func (s *GetChatappUploadAuthorizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappUploadAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappUploadAuthorizationResponseBody) SetAccessDeniedDetail(v string) *GetChatappUploadAuthorizationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetCode(v string) *GetChatappUploadAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetData(v *GetChatappUploadAuthorizationResponseBodyData) *GetChatappUploadAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetMessage(v string) *GetChatappUploadAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) SetRequestId(v string) *GetChatappUploadAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatappUploadAuthorizationResponseBodyData struct {
	// The AccessKey ID that is used to authorize a user to upload a file to Object Storage Service (OSS).
	//
	// example:
	//
	// 2skeuurfj****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret that is used to authorize a user to upload a file to OSS.
	//
	// example:
	//
	// skdkdukeuuuu****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The name of the bucket to which a file is uploaded in OSS.
	//
	// example:
	//
	// oss
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The directory to which the file is uploaded in Object Storage Service (OSS).
	//
	// example:
	//
	// 1000102939
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The address of the OSS server to which a file is uploaded.
	//
	// example:
	//
	// https://oss.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 3600
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The security token.
	//
	// example:
	//
	// dkdieiii**
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetChatappUploadAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatappUploadAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetDir() *string {
	return s.Dir
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetEndPoint() *string {
	return s.EndPoint
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetExpire() *int32 {
	return s.Expire
}

func (s *GetChatappUploadAuthorizationResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetAccessKeyId(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetAccessKeySecret(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetBucketName(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetDir(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.Dir = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetEndPoint(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetExpire(v int32) *GetChatappUploadAuthorizationResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) SetSecurityToken(v string) *GetChatappUploadAuthorizationResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetChatappUploadAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
