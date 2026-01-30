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
	Data     []*ListOssBucketScanInfoResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListOssBucketScanInfoResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2220FE66-76EF-5D9D-A0EE-3221CC******
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOssBucketScanInfoResponseBodyData struct {
	// example:
	//
	// testBucket******
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 1
	ConfigStatus *int32 `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	// example:
	//
	// 1
	DecompressStatus *int32 `json:"DecompressStatus,omitempty" xml:"DecompressStatus,omitempty"`
	// example:
	//
	// 0
	HighRisk *int64 `json:"HighRisk,omitempty" xml:"HighRisk,omitempty"`
	// example:
	//
	// 1698388233884
	LastScanEndTime *int64 `json:"LastScanEndTime,omitempty" xml:"LastScanEndTime,omitempty"`
	// example:
	//
	// 1698388233883
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// example:
	//
	// 1
	LowRisk *int64 `json:"LowRisk,omitempty" xml:"LowRisk,omitempty"`
	// example:
	//
	// 0
	MediumRisk *int64 `json:"MediumRisk,omitempty" xml:"MediumRisk,omitempty"`
	// example:
	//
	// Unsupported Region。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	ScanObject *int64 `json:"ScanObject,omitempty" xml:"ScanObject,omitempty"`
	// example:
	//
	// true
	Scanned *bool `json:"Scanned,omitempty" xml:"Scanned,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// IA
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
	// example:
	//
	// 2
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
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 55
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
