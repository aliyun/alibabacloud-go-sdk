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
}

type CreateOssBucketScanTaskRequest struct {
	// Specifies whether to match the prefixes of all objects.
	//
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// The names of the buckets.
	//
	// This parameter is required.
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
	// The maximum number of objects that can be extracted during decompression. Valid values: 1 to 1000. If the maximum number of objects that can be extracted is reached, the decompression operation immediately ends and the detection of extracted objects is not affected.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum number of decompression levels when multi-level packages are decompressed. Valid values: 1 to 5. If the maximum number of decompression levels is reached, the decompression operation immediately ends and the detection of extracted objects is not affected.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The decryption methods.
	DecryptionList []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	// The suffixes of the objects that you do not want to check.
	ExcludeKeySuffixList []*string `json:"ExcludeKeySuffixList,omitempty" xml:"ExcludeKeySuffixList,omitempty" type:"Repeated"`
	// The prefixes of the objects.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The suffixes of the objects that you want to check.
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// The timestamp when the object was last modified. The time must be later than the timestamp that you specify. Unit: milliseconds.
	//
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// The check mode. Valid values:
	//
	// 	- **1**: checks all objects in the bucket.
	//
	// 	- **2**: checks only new objects in the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ScanMode *int32 `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
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

func (s *CreateOssBucketScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
