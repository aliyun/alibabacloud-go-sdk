// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUploadInfoResponseBody
	GetCode() *string
	SetData(v *GetUploadInfoResponseBodyData) *GetUploadInfoResponseBody
	GetData() *GetUploadInfoResponseBodyData
	SetMessage(v string) *GetUploadInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUploadInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUploadInfoResponseBody
	GetSuccess() *bool
}

type GetUploadInfoResponseBody struct {
	// The return code of the API. A value of 200 indicates success. If the call fails, the corresponding error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file upload endpoint and temporary credentials.
	Data *GetUploadInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. It is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 8737D130-BFD0-5D51-96F6-C08EB1139A25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUploadInfoResponseBody) GetData() *GetUploadInfoResponseBodyData {
	return s.Data
}

func (s *GetUploadInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUploadInfoResponseBody) SetCode(v string) *GetUploadInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetData(v *GetUploadInfoResponseBodyData) *GetUploadInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadInfoResponseBody) SetMessage(v string) *GetUploadInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetRequestId(v string) *GetUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetSuccess(v bool) *GetUploadInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetUploadInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUploadInfoResponseBodyData struct {
	// The AccessKey ID of the STS temporary access credential.
	//
	// example:
	//
	// STS.NUexample****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey Secret of the STS temporary access credential. It is used together with AccessKeyId and StsToken.
	//
	// example:
	//
	// secret-example****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The name of the destination OSS bucket for the upload.
	//
	// example:
	//
	// appstream-example
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// Not supported.
	//
	// example:
	//
	// 123
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The OSS directory prefix where files are allowed to be uploaded. The path of the uploaded object must be located under this directory.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The OSS endpoint for public network access.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	OssPoint *string `json:"OssPoint,omitempty" xml:"OssPoint,omitempty"`
	// The internal OSS endpoint.
	//
	// example:
	//
	// oss-cn-shanghai-internal.aliyuncs.com
	OssPointInternal *string `json:"OssPointInternal,omitempty" xml:"OssPointInternal,omitempty"`
	// The STS temporary security token. It is used together with AccessKeyId and AccessKeySecret for file uploads.
	//
	// example:
	//
	// CAISexample****
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s GetUploadInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUploadInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetUploadInfoResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetUploadInfoResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetUploadInfoResponseBodyData) GetExpiration() *string {
	return s.Expiration
}

func (s *GetUploadInfoResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetUploadInfoResponseBodyData) GetOssPoint() *string {
	return s.OssPoint
}

func (s *GetUploadInfoResponseBodyData) GetOssPointInternal() *string {
	return s.OssPointInternal
}

func (s *GetUploadInfoResponseBodyData) GetStsToken() *string {
	return s.StsToken
}

func (s *GetUploadInfoResponseBodyData) SetAccessKeyId(v string) *GetUploadInfoResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetAccessKeySecret(v string) *GetUploadInfoResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetBucketName(v string) *GetUploadInfoResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetExpiration(v string) *GetUploadInfoResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetFilePath(v string) *GetUploadInfoResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetOssPoint(v string) *GetUploadInfoResponseBodyData {
	s.OssPoint = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetOssPointInternal(v string) *GetUploadInfoResponseBodyData {
	s.OssPointInternal = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) SetStsToken(v string) *GetUploadInfoResponseBodyData {
	s.StsToken = &v
	return s
}

func (s *GetUploadInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
