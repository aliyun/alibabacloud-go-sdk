// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssBucketScanInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ListOssBucketScanInfoRequest
	GetBucketName() *string
	SetCurrentPage(v int32) *ListOssBucketScanInfoRequest
	GetCurrentPage() *int32
	SetFuzzBucketName(v string) *ListOssBucketScanInfoRequest
	GetFuzzBucketName() *string
	SetFuzzFileSystemName(v string) *ListOssBucketScanInfoRequest
	GetFuzzFileSystemName() *string
	SetHasRisk(v int32) *ListOssBucketScanInfoRequest
	GetHasRisk() *int32
	SetLang(v string) *ListOssBucketScanInfoRequest
	GetLang() *string
	SetPageSize(v int32) *ListOssBucketScanInfoRequest
	GetPageSize() *int32
	SetSource(v string) *ListOssBucketScanInfoRequest
	GetSource() *string
	SetStatus(v int32) *ListOssBucketScanInfoRequest
	GetStatus() *int32
}

type ListOssBucketScanInfoRequest struct {
	// The bucket name.
	//
	// example:
	//
	// iboxpublic****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The page number of the current page in a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The bucket name for fuzzy match.
	//
	// example:
	//
	// test
	FuzzBucketName *string `json:"FuzzBucketName,omitempty" xml:"FuzzBucketName,omitempty"`
	// The NAS file system name for fuzzy match.
	//
	// example:
	//
	// test
	FuzzFileSystemName *string `json:"FuzzFileSystemName,omitempty" xml:"FuzzFileSystemName,omitempty"`
	// Specifies whether risky files are detected. Valid values:
	//
	// - **0**: No risks detected.
	//
	// - **1**: Risky files exist.
	//
	// example:
	//
	// 1
	HasRisk *int32 `json:"HasRisk,omitempty" xml:"HasRisk,omitempty"`
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page in a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The business source. Valid values:
	//
	// - **OSS**: OSS
	//
	// - **NAS**: NAS
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The detection status. Valid values:
	//
	// - **1**: Not scanned.
	//
	// - **2**: Full scan in progress.
	//
	// - **3**: Incremental scan in progress.
	//
	// - **4**: Scanned.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOssBucketScanInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketScanInfoRequest) GoString() string {
	return s.String()
}

func (s *ListOssBucketScanInfoRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListOssBucketScanInfoRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssBucketScanInfoRequest) GetFuzzBucketName() *string {
	return s.FuzzBucketName
}

func (s *ListOssBucketScanInfoRequest) GetFuzzFileSystemName() *string {
	return s.FuzzFileSystemName
}

func (s *ListOssBucketScanInfoRequest) GetHasRisk() *int32 {
	return s.HasRisk
}

func (s *ListOssBucketScanInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *ListOssBucketScanInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssBucketScanInfoRequest) GetSource() *string {
	return s.Source
}

func (s *ListOssBucketScanInfoRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListOssBucketScanInfoRequest) SetBucketName(v string) *ListOssBucketScanInfoRequest {
	s.BucketName = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetCurrentPage(v int32) *ListOssBucketScanInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetFuzzBucketName(v string) *ListOssBucketScanInfoRequest {
	s.FuzzBucketName = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetFuzzFileSystemName(v string) *ListOssBucketScanInfoRequest {
	s.FuzzFileSystemName = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetHasRisk(v int32) *ListOssBucketScanInfoRequest {
	s.HasRisk = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetLang(v string) *ListOssBucketScanInfoRequest {
	s.Lang = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetPageSize(v int32) *ListOssBucketScanInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetSource(v string) *ListOssBucketScanInfoRequest {
	s.Source = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) SetStatus(v int32) *ListOssBucketScanInfoRequest {
	s.Status = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) Validate() error {
	return dara.Validate(s)
}
