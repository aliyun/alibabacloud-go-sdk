// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListOssScanConfigResponseBodyData) *ListOssScanConfigResponseBody
	GetData() []*ListOssScanConfigResponseBodyData
	SetPageInfo(v *ListOssScanConfigResponseBodyPageInfo) *ListOssScanConfigResponseBody
	GetPageInfo() *ListOssScanConfigResponseBodyPageInfo
	SetRequestId(v string) *ListOssScanConfigResponseBody
	GetRequestId() *string
}

type ListOssScanConfigResponseBody struct {
	// The data returned.
	Data []*ListOssScanConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListOssScanConfigResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E10BAF1C-A6C5-51E2-866C-76D5922E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOssScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOssScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListOssScanConfigResponseBody) GetData() []*ListOssScanConfigResponseBodyData {
	return s.Data
}

func (s *ListOssScanConfigResponseBody) GetPageInfo() *ListOssScanConfigResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListOssScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOssScanConfigResponseBody) SetData(v []*ListOssScanConfigResponseBodyData) *ListOssScanConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListOssScanConfigResponseBody) SetPageInfo(v *ListOssScanConfigResponseBodyPageInfo) *ListOssScanConfigResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListOssScanConfigResponseBody) SetRequestId(v string) *ListOssScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOssScanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOssScanConfigResponseBodyData struct {
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
	// 10
	BucketCount *int32 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
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
	// Indicates whether the policy is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the scan ends. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 06:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 443496
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The prefixes of the objects that are scanned.
	KeyPrefixList []*string `json:"KeyPrefixList,omitempty" xml:"KeyPrefixList,omitempty" type:"Repeated"`
	// The suffixes of the objects that are scanned.
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
	// 1698388233883
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Whether to enable real-time incremental detection. When this parameter is set to true, the parameters ScanDayList, StartTime, and EndTime are not effective.
	//
	// example:
	//
	// true
	RealTimeIncr *bool `json:"RealTimeIncr,omitempty" xml:"RealTimeIncr,omitempty"`
	// The days on which the scan is executed in a week.
	ScanDayList []*int32 `json:"ScanDayList,omitempty" xml:"ScanDayList,omitempty" type:"Repeated"`
	// The time when the scan starts. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListOssScanConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOssScanConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOssScanConfigResponseBodyData) GetAllKeyPrefix() *bool {
	return s.AllKeyPrefix
}

func (s *ListOssScanConfigResponseBodyData) GetBucketCount() *int32 {
	return s.BucketCount
}

func (s *ListOssScanConfigResponseBodyData) GetBucketNameList() []*string {
	return s.BucketNameList
}

func (s *ListOssScanConfigResponseBodyData) GetDecompressMaxFileCount() *int32 {
	return s.DecompressMaxFileCount
}

func (s *ListOssScanConfigResponseBodyData) GetDecompressMaxLayer() *int32 {
	return s.DecompressMaxLayer
}

func (s *ListOssScanConfigResponseBodyData) GetDecryptionList() []*string {
	return s.DecryptionList
}

func (s *ListOssScanConfigResponseBodyData) GetEnable() *int32 {
	return s.Enable
}

func (s *ListOssScanConfigResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListOssScanConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListOssScanConfigResponseBodyData) GetKeyPrefixList() []*string {
	return s.KeyPrefixList
}

func (s *ListOssScanConfigResponseBodyData) GetKeySuffixList() []*string {
	return s.KeySuffixList
}

func (s *ListOssScanConfigResponseBodyData) GetLastModifiedStartTime() *int64 {
	return s.LastModifiedStartTime
}

func (s *ListOssScanConfigResponseBodyData) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *ListOssScanConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListOssScanConfigResponseBodyData) GetRealTimeIncr() *bool {
	return s.RealTimeIncr
}

func (s *ListOssScanConfigResponseBodyData) GetScanDayList() []*int32 {
	return s.ScanDayList
}

func (s *ListOssScanConfigResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOssScanConfigResponseBodyData) SetAllKeyPrefix(v bool) *ListOssScanConfigResponseBodyData {
	s.AllKeyPrefix = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetBucketCount(v int32) *ListOssScanConfigResponseBodyData {
	s.BucketCount = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetBucketNameList(v []*string) *ListOssScanConfigResponseBodyData {
	s.BucketNameList = v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetDecompressMaxFileCount(v int32) *ListOssScanConfigResponseBodyData {
	s.DecompressMaxFileCount = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetDecompressMaxLayer(v int32) *ListOssScanConfigResponseBodyData {
	s.DecompressMaxLayer = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetDecryptionList(v []*string) *ListOssScanConfigResponseBodyData {
	s.DecryptionList = v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetEnable(v int32) *ListOssScanConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetEndTime(v string) *ListOssScanConfigResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetId(v int64) *ListOssScanConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetKeyPrefixList(v []*string) *ListOssScanConfigResponseBodyData {
	s.KeyPrefixList = v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetKeySuffixList(v []*string) *ListOssScanConfigResponseBodyData {
	s.KeySuffixList = v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetLastModifiedStartTime(v int64) *ListOssScanConfigResponseBodyData {
	s.LastModifiedStartTime = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetLastUpdateTime(v int64) *ListOssScanConfigResponseBodyData {
	s.LastUpdateTime = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetName(v string) *ListOssScanConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetRealTimeIncr(v bool) *ListOssScanConfigResponseBodyData {
	s.RealTimeIncr = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetScanDayList(v []*int32) *ListOssScanConfigResponseBodyData {
	s.ScanDayList = v
	return s
}

func (s *ListOssScanConfigResponseBodyData) SetStartTime(v string) *ListOssScanConfigResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListOssScanConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListOssScanConfigResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOssScanConfigResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOssScanConfigResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListOssScanConfigResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssScanConfigResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssScanConfigResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOssScanConfigResponseBodyPageInfo) SetCurrentPage(v int32) *ListOssScanConfigResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListOssScanConfigResponseBodyPageInfo) SetPageSize(v int32) *ListOssScanConfigResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListOssScanConfigResponseBodyPageInfo) SetTotalCount(v int32) *ListOssScanConfigResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListOssScanConfigResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
