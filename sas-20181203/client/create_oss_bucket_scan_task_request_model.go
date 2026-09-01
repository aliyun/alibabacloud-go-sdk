// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssBucketScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllKeyPrefix(v bool) *CreateOssBucketScanTaskRequest
	GetAllKeyPrefix() *bool
	SetBucketNameList(v []*string) *CreateOssBucketScanTaskRequest
	GetBucketNameList() []*string
	SetDecompressMaxFileCount(v int32) *CreateOssBucketScanTaskRequest
	GetDecompressMaxFileCount() *int32
	SetDecompressMaxLayer(v int32) *CreateOssBucketScanTaskRequest
	GetDecompressMaxLayer() *int32
	SetDecryptionList(v []*string) *CreateOssBucketScanTaskRequest
	GetDecryptionList() []*string
	SetExcludeKeySuffixList(v []*string) *CreateOssBucketScanTaskRequest
	GetExcludeKeySuffixList() []*string
	SetKeyPrefixList(v []*string) *CreateOssBucketScanTaskRequest
	GetKeyPrefixList() []*string
	SetKeySuffixList(v []*string) *CreateOssBucketScanTaskRequest
	GetKeySuffixList() []*string
	SetLastModifiedStartTime(v int64) *CreateOssBucketScanTaskRequest
	GetLastModifiedStartTime() *int64
	SetScanMode(v int32) *CreateOssBucketScanTaskRequest
	GetScanMode() *int32
	SetSource(v string) *CreateOssBucketScanTaskRequest
	GetSource() *string
}

type CreateOssBucketScanTaskRequest struct {
	// Specifies whether to match all prefixes. If this parameter is set to true, the KeyPrefixList parameter does not take effect.
	//
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// The list of bucket names.
	//
	// This parameter is required.
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
	// The maximum number of files to decompress. The minimum value is 1 and the maximum value is 1000. When the maximum number of decompressed files is exceeded, the decompression operation ends immediately. The detection of files that have already been decompressed is not affected.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum number of decompression layers when multiple levels of compressed packages are nested. The minimum value is 1 and the maximum value is 5. When the maximum number of decompression layers is exceeded, the decompression operation ends immediately. The detection of files that have already been decompressed is not affected.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The list of decryption types.
	DecryptionList []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	// The list of file suffixes to exclude from detection.
	ExcludeKeySuffixList []*string `json:"ExcludeKeySuffixList,omitempty" xml:"ExcludeKeySuffixList,omitempty" type:"Repeated"`
	// The prefix list of files.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The list of file suffixes.
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// Specifies that only files whose last modification time is after the specified timestamp are detected. Unit: milliseconds.
	//
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// The detection mode. Valid values:
	//
	// - **1**: Full file detection.
	//
	// - **2**: Incremental file detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ScanMode *int32 `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The business source. Valid values:
	//
	// - **OSS**: OSS
	//
	// - **NAS**: NAS
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateOssBucketScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOssBucketScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOssBucketScanTaskRequest) GetAllKeyPrefix() *bool {
	return s.AllKeyPrefix
}

func (s *CreateOssBucketScanTaskRequest) GetBucketNameList() []*string {
	return s.BucketNameList
}

func (s *CreateOssBucketScanTaskRequest) GetDecompressMaxFileCount() *int32 {
	return s.DecompressMaxFileCount
}

func (s *CreateOssBucketScanTaskRequest) GetDecompressMaxLayer() *int32 {
	return s.DecompressMaxLayer
}

func (s *CreateOssBucketScanTaskRequest) GetDecryptionList() []*string {
	return s.DecryptionList
}

func (s *CreateOssBucketScanTaskRequest) GetExcludeKeySuffixList() []*string {
	return s.ExcludeKeySuffixList
}

func (s *CreateOssBucketScanTaskRequest) GetKeyPrefixList() []*string {
	return s.KeyPrefixList
}

func (s *CreateOssBucketScanTaskRequest) GetKeySuffixList() []*string {
	return s.KeySuffixList
}

func (s *CreateOssBucketScanTaskRequest) GetLastModifiedStartTime() *int64 {
	return s.LastModifiedStartTime
}

func (s *CreateOssBucketScanTaskRequest) GetScanMode() *int32 {
	return s.ScanMode
}

func (s *CreateOssBucketScanTaskRequest) GetSource() *string {
	return s.Source
}

func (s *CreateOssBucketScanTaskRequest) SetAllKeyPrefix(v bool) *CreateOssBucketScanTaskRequest {
	s.AllKeyPrefix = &v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetBucketNameList(v []*string) *CreateOssBucketScanTaskRequest {
	s.BucketNameList = v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetDecompressMaxFileCount(v int32) *CreateOssBucketScanTaskRequest {
	s.DecompressMaxFileCount = &v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetDecompressMaxLayer(v int32) *CreateOssBucketScanTaskRequest {
	s.DecompressMaxLayer = &v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetDecryptionList(v []*string) *CreateOssBucketScanTaskRequest {
	s.DecryptionList = v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetExcludeKeySuffixList(v []*string) *CreateOssBucketScanTaskRequest {
	s.ExcludeKeySuffixList = v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetKeyPrefixList(v []*string) *CreateOssBucketScanTaskRequest {
	s.KeyPrefixList = v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetKeySuffixList(v []*string) *CreateOssBucketScanTaskRequest {
	s.KeySuffixList = v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetLastModifiedStartTime(v int64) *CreateOssBucketScanTaskRequest {
	s.LastModifiedStartTime = &v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetScanMode(v int32) *CreateOssBucketScanTaskRequest {
	s.ScanMode = &v
	return s
}

func (s *CreateOssBucketScanTaskRequest) SetSource(v string) *CreateOssBucketScanTaskRequest {
	s.Source = &v
	return s
}

func (s *CreateOssBucketScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
