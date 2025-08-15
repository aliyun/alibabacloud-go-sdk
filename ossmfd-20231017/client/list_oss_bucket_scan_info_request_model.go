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
	SetHasRisk(v int32) *ListOssBucketScanInfoRequest
	GetHasRisk() *int32
	SetLang(v string) *ListOssBucketScanInfoRequest
	GetLang() *string
	SetPageSize(v int32) *ListOssBucketScanInfoRequest
	GetPageSize() *int32
	SetStatus(v int32) *ListOssBucketScanInfoRequest
	GetStatus() *int32
}

type ListOssBucketScanInfoRequest struct {
	// example:
	//
	// testBucket******
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// test******
	FuzzBucketName *string `json:"FuzzBucketName,omitempty" xml:"FuzzBucketName,omitempty"`
	// example:
	//
	// 1
	HasRisk *int32 `json:"HasRisk,omitempty" xml:"HasRisk,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListOssBucketScanInfoRequest) GetHasRisk() *int32 {
	return s.HasRisk
}

func (s *ListOssBucketScanInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *ListOssBucketScanInfoRequest) GetPageSize() *int32 {
	return s.PageSize
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

func (s *ListOssBucketScanInfoRequest) SetStatus(v int32) *ListOssBucketScanInfoRequest {
	s.Status = &v
	return s
}

func (s *ListOssBucketScanInfoRequest) Validate() error {
	return dara.Validate(s)
}
