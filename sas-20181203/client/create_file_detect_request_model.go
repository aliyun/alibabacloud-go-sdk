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
	// Specifies whether to decompress the archive for detection. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > This parameter is not supported when `Type` is set to `6`.
	//
	// example:
	//
	// false
	Decompress *bool `json:"Decompress,omitempty" xml:"Decompress,omitempty"`
	// The maximum number of files that can be decompressed from an archive. The maximum value is 1000.
	//
	// This parameter is required if you set `Decompress` to `true`.
	//
	// > This parameter is not supported when `Type` is set to `6`.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum number of decompression layers for nested archives. The maximum value is 5.
	//
	// This parameter is required if you set `Decompress` to `true`.
	//
	// > This parameter is not supported when `Type` is set to `6`.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The download link for the file. You can provide a public URL to trigger file detection without uploading the file.
	//
	// > Skill archives can be submitted only by providing a download link. Therefore, this parameter is required when `Type` is set to `6`.
	//
	// example:
	//
	// https://xxxxxxxx.oss-cn-hangzhou-1.aliyuncs.com/xxxxx/xxxxxxxxxxxxxx?Expires=1671448125&OSSAccessKeyId=xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The unique identifier of the file.
	//
	// This parameter is required if `Type` is `0`. Its value must be the MD5 or SHA-256 hash of the file.
	//
	// If you set `Type` to `6`, you do not need to specify this parameter. The operation returns the file\\"s unique identifier in the response.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The storage key of the file in an Object Storage Service (OSS) bucket.
	//
	// If you submit the file by using the `DownloadUrl` parameter, you can leave this parameter empty. To obtain the value of this parameter, call the [CreateFileDetectUploadUrl](~~CreateFileDetectUploadUrl~~) operation.
	//
	// > This parameter is not supported when `Type` is set to `6`.
	//
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The IP address of the source.
	//
	// example:
	//
	// 115.213.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the file to detect. Valid values:
	//
	// - **0**: Malicious file detection
	//
	// - **6**: Skill archive detection
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
