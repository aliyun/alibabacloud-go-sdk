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
	// The data returned.
	Data *GetOssScanConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E10BAF1C-A6C5-51E2-866C-76D5922E****
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
	// Indicates whether the prefixes of all objects are matched.
	//
	// example:
	//
	// true
	AllKeyPrefix *bool `json:"AllKeyPrefix,omitempty" xml:"AllKeyPrefix,omitempty"`
	// The number of buckets.
	//
	// example:
	//
	// 1
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// The name of the bucket.
	//
	// example:
	//
	// hz-new01****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The names of the buckets.
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
	// Indicates whether the check policy is enabled. Valid values:
	//
	// 	- **1**: enabled.
	//
	// 	- **0**: disabled.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end time of the check. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 1274****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The prefixes of the objects.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The suffixes of the objects that are checked.
	KeySuffixList []*string `json:"KeySuffixList,omitempty" xml:"KeySuffixList,omitempty" type:"Repeated"`
	// The timestamp when the object was last modified. The time must be later than the timestamp that you specify. Unit: milliseconds.
	//
	// example:
	//
	// 1724301769834
	LastModifiedStartTime *int64 `json:"LastModifiedStartTime,omitempty" xml:"LastModifiedStartTime,omitempty"`
	// The timestamp when the configuration was last modified.
	//
	// example:
	//
	// 1702025633079
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test0104
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Whether to enable real-time incremental detection. When this parameter is set to true, the parameters ScanDayList, StartTime, and EndTime are not effective.
	//
	// example:
	//
	// true
	RealTimeIncr *bool `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	// The days when the check is performed. The value indicates the days of the week.
	ScanDayList []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
	// The start time of the check. The time is in the HH:mm:ss format.
	//
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

func (s *GetOssScanConfigResponseBodyData) GetBucketNameList() []*string {
	return s.BucketNameList
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

func (s *GetOssScanConfigResponseBodyData) SetBucketNameList(v []*string) *GetOssScanConfigResponseBodyData {
	s.BucketNameList = v
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
