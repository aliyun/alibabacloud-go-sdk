// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDecompress(v bool) *CreateFileDetectRequest
	GetDecompress() *bool
	SetDecompressMaxFileCount(v int32) *CreateFileDetectRequest
	GetDecompressMaxFileCount() *int32
	SetDecompressMaxLayer(v int32) *CreateFileDetectRequest
	GetDecompressMaxLayer() *int32
	SetDownloadUrl(v string) *CreateFileDetectRequest
	GetDownloadUrl() *string
	SetHashKey(v string) *CreateFileDetectRequest
	GetHashKey() *string
	SetOssKey(v string) *CreateFileDetectRequest
	GetOssKey() *string
	SetSourceIp(v string) *CreateFileDetectRequest
	GetSourceIp() *string
	SetType(v int32) *CreateFileDetectRequest
	GetType() *int32
}

type CreateFileDetectRequest struct {
	// Specifies whether to identify and decompress compressed files. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	Decompress *bool `json:"Decompress,omitempty" xml:"Decompress,omitempty"`
	// The maximum number of files to decompress. Maximum value: 1000.
	//
	// This parameter is required when Decompress is set to true.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum number of decompression layers when compressed files are nested within a compressed package. Maximum value: 5.
	//
	// This parameter is required when Decompress is set to true.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The download URL of the file. You can pass in a file download URL (public URL) to directly trigger file detection without uploading the file in advance.
	//
	// example:
	//
	// https://xxxxxxxx.oss-cn-hangzhou-1.aliyuncs.com/xxxxx/xxxxxxxxxxxxxx?Expires=1671448125&OSSAccessKeyId=xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The unique identifier of the file. This parameter is required and must be the MD5 or SHA-256 of the file.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The storage key of the file in the OSS bucket.
	//
	// If you push the file for detection by using DownloadUrl, this parameter is optional. This parameter is obtained from the [CreateFileDetectUploadUrl](~~CreateFileDetectUploadUrl~~) operation.
	//
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 115.213.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of file to detect. Valid values:
	//
	// - **0**: malicious file detection
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFileDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectRequest) GoString() string {
	return s.String()
}

func (s *CreateFileDetectRequest) GetDecompress() *bool {
	return s.Decompress
}

func (s *CreateFileDetectRequest) GetDecompressMaxFileCount() *int32 {
	return s.DecompressMaxFileCount
}

func (s *CreateFileDetectRequest) GetDecompressMaxLayer() *int32 {
	return s.DecompressMaxLayer
}

func (s *CreateFileDetectRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateFileDetectRequest) GetHashKey() *string {
	return s.HashKey
}

func (s *CreateFileDetectRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateFileDetectRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateFileDetectRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateFileDetectRequest) SetDecompress(v bool) *CreateFileDetectRequest {
	s.Decompress = &v
	return s
}

func (s *CreateFileDetectRequest) SetDecompressMaxFileCount(v int32) *CreateFileDetectRequest {
	s.DecompressMaxFileCount = &v
	return s
}

func (s *CreateFileDetectRequest) SetDecompressMaxLayer(v int32) *CreateFileDetectRequest {
	s.DecompressMaxLayer = &v
	return s
}

func (s *CreateFileDetectRequest) SetDownloadUrl(v string) *CreateFileDetectRequest {
	s.DownloadUrl = &v
	return s
}

func (s *CreateFileDetectRequest) SetHashKey(v string) *CreateFileDetectRequest {
	s.HashKey = &v
	return s
}

func (s *CreateFileDetectRequest) SetOssKey(v string) *CreateFileDetectRequest {
	s.OssKey = &v
	return s
}

func (s *CreateFileDetectRequest) SetSourceIp(v string) *CreateFileDetectRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateFileDetectRequest) SetType(v int32) *CreateFileDetectRequest {
	s.Type = &v
	return s
}

func (s *CreateFileDetectRequest) Validate() error {
	return dara.Validate(s)
}
