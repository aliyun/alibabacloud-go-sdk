// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllKeyPrefix(v bool) *UpdateOssScanConfigRequest
	GetAllKeyPrefix() *bool
	SetBucketNameList(v []*string) *UpdateOssScanConfigRequest
	GetBucketNameList() []*string
	SetDecompressMaxFileCount(v int32) *UpdateOssScanConfigRequest
	GetDecompressMaxFileCount() *int32
	SetDecompressMaxLayer(v int32) *UpdateOssScanConfigRequest
	GetDecompressMaxLayer() *int32
	SetDecryptionList(v []*string) *UpdateOssScanConfigRequest
	GetDecryptionList() []*string
	SetEnable(v int32) *UpdateOssScanConfigRequest
	GetEnable() *int32
	SetEndTime(v string) *UpdateOssScanConfigRequest
	GetEndTime() *string
	SetId(v string) *UpdateOssScanConfigRequest
	GetId() *string
	SetKeyPrefixList(v []*string) *UpdateOssScanConfigRequest
	GetKeyPrefixList() []*string
	SetKeySuffixList(v []*string) *UpdateOssScanConfigRequest
	GetKeySuffixList() []*string
	SetLastModifiedStartTime(v int64) *UpdateOssScanConfigRequest
	GetLastModifiedStartTime() *int64
	SetName(v string) *UpdateOssScanConfigRequest
	GetName() *string
	SetRealTimeIncr(v bool) *UpdateOssScanConfigRequest
	GetRealTimeIncr() *bool
	SetScanDayList(v []*int32) *UpdateOssScanConfigRequest
	GetScanDayList() []*int32
	SetStartTime(v string) *UpdateOssScanConfigRequest
	GetStartTime() *string
}

type UpdateOssScanConfigRequest struct {
	// Specifies whether to match all file prefixes.
	//
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// The list of bucket names.
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
	// The maximum number of files to decompress. Minimum value: 1. Maximum value: 1000. If the maximum number of decompressed files is exceeded, the decompression operation stops. The detection of already decompressed files is not affected.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum number of decompression layers for nested compressed files. Minimum value: 1. Maximum value: 5. If the maximum number of decompression layers is exceeded, the decompression operation stops. The detection of already decompressed files is not affected.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The list of decryption types.
	DecryptionList []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	// Specifies whether to enable the scan policy. Valid values:
	//
	// - **1**: Enable.
	//
	// - **0**: Disable.
	//
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The scan end time in the HH:mm:ss format.
	//
	// example:
	//
	// 00:00:01
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The scan policy ID.
	//
	// example:
	//
	// 1141****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The prefix list of files.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The list of file suffixes.
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// Scans files whose last modification time is after the specified timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// The scan policy name.
	//
	// example:
	//
	// testStrategy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable real-time incremental detection. If this parameter is set to true, the ScanDayList, StartTime, and EndTime parameters do not take effect.
	//
	// example:
	//
	// true
	RealTimeIncr *bool `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	// The scan days. The number indicates the day of the week.
	ScanDayList []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
	// The scan start time in the HH:mm:ss format.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateOssScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssScanConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssScanConfigRequest) GetAllKeyPrefix() *bool {
	return s.AllKeyPrefix
}

func (s *UpdateOssScanConfigRequest) GetBucketNameList() []*string {
	return s.BucketNameList
}

func (s *UpdateOssScanConfigRequest) GetDecompressMaxFileCount() *int32 {
	return s.DecompressMaxFileCount
}

func (s *UpdateOssScanConfigRequest) GetDecompressMaxLayer() *int32 {
	return s.DecompressMaxLayer
}

func (s *UpdateOssScanConfigRequest) GetDecryptionList() []*string {
	return s.DecryptionList
}

func (s *UpdateOssScanConfigRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateOssScanConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateOssScanConfigRequest) GetId() *string {
	return s.Id
}

func (s *UpdateOssScanConfigRequest) GetKeyPrefixList() []*string {
	return s.KeyPrefixList
}

func (s *UpdateOssScanConfigRequest) GetKeySuffixList() []*string {
	return s.KeySuffixList
}

func (s *UpdateOssScanConfigRequest) GetLastModifiedStartTime() *int64 {
	return s.LastModifiedStartTime
}

func (s *UpdateOssScanConfigRequest) GetName() *string {
	return s.Name
}

func (s *UpdateOssScanConfigRequest) GetRealTimeIncr() *bool {
	return s.RealTimeIncr
}

func (s *UpdateOssScanConfigRequest) GetScanDayList() []*int32 {
	return s.ScanDayList
}

func (s *UpdateOssScanConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateOssScanConfigRequest) SetAllKeyPrefix(v bool) *UpdateOssScanConfigRequest {
	s.AllKeyPrefix = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetBucketNameList(v []*string) *UpdateOssScanConfigRequest {
	s.BucketNameList = v
	return s
}

func (s *UpdateOssScanConfigRequest) SetDecompressMaxFileCount(v int32) *UpdateOssScanConfigRequest {
	s.DecompressMaxFileCount = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetDecompressMaxLayer(v int32) *UpdateOssScanConfigRequest {
	s.DecompressMaxLayer = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetDecryptionList(v []*string) *UpdateOssScanConfigRequest {
	s.DecryptionList = v
	return s
}

func (s *UpdateOssScanConfigRequest) SetEnable(v int32) *UpdateOssScanConfigRequest {
	s.Enable = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetEndTime(v string) *UpdateOssScanConfigRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetId(v string) *UpdateOssScanConfigRequest {
	s.Id = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetKeyPrefixList(v []*string) *UpdateOssScanConfigRequest {
	s.KeyPrefixList = v
	return s
}

func (s *UpdateOssScanConfigRequest) SetKeySuffixList(v []*string) *UpdateOssScanConfigRequest {
	s.KeySuffixList = v
	return s
}

func (s *UpdateOssScanConfigRequest) SetLastModifiedStartTime(v int64) *UpdateOssScanConfigRequest {
	s.LastModifiedStartTime = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetName(v string) *UpdateOssScanConfigRequest {
	s.Name = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetRealTimeIncr(v bool) *UpdateOssScanConfigRequest {
	s.RealTimeIncr = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetScanDayList(v []*int32) *UpdateOssScanConfigRequest {
	s.ScanDayList = v
	return s
}

func (s *UpdateOssScanConfigRequest) SetStartTime(v string) *UpdateOssScanConfigRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
