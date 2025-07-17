// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUploadTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeUploadTokenResponseBody
	GetCode() *int32
	SetData(v *DescribeUploadTokenResponseBodyData) *DescribeUploadTokenResponseBody
	GetData() *DescribeUploadTokenResponseBodyData
	SetMsg(v string) *DescribeUploadTokenResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeUploadTokenResponseBody
	GetRequestId() *string
}

type DescribeUploadTokenResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeUploadTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUploadTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUploadTokenResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeUploadTokenResponseBody) GetData() *DescribeUploadTokenResponseBodyData {
	return s.Data
}

func (s *DescribeUploadTokenResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeUploadTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUploadTokenResponseBody) SetCode(v int32) *DescribeUploadTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUploadTokenResponseBody) SetData(v *DescribeUploadTokenResponseBodyData) *DescribeUploadTokenResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUploadTokenResponseBody) SetMsg(v string) *DescribeUploadTokenResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeUploadTokenResponseBody) SetRequestId(v string) *DescribeUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUploadTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUploadTokenResponseBodyData struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NUEUjvDqMuvH6oQA1TXxxH4wVR
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// xxxx
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// oss-cip-shanghai
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The time when the file sharing link expires.
	//
	// example:
	//
	// 1720577200
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The file prefix.
	//
	// example:
	//
	// upload/1xxb89/
	FileNamePrefix *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	// the oss intranet point.
	//
	// example:
	//
	// https://oss-cn-shanghai-internal.aliyuncs.com
	OssInternalEndPoint *string `json:"OssInternalEndPoint,omitempty" xml:"OssInternalEndPoint,omitempty"`
	// the oss internet point.
	//
	// example:
	//
	// https://oss-cn-shanghai.aliyuncs.com
	OssInternetEndPoint *string `json:"OssInternetEndPoint,omitempty" xml:"OssInternetEndPoint,omitempty"`
	// The security token.
	//
	// example:
	//
	// xxxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUploadTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUploadTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUploadTokenResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeUploadTokenResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *DescribeUploadTokenResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeUploadTokenResponseBodyData) GetExpiration() *int32 {
	return s.Expiration
}

func (s *DescribeUploadTokenResponseBodyData) GetFileNamePrefix() *string {
	return s.FileNamePrefix
}

func (s *DescribeUploadTokenResponseBodyData) GetOssInternalEndPoint() *string {
	return s.OssInternalEndPoint
}

func (s *DescribeUploadTokenResponseBodyData) GetOssInternetEndPoint() *string {
	return s.OssInternetEndPoint
}

func (s *DescribeUploadTokenResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUploadTokenResponseBodyData) SetAccessKeyId(v string) *DescribeUploadTokenResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetAccessKeySecret(v string) *DescribeUploadTokenResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetBucketName(v string) *DescribeUploadTokenResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetExpiration(v int32) *DescribeUploadTokenResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetFileNamePrefix(v string) *DescribeUploadTokenResponseBodyData {
	s.FileNamePrefix = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetOssInternalEndPoint(v string) *DescribeUploadTokenResponseBodyData {
	s.OssInternalEndPoint = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetOssInternetEndPoint(v string) *DescribeUploadTokenResponseBodyData {
	s.OssInternetEndPoint = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetSecurityToken(v string) *DescribeUploadTokenResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
