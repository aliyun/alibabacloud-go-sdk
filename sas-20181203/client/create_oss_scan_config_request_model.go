// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllKeyPrefix(v bool) *CreateOssScanConfigRequest
	GetAllKeyPrefix() *bool
	SetBucketNameList(v []*string) *CreateOssScanConfigRequest
	GetBucketNameList() []*string
	SetDecompressMaxFileCount(v int32) *CreateOssScanConfigRequest
	GetDecompressMaxFileCount() *int32
	SetDecompressMaxLayer(v int32) *CreateOssScanConfigRequest
	GetDecompressMaxLayer() *int32
	SetDecryptionList(v []*string) *CreateOssScanConfigRequest
	GetDecryptionList() []*string
	SetEnable(v int32) *CreateOssScanConfigRequest
	GetEnable() *int32
	SetEndTime(v string) *CreateOssScanConfigRequest
	GetEndTime() *string
	SetKeyPrefixList(v []*string) *CreateOssScanConfigRequest
	GetKeyPrefixList() []*string
	SetKeySuffixList(v []*string) *CreateOssScanConfigRequest
	GetKeySuffixList() []*string
	SetLastModifiedStartTime(v int64) *CreateOssScanConfigRequest
	GetLastModifiedStartTime() *int64
	SetName(v string) *CreateOssScanConfigRequest
	GetName() *string
	SetRealTimeIncr(v bool) *CreateOssScanConfigRequest
	GetRealTimeIncr() *bool
	SetScanDayList(v []*int32) *CreateOssScanConfigRequest
	GetScanDayList() []*int32
	SetStartTime(v string) *CreateOssScanConfigRequest
	GetStartTime() *string
}

type CreateOssScanConfigRequest struct {
	// Specifies whether to match the prefixes of all objects.
	//
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// The names of buckets.
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
	// Specifies whether to enable the policy. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the scan ends. The time must be in the HH:mm:ss format.
	//
	// example:
	//
	// 01:01:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The prefixes of the objects.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The suffixes of the files to scan.
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// The timestamp when the object was last modified. The time must be later than the timestamp that you specify. Unit: milliseconds.
	//
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// runtime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Whether to enable real-time incremental detection. When this parameter is set to true, the parameters ScanDayList, StartTime, and EndTime are not effective.
	//
	// example:
	//
	// true
	RealTimeIncr *bool `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	// The days on which the scan is executed in a week.
	ScanDayList []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
	// The time when the scan starts. The time must be in the HH:mm:ss format.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateOssScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOssScanConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateOssScanConfigRequest) GetAllKeyPrefix() *bool {
	return s.AllKeyPrefix
}

func (s *CreateOssScanConfigRequest) GetBucketNameList() []*string {
	return s.BucketNameList
}

func (s *CreateOssScanConfigRequest) GetDecompressMaxFileCount() *int32 {
	return s.DecompressMaxFileCount
}

func (s *CreateOssScanConfigRequest) GetDecompressMaxLayer() *int32 {
	return s.DecompressMaxLayer
}

func (s *CreateOssScanConfigRequest) GetDecryptionList() []*string {
	return s.DecryptionList
}

func (s *CreateOssScanConfigRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *CreateOssScanConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateOssScanConfigRequest) GetKeyPrefixList() []*string {
	return s.KeyPrefixList
}

func (s *CreateOssScanConfigRequest) GetKeySuffixList() []*string {
	return s.KeySuffixList
}

func (s *CreateOssScanConfigRequest) GetLastModifiedStartTime() *int64 {
	return s.LastModifiedStartTime
}

func (s *CreateOssScanConfigRequest) GetName() *string {
	return s.Name
}

func (s *CreateOssScanConfigRequest) GetRealTimeIncr() *bool {
	return s.RealTimeIncr
}

func (s *CreateOssScanConfigRequest) GetScanDayList() []*int32 {
	return s.ScanDayList
}

func (s *CreateOssScanConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateOssScanConfigRequest) SetAllKeyPrefix(v bool) *CreateOssScanConfigRequest {
	s.AllKeyPrefix = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetBucketNameList(v []*string) *CreateOssScanConfigRequest {
	s.BucketNameList = v
	return s
}

func (s *CreateOssScanConfigRequest) SetDecompressMaxFileCount(v int32) *CreateOssScanConfigRequest {
	s.DecompressMaxFileCount = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetDecompressMaxLayer(v int32) *CreateOssScanConfigRequest {
	s.DecompressMaxLayer = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetDecryptionList(v []*string) *CreateOssScanConfigRequest {
	s.DecryptionList = v
	return s
}

func (s *CreateOssScanConfigRequest) SetEnable(v int32) *CreateOssScanConfigRequest {
	s.Enable = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetEndTime(v string) *CreateOssScanConfigRequest {
	s.EndTime = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetKeyPrefixList(v []*string) *CreateOssScanConfigRequest {
	s.KeyPrefixList = v
	return s
}

func (s *CreateOssScanConfigRequest) SetKeySuffixList(v []*string) *CreateOssScanConfigRequest {
	s.KeySuffixList = v
	return s
}

func (s *CreateOssScanConfigRequest) SetLastModifiedStartTime(v int64) *CreateOssScanConfigRequest {
	s.LastModifiedStartTime = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetName(v string) *CreateOssScanConfigRequest {
	s.Name = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetRealTimeIncr(v bool) *CreateOssScanConfigRequest {
	s.RealTimeIncr = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetScanDayList(v []*int32) *CreateOssScanConfigRequest {
	s.ScanDayList = v
	return s
}

func (s *CreateOssScanConfigRequest) SetStartTime(v string) *CreateOssScanConfigRequest {
	s.StartTime = &v
	return s
}

func (s *CreateOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
