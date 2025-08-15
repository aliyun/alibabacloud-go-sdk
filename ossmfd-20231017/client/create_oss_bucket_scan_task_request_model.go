// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssBucketScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllKeyPrefix(v string) *CreateOssBucketScanTaskRequest
	GetAllKeyPrefix() *string
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
	// example:
	//
	// true
	AllKeyPrefix *string `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// This parameter is required.
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// example:
	//
	// 1
	DecompressMaxLayer   *int32    `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	DecryptionList       []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	ExcludeKeySuffixList []*string `json:"ExcludeKeySuffixList,omitempty" xml:"ExcludeKeySuffixList,omitempty" type:"Repeated"`
	KeyPrefixList        []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	KeySuffixList        []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
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

func (s *CreateOssBucketScanTaskRequest) GetAllKeyPrefix() *string {
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

func (s *CreateOssBucketScanTaskRequest) SetAllKeyPrefix(v string) *CreateOssBucketScanTaskRequest {
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
