// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssBucketScanInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListOssBucketScanInfoResponseBodyData) *ListOssBucketScanInfoResponseBody
	GetData() []*ListOssBucketScanInfoResponseBodyData
	SetPageInfo(v *ListOssBucketScanInfoResponseBodyPageInfo) *ListOssBucketScanInfoResponseBody
	GetPageInfo() *ListOssBucketScanInfoResponseBodyPageInfo
	SetRequestId(v string) *ListOssBucketScanInfoResponseBody
	GetRequestId() *string
}

type ListOssBucketScanInfoResponseBody struct {
	// The data returned.
	Data []*ListOssBucketScanInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page information.
	PageInfo *ListOssBucketScanInfoResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6C578F36-92D2-552C-8AA0-86EB1AC2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOssBucketScanInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketScanInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListOssBucketScanInfoResponseBody) GetData() []*ListOssBucketScanInfoResponseBodyData {
	return s.Data
}

func (s *ListOssBucketScanInfoResponseBody) GetPageInfo() *ListOssBucketScanInfoResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListOssBucketScanInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOssBucketScanInfoResponseBody) SetData(v []*ListOssBucketScanInfoResponseBodyData) *ListOssBucketScanInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListOssBucketScanInfoResponseBody) SetPageInfo(v *ListOssBucketScanInfoResponseBodyPageInfo) *ListOssBucketScanInfoResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListOssBucketScanInfoResponseBody) SetRequestId(v string) *ListOssBucketScanInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOssBucketScanInfoResponseBodyData struct {
	// The name of the bucket.
	//
	// example:
	//
	// hz-new02****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// Configuration status, valid values:
	//
	// - **0**: No Configuration.
	//
	// - **1**: Not Open.
	//
	// - **2**: Open.
	//
	// example:
	//
	// 1
	ConfigStatus *int32 `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	// Bucket decompression configuration status, valid values:
	//
	// - **0**: Decompression not configured.
	//
	// - **1**: Decompression configured.
	//
	// example:
	//
	// 0
	DecompressStatus *int32 `json:"DecompressStatus,omitempty" xml:"DecompressStatus,omitempty"`
	// The number of high-risk objects.
	//
	// example:
	//
	// 0
	HighRisk *int64 `json:"HighRisk,omitempty" xml:"HighRisk,omitempty"`
	// The time when the most recent check ended. Unit: milliseconds.
	//
	// example:
	//
	// 1698388233883
	LastScanEndTime *int64 `json:"LastScanEndTime,omitempty" xml:"LastScanEndTime,omitempty"`
	// The time when the bucket was last checked. Unit: milliseconds.
	//
	// example:
	//
	// 1698388233883
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The number of low-risk objects.
	//
	// example:
	//
	// 0
	LowRisk *int64 `json:"LowRisk,omitempty" xml:"LowRisk,omitempty"`
	// The number of medium-risk objects.
	//
	// example:
	//
	// 0
	MediumRisk *int64 `json:"MediumRisk,omitempty" xml:"MediumRisk,omitempty"`
	// The reason why the bucket cannot be checked.
	//
	// example:
	//
	// Unsupported Region.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of objects that are checked.
	//
	// example:
	//
	// 100
	ScanObject *int64 `json:"ScanObject,omitempty" xml:"ScanObject,omitempty"`
	// Indicates whether the bucket is checked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Scanned *bool `json:"Scanned,omitempty" xml:"Scanned,omitempty"`
	// The check status of the bucket. Valid values:
	//
	// 	- **1**: The bucket is not checked.
	//
	// 	- **2**: All objects in the bucket are being checked.
	//
	// 	- **3**: Only new objects are being checked.
	//
	// 	- **4**: The bucket is checked.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the bucket. Valid values:
	//
	// 	- **Standard**
	//
	// 	- **IA**
	//
	// 	- **Archive**
	//
	// 	- **ColdArchive**
	//
	// example:
	//
	// Archive
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// Indicates whether the bucket can be checked. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
	// The total number of objects in the bucket.
	//
	// example:
	//
	// 100
	TotalObject *int64 `json:"TotalObject,omitempty" xml:"TotalObject,omitempty"`
}

func (s ListOssBucketScanInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketScanInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOssBucketScanInfoResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *ListOssBucketScanInfoResponseBodyData) GetConfigStatus() *int32 {
	return s.ConfigStatus
}

func (s *ListOssBucketScanInfoResponseBodyData) GetDecompressStatus() *int32 {
	return s.DecompressStatus
}

func (s *ListOssBucketScanInfoResponseBodyData) GetHighRisk() *int64 {
	return s.HighRisk
}

func (s *ListOssBucketScanInfoResponseBodyData) GetLastScanEndTime() *int64 {
	return s.LastScanEndTime
}

func (s *ListOssBucketScanInfoResponseBodyData) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *ListOssBucketScanInfoResponseBodyData) GetLowRisk() *int64 {
	return s.LowRisk
}

func (s *ListOssBucketScanInfoResponseBodyData) GetMediumRisk() *int64 {
	return s.MediumRisk
}

func (s *ListOssBucketScanInfoResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListOssBucketScanInfoResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOssBucketScanInfoResponseBodyData) GetScanObject() *int64 {
	return s.ScanObject
}

func (s *ListOssBucketScanInfoResponseBodyData) GetScanned() *bool {
	return s.Scanned
}

func (s *ListOssBucketScanInfoResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListOssBucketScanInfoResponseBodyData) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ListOssBucketScanInfoResponseBodyData) GetSupport() *bool {
	return s.Support
}

func (s *ListOssBucketScanInfoResponseBodyData) GetTotalObject() *int64 {
	return s.TotalObject
}

func (s *ListOssBucketScanInfoResponseBodyData) SetBucketName(v string) *ListOssBucketScanInfoResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetConfigStatus(v int32) *ListOssBucketScanInfoResponseBodyData {
	s.ConfigStatus = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetDecompressStatus(v int32) *ListOssBucketScanInfoResponseBodyData {
	s.DecompressStatus = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetHighRisk(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.HighRisk = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetLastScanEndTime(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.LastScanEndTime = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetLastScanTime(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.LastScanTime = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetLowRisk(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.LowRisk = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetMediumRisk(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.MediumRisk = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetMessage(v string) *ListOssBucketScanInfoResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetRegionId(v string) *ListOssBucketScanInfoResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetScanObject(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.ScanObject = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetScanned(v bool) *ListOssBucketScanInfoResponseBodyData {
	s.Scanned = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetStatus(v int32) *ListOssBucketScanInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetStorageClass(v string) *ListOssBucketScanInfoResponseBodyData {
	s.StorageClass = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetSupport(v bool) *ListOssBucketScanInfoResponseBodyData {
	s.Support = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) SetTotalObject(v int64) *ListOssBucketScanInfoResponseBodyData {
	s.TotalObject = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListOssBucketScanInfoResponseBodyPageInfo struct {
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
	// 165
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOssBucketScanInfoResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketScanInfoResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) SetCurrentPage(v int32) *ListOssBucketScanInfoResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) SetPageSize(v int32) *ListOssBucketScanInfoResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) SetTotalCount(v int32) *ListOssBucketScanInfoResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListOssBucketScanInfoResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
