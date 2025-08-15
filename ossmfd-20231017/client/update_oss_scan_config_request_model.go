// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllKeyPrefix(v string) *UpdateOssScanConfigRequest
	GetAllKeyPrefix() *string
	SetBucketName(v string) *UpdateOssScanConfigRequest
	GetBucketName() *string
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
	// example:
	//
	// true
	AllKeyPrefix *string `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// example:
	//
	// testBucket******
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 1000
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// example:
	//
	// 1
	DecompressMaxLayer *int32    `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	DecryptionList     []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 00:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 100******
	Id            *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// example:
	//
	// testPolicy******
	Name         *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	RealTimeIncr *bool    `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	ScanDayList  []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
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

func (s *UpdateOssScanConfigRequest) GetAllKeyPrefix() *string {
	return s.AllKeyPrefix
}

func (s *UpdateOssScanConfigRequest) GetBucketName() *string {
	return s.BucketName
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

func (s *UpdateOssScanConfigRequest) SetAllKeyPrefix(v string) *UpdateOssScanConfigRequest {
	s.AllKeyPrefix = &v
	return s
}

func (s *UpdateOssScanConfigRequest) SetBucketName(v string) *UpdateOssScanConfigRequest {
	s.BucketName = &v
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
