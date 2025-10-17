// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempAccessTokenIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TempAccessTokenIntlResponseBody
	GetCode() *string
	SetData(v *TempAccessTokenIntlResponseBodyData) *TempAccessTokenIntlResponseBody
	GetData() *TempAccessTokenIntlResponseBodyData
	SetMessage(v string) *TempAccessTokenIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *TempAccessTokenIntlResponseBody
	GetRequestId() *string
}

type TempAccessTokenIntlResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return result.
	Data *TempAccessTokenIntlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TempAccessTokenIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TempAccessTokenIntlResponseBody) GoString() string {
	return s.String()
}

func (s *TempAccessTokenIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *TempAccessTokenIntlResponseBody) GetData() *TempAccessTokenIntlResponseBodyData {
	return s.Data
}

func (s *TempAccessTokenIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TempAccessTokenIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TempAccessTokenIntlResponseBody) SetCode(v string) *TempAccessTokenIntlResponseBody {
	s.Code = &v
	return s
}

func (s *TempAccessTokenIntlResponseBody) SetData(v *TempAccessTokenIntlResponseBodyData) *TempAccessTokenIntlResponseBody {
	s.Data = v
	return s
}

func (s *TempAccessTokenIntlResponseBody) SetMessage(v string) *TempAccessTokenIntlResponseBody {
	s.Message = &v
	return s
}

func (s *TempAccessTokenIntlResponseBody) SetRequestId(v string) *TempAccessTokenIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempAccessTokenIntlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TempAccessTokenIntlResponseBodyData struct {
	// AccessKeyId for temporary file upload credentials.
	//
	// example:
	//
	// ***
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// Temporary authorization secret.
	//
	// example:
	//
	// 3hxuRpEJ3Jv2Rtzyg4HooFCYqps762XcNtzhn19wQymk
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// Bucket name.
	//
	// example:
	//
	// liyi
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// File prefix.
	//
	// example:
	//
	// test001
	FileNamePrefix *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	// OSS endpoint.
	//
	// example:
	//
	// ossEndPoint
	OssEndPoint *string `json:"OssEndPoint,omitempty" xml:"OssEndPoint,omitempty"`
	// Security token for temporary file upload credentials.
	//
	// example:
	//
	// ***
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s TempAccessTokenIntlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TempAccessTokenIntlResponseBodyData) GoString() string {
	return s.String()
}

func (s *TempAccessTokenIntlResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *TempAccessTokenIntlResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *TempAccessTokenIntlResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *TempAccessTokenIntlResponseBodyData) GetFileNamePrefix() *string {
	return s.FileNamePrefix
}

func (s *TempAccessTokenIntlResponseBodyData) GetOssEndPoint() *string {
	return s.OssEndPoint
}

func (s *TempAccessTokenIntlResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *TempAccessTokenIntlResponseBodyData) SetAccessKeyId(v string) *TempAccessTokenIntlResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *TempAccessTokenIntlResponseBodyData) SetAccessKeySecret(v string) *TempAccessTokenIntlResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *TempAccessTokenIntlResponseBodyData) SetBucketName(v string) *TempAccessTokenIntlResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *TempAccessTokenIntlResponseBodyData) SetFileNamePrefix(v string) *TempAccessTokenIntlResponseBodyData {
	s.FileNamePrefix = &v
	return s
}

func (s *TempAccessTokenIntlResponseBodyData) SetOssEndPoint(v string) *TempAccessTokenIntlResponseBodyData {
	s.OssEndPoint = &v
	return s
}

func (s *TempAccessTokenIntlResponseBodyData) SetSecurityToken(v string) *TempAccessTokenIntlResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *TempAccessTokenIntlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
