// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectScanEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchType(v string) *ListObjectScanEventRequest
	GetBatchType() *string
	SetBucketName(v string) *ListObjectScanEventRequest
	GetBucketName() *string
	SetCurrentPage(v int32) *ListObjectScanEventRequest
	GetCurrentPage() *int32
	SetEventId(v int64) *ListObjectScanEventRequest
	GetEventId() *int64
	SetEventName(v string) *ListObjectScanEventRequest
	GetEventName() *string
	SetLang(v string) *ListObjectScanEventRequest
	GetLang() *string
	SetMd5(v string) *ListObjectScanEventRequest
	GetMd5() *string
	SetOssKey(v string) *ListObjectScanEventRequest
	GetOssKey() *string
	SetPageSize(v int32) *ListObjectScanEventRequest
	GetPageSize() *int32
	SetParentEventId(v int64) *ListObjectScanEventRequest
	GetParentEventId() *int64
	SetRiskLevel(v string) *ListObjectScanEventRequest
	GetRiskLevel() *string
	SetSource(v string) *ListObjectScanEventRequest
	GetSource() *string
	SetStatus(v int32) *ListObjectScanEventRequest
	GetStatus() *int32
	SetTimeEnd(v int64) *ListObjectScanEventRequest
	GetTimeEnd() *int64
	SetTimeStart(v int64) *ListObjectScanEventRequest
	GetTimeStart() *int64
}

type ListObjectScanEventRequest struct {
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// ltrbuck****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EventId     *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// WebShell
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The MD5 hash value of the file.
	//
	// example:
	//
	// 0552c44e243abdea1729d4507bce****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The key of the file that is stored in an OSS bucket.
	//
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the alert that is generated for the package to which the subfile belongs.
	//
	// example:
	//
	// 1
	ParentEventId *int64 `json:"ParentEventId,omitempty" xml:"ParentEventId,omitempty"`
	// The risk level of the alert. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The method that is used to detect the malicious file. Valid values:
	//
	// 	- **API**: uses API operations.
	//
	// 	- **OSS**: uses Object Storage Service (OSS) file check.
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end of the time range during which the exception is detected.
	//
	// example:
	//
	// 1683862286000
	TimeEnd *int64 `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// The beginning of the time range during which the exception is detected.
	//
	// example:
	//
	// 1683603086000
	TimeStart *int64 `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
}

func (s ListObjectScanEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListObjectScanEventRequest) GoString() string {
	return s.String()
}

func (s *ListObjectScanEventRequest) GetBatchType() *string {
	return s.BatchType
}

func (s *ListObjectScanEventRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListObjectScanEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListObjectScanEventRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *ListObjectScanEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *ListObjectScanEventRequest) GetLang() *string {
	return s.Lang
}

func (s *ListObjectScanEventRequest) GetMd5() *string {
	return s.Md5
}

func (s *ListObjectScanEventRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *ListObjectScanEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListObjectScanEventRequest) GetParentEventId() *int64 {
	return s.ParentEventId
}

func (s *ListObjectScanEventRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListObjectScanEventRequest) GetSource() *string {
	return s.Source
}

func (s *ListObjectScanEventRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListObjectScanEventRequest) GetTimeEnd() *int64 {
	return s.TimeEnd
}

func (s *ListObjectScanEventRequest) GetTimeStart() *int64 {
	return s.TimeStart
}

func (s *ListObjectScanEventRequest) SetBatchType(v string) *ListObjectScanEventRequest {
	s.BatchType = &v
	return s
}

func (s *ListObjectScanEventRequest) SetBucketName(v string) *ListObjectScanEventRequest {
	s.BucketName = &v
	return s
}

func (s *ListObjectScanEventRequest) SetCurrentPage(v int32) *ListObjectScanEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListObjectScanEventRequest) SetEventId(v int64) *ListObjectScanEventRequest {
	s.EventId = &v
	return s
}

func (s *ListObjectScanEventRequest) SetEventName(v string) *ListObjectScanEventRequest {
	s.EventName = &v
	return s
}

func (s *ListObjectScanEventRequest) SetLang(v string) *ListObjectScanEventRequest {
	s.Lang = &v
	return s
}

func (s *ListObjectScanEventRequest) SetMd5(v string) *ListObjectScanEventRequest {
	s.Md5 = &v
	return s
}

func (s *ListObjectScanEventRequest) SetOssKey(v string) *ListObjectScanEventRequest {
	s.OssKey = &v
	return s
}

func (s *ListObjectScanEventRequest) SetPageSize(v int32) *ListObjectScanEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListObjectScanEventRequest) SetParentEventId(v int64) *ListObjectScanEventRequest {
	s.ParentEventId = &v
	return s
}

func (s *ListObjectScanEventRequest) SetRiskLevel(v string) *ListObjectScanEventRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListObjectScanEventRequest) SetSource(v string) *ListObjectScanEventRequest {
	s.Source = &v
	return s
}

func (s *ListObjectScanEventRequest) SetStatus(v int32) *ListObjectScanEventRequest {
	s.Status = &v
	return s
}

func (s *ListObjectScanEventRequest) SetTimeEnd(v int64) *ListObjectScanEventRequest {
	s.TimeEnd = &v
	return s
}

func (s *ListObjectScanEventRequest) SetTimeStart(v int64) *ListObjectScanEventRequest {
	s.TimeStart = &v
	return s
}

func (s *ListObjectScanEventRequest) Validate() error {
	return dara.Validate(s)
}
