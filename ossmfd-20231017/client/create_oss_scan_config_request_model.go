// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllKeyPrefix(v string) *CreateOssScanConfigRequest
	GetAllKeyPrefix() *string
	SetBucketName(v string) *CreateOssScanConfigRequest
	GetBucketName() *string
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
	// example:
	//
	// true
	AllKeyPrefix *string `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testBucketName****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 100
	DecompressMaxFileCount *int32 `json:"DecompressMaxFileCount,omitempty" xml:"DecompressMaxFileCount,omitempty"`
	// example:
	//
	// 1
	DecompressMaxLayer *int32    `json:"DecompressMaxLayer,omitempty" xml:"DecompressMaxLayer,omitempty"`
	DecryptionList     []*string `json:"DecryptionList,omitempty" xml:"DecryptionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 00:10:00
	EndTime       *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// example:
	//
	// testName******
	Name         *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	RealTimeIncr *bool    `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	ScanDayList  []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
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

func (s *CreateOssScanConfigRequest) GetAllKeyPrefix() *string {
	return s.AllKeyPrefix
}

func (s *CreateOssScanConfigRequest) GetBucketName() *string {
	return s.BucketName
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

func (s *CreateOssScanConfigRequest) SetAllKeyPrefix(v string) *CreateOssScanConfigRequest {
	s.AllKeyPrefix = &v
	return s
}

func (s *CreateOssScanConfigRequest) SetBucketName(v string) *CreateOssScanConfigRequest {
	s.BucketName = &v
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
