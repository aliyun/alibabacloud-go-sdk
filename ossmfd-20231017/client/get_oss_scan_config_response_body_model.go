// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetOssScanConfigResponseBodyData) *GetOssScanConfigResponseBody
	GetData() *GetOssScanConfigResponseBodyData
	SetRequestId(v string) *GetOssScanConfigResponseBody
	GetRequestId() *string
}

type GetOssScanConfigResponseBody struct {
	Data *GetOssScanConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E14CECBF-8CAC-520C-ACC3-9503D5******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssScanConfigResponseBody) GetData() *GetOssScanConfigResponseBodyData {
	return s.Data
}

func (s *GetOssScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssScanConfigResponseBody) SetData(v *GetOssScanConfigResponseBodyData) *GetOssScanConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetOssScanConfigResponseBody) SetRequestId(v string) *GetOssScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssScanConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssScanConfigResponseBodyData struct {
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// example:
	//
	// 10
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// example:
	//
	// testBucket******
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 1
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
	// 01:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 13******
	Id            *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// example:
	//
	// 1702025633079
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// example:
	//
	// 1702025633079
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
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

func (s GetOssScanConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOssScanConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssScanConfigResponseBodyData) GetAllKeyPrefix() *bool {
	return s.AllKeyPrefix
}

func (s *GetOssScanConfigResponseBodyData) GetBucketCount() *int32 {
	return s.BucketCount
}

func (s *GetOssScanConfigResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetOssScanConfigResponseBodyData) GetDecompressMaxFileCount() *int32 {
	return s.DecompressMaxFileCount
}

func (s *GetOssScanConfigResponseBodyData) GetDecompressMaxLayer() *int32 {
	return s.DecompressMaxLayer
}

func (s *GetOssScanConfigResponseBodyData) GetDecryptionList() []*string {
	return s.DecryptionList
}

func (s *GetOssScanConfigResponseBodyData) GetEnable() *int32 {
	return s.Enable
}

func (s *GetOssScanConfigResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOssScanConfigResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetOssScanConfigResponseBodyData) GetKeyPrefixList() []*string {
	return s.KeyPrefixList
}

func (s *GetOssScanConfigResponseBodyData) GetKeySuffixList() []*string {
	return s.KeySuffixList
}

func (s *GetOssScanConfigResponseBodyData) GetLastModifiedStartTime() *int64 {
	return s.LastModifiedStartTime
}

func (s *GetOssScanConfigResponseBodyData) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *GetOssScanConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetOssScanConfigResponseBodyData) GetRealTimeIncr() *bool {
	return s.RealTimeIncr
}

func (s *GetOssScanConfigResponseBodyData) GetScanDayList() []*int32 {
	return s.ScanDayList
}

func (s *GetOssScanConfigResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOssScanConfigResponseBodyData) SetAllKeyPrefix(v bool) *GetOssScanConfigResponseBodyData {
	s.AllKeyPrefix = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetBucketCount(v int32) *GetOssScanConfigResponseBodyData {
	s.BucketCount = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetBucketName(v string) *GetOssScanConfigResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetDecompressMaxFileCount(v int32) *GetOssScanConfigResponseBodyData {
	s.DecompressMaxFileCount = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetDecompressMaxLayer(v int32) *GetOssScanConfigResponseBodyData {
	s.DecompressMaxLayer = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetDecryptionList(v []*string) *GetOssScanConfigResponseBodyData {
	s.DecryptionList = v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetEnable(v int32) *GetOssScanConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetEndTime(v string) *GetOssScanConfigResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetId(v string) *GetOssScanConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetKeyPrefixList(v []*string) *GetOssScanConfigResponseBodyData {
	s.KeyPrefixList = v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetKeySuffixList(v []*string) *GetOssScanConfigResponseBodyData {
	s.KeySuffixList = v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetLastModifiedStartTime(v int64) *GetOssScanConfigResponseBodyData {
	s.LastModifiedStartTime = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetLastUpdateTime(v int64) *GetOssScanConfigResponseBodyData {
	s.LastUpdateTime = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetName(v string) *GetOssScanConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetRealTimeIncr(v bool) *GetOssScanConfigResponseBodyData {
	s.RealTimeIncr = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetScanDayList(v []*int32) *GetOssScanConfigResponseBodyData {
	s.ScanDayList = v
	return s
}

func (s *GetOssScanConfigResponseBodyData) SetStartTime(v string) *GetOssScanConfigResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetOssScanConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
