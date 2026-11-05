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
	SetAutoAdd(v int32) *CreateOssScanConfigRequest
	GetAutoAdd() *int32
	SetBucketNameList(v []*string) *CreateOssScanConfigRequest
	GetBucketNameList() []*string
	SetClientToken(v string) *CreateOssScanConfigRequest
	GetClientToken() *string
	SetDecompressMaxFileCount(v int32) *CreateOssScanConfigRequest
	GetDecompressMaxFileCount() *int32
	SetDecompressMaxLayer(v int32) *CreateOssScanConfigRequest
	GetDecompressMaxLayer() *int32
	SetDecryptionList(v []*string) *CreateOssScanConfigRequest
	GetDecryptionList() []*string
	SetDryRun(v bool) *CreateOssScanConfigRequest
	GetDryRun() *bool
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
	SetSource(v string) *CreateOssScanConfigRequest
	GetSource() *string
	SetStartTime(v string) *CreateOssScanConfigRequest
	GetStartTime() *string
}

type CreateOssScanConfigRequest struct {
	// Specifies whether to match all prefixes. If this parameter is set to true, the KeyPrefixList parameter does not take effect.
	//
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// Specifies whether to add new OSS buckets to this policy by default. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// 0
	AutoAdd *int32 `json:"AutoAdd,omitempty" xml:"AutoAdd,omitempty"`
	// The list of bucket names.
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of files to decompress. Minimum value: 1. Maximum value: 1000. If the number of files exceeds this value, decompression stops, but the scan of already decompressed files is not affected.
	//
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// The maximum number of decompression layers for nested archives. Minimum value: 1. Maximum value: 5. If the nesting depth exceeds this value, decompression stops, but the scan of already decompressed files is not affected.
	//
	// example:
	//
	// 1
	DecompressMaxLayer *int32 `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	// The list of decryption types.
	DecryptionList []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without executing the actual operation.
	//
	// - **false**: executes the actual operation.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable the policy. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end time of the scan, in the HH:mm:ss format.
	//
	// example:
	//
	// 01:01:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The file prefix list.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The list of file name extensions to scan.
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// The timestamp after which the files to be scanned were last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable real-time incremental scanning. If this parameter is set to true, the ScanDayList, StartTime, and EndTime parameters do not take effect.
	//
	// example:
	//
	// true
	RealTimeIncr *bool `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	// The days of the week when the scan is performed, represented by numbers.
	ScanDayList []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
	// The source service. Valid values:
	//
	// - **OSS**: OSS.
	//
	// - **NAS**: NAS.
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time of the scan, in the HH:mm:ss format.
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

func (s *CreateOssScanConfigRequest) GetAutoAdd() *int32 {
	return s.AutoAdd
}

func (s *CreateOssScanConfigRequest) GetBucketNameList() []*string {
	return s.BucketNameList
}

func (s *CreateOssScanConfigRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateOssScanConfigRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateOssScanConfigRequest) GetSource() *string {
	return s.Source
}

func (s *CreateOssScanConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateOssScanConfigRequest) SetAllKeyPrefix(v bool) *CreateOssScanConfigRequest {
	s.AllKeyPrefix = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetAutoAdd(v int32) *CreateOssScanConfigRequest {
	s.AutoAdd = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetBucketNameList(v []*string) *CreateOssScanConfigRequest {
	s.BucketNameList = v
	return s
}

func (s *CreateOssScanConfigRequest) SetClientToken(v string) *CreateOssScanConfigRequest {
	s.ClientToken = &v
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

func (s *CreateOssScanConfigRequest) SetDryRun(v bool) *CreateOssScanConfigRequest {
	s.DryRun = &v
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

func (s *CreateOssScanConfigRequest) SetSource(v string) *CreateOssScanConfigRequest {
	s.Source = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetStartTime(v string) *CreateOssScanConfigRequest {
	s.StartTime = &v
	return s
}

func (s *CreateOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
