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
	// Whether to decompress or not. Valid values:
	//
	// - true: To decompress.
	//
	// - false: Not to decompress.
	//
	// example:
	//
	// false
	Decompress *bool `json:"Decompress,omitempty" xml:"Decompress,omitempty"`
	// The maximum number of files for decompression. The minimum value is 1, and the maximum value is 1000. If the decompression level exceeds the maximum, the decompression operation will be terminated, but the detection of decompressed files will not be affected.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum level of decompression when dealing with nested compressed files with multiple levels. The minimum value is 1, and the maximum value is 5. If the decompression level exceeds the maximum, the decompression operation will be terminated, but the detection of decompressed files will not be affected.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The URL that is used to download the file. You can specify this parameter to trigger file detection without the need to upload the file in advance.
	//
	// example:
	//
	// https://xxxxxxxx.oss-cn-hangzhou-1.aliyuncs.com/xxxxx/xxxxxxxxxxxxxx?Expires=1671448125&OSSAccessKeyId=xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The identifier of the file. Only MD5 hash values are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The key of the file that is stored in the Object Storage Service (OSS) bucket. You can call the [CreateFileDetectUploadUrl](~~CreateFileDetectUploadUrl~~) operation to query the keys of files.
	//
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 115.213.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the file. Valid values:
	//
	// 	- **0**: unknown files
	//
	// 	- **1**: binary files
	//
	// 	- **2**: webshell files
	//
	// 	- **4**: script files
	//
	// >  If you do not know the type of the file, set this parameter to 0.
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
