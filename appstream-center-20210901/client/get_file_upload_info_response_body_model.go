// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileUploadInfoResponseBodyData) *GetFileUploadInfoResponseBody
	GetData() *GetFileUploadInfoResponseBodyData
	SetRequestId(v string) *GetFileUploadInfoResponseBody
	GetRequestId() *string
}

type GetFileUploadInfoResponseBody struct {
	// Returns None.
	Data *GetFileUploadInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBody) GetData() *GetFileUploadInfoResponseBodyData {
	return s.Data
}

func (s *GetFileUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileUploadInfoResponseBody) SetData(v *GetFileUploadInfoResponseBodyData) *GetFileUploadInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetRequestId(v string) *GetFileUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileUploadInfoResponseBodyData struct {
	// The temporary AccessKey ID returned by Security Token Service (STS).
	//
	// example:
	//
	// LTA****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// appstream-*
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 600
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The folder path.
	//
	// example:
	//
	// cn-shanghai/aig_upm/***
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The maximum file size in bytes.
	//
	// example:
	//
	// 52428800
	MaxFileSize *int64 `json:"MaxFileSize,omitempty" xml:"MaxFileSize,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	OssPoint *string `json:"OssPoint,omitempty" xml:"OssPoint,omitempty"`
	// The PostObject policy (Base64-encoded).
	//
	// example:
	//
	// ***
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F591F0EA-AA10-52D2-ADA3-68397887B17C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The PostObject policy signature (HMAC-SHA1).
	//
	// example:
	//
	// ****************************
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The temporary token returned by STS.
	//
	// example:
	//
	// C*****
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s GetFileUploadInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetFileUploadInfoResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetFileUploadInfoResponseBodyData) GetExpiration() *string {
	return s.Expiration
}

func (s *GetFileUploadInfoResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetFileUploadInfoResponseBodyData) GetMaxFileSize() *int64 {
	return s.MaxFileSize
}

func (s *GetFileUploadInfoResponseBodyData) GetOssPoint() *string {
	return s.OssPoint
}

func (s *GetFileUploadInfoResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetFileUploadInfoResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileUploadInfoResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetFileUploadInfoResponseBodyData) GetStsToken() *string {
	return s.StsToken
}

func (s *GetFileUploadInfoResponseBodyData) SetAccessKeyId(v string) *GetFileUploadInfoResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetBucketName(v string) *GetFileUploadInfoResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetExpiration(v string) *GetFileUploadInfoResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetFilePath(v string) *GetFileUploadInfoResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetMaxFileSize(v int64) *GetFileUploadInfoResponseBodyData {
	s.MaxFileSize = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetOssPoint(v string) *GetFileUploadInfoResponseBodyData {
	s.OssPoint = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetPolicy(v string) *GetFileUploadInfoResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetRequestId(v string) *GetFileUploadInfoResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetSignature(v string) *GetFileUploadInfoResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) SetStsToken(v string) *GetFileUploadInfoResponseBodyData {
	s.StsToken = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
