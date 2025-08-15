// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectScanEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ListObjectScanEventRequest
	GetBucketName() *string
	SetCurrentPage(v int32) *ListObjectScanEventRequest
	GetCurrentPage() *int32
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
	SetTimeEnd(v int64) *ListObjectScanEventRequest
	GetTimeEnd() *int64
	SetTimeStart(v int64) *ListObjectScanEventRequest
	GetTimeStart() *int64
}

type ListObjectScanEventRequest struct {
	// example:
	//
	// testBucket******
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// e991e62f484bb6accd676e57a9******
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	ParentEventId *int64 `json:"ParentEventId,omitempty" xml:"ParentEventId,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1683862286000
	TimeEnd *int64 `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
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

func (s *ListObjectScanEventRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListObjectScanEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
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

func (s *ListObjectScanEventRequest) GetTimeEnd() *int64 {
	return s.TimeEnd
}

func (s *ListObjectScanEventRequest) GetTimeStart() *int64 {
	return s.TimeStart
}

func (s *ListObjectScanEventRequest) SetBucketName(v string) *ListObjectScanEventRequest {
	s.BucketName = &v
	return s
}

func (s *ListObjectScanEventRequest) SetCurrentPage(v int32) *ListObjectScanEventRequest {
	s.CurrentPage = &v
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
