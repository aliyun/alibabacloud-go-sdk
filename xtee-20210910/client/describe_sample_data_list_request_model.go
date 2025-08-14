// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleDataListRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeSampleDataListRequest
	GetCurrentPage() *string
	SetDeleteTag(v string) *DescribeSampleDataListRequest
	GetDeleteTag() *string
	SetPageSize(v string) *DescribeSampleDataListRequest
	GetPageSize() *string
	SetQueryContent(v string) *DescribeSampleDataListRequest
	GetQueryContent() *string
	SetRegId(v string) *DescribeSampleDataListRequest
	GetRegId() *string
	SetSampleId(v int64) *DescribeSampleDataListRequest
	GetSampleId() *int64
	SetScene(v string) *DescribeSampleDataListRequest
	GetScene() *string
	SetStatus(v string) *DescribeSampleDataListRequest
	GetStatus() *string
}

type DescribeSampleDataListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Deletion status
	//
	// example:
	//
	// DELETE
	DeleteTag *string `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Query content
	//
	// example:
	//
	// 手机号
	QueryContent *string `json:"queryContent,omitempty" xml:"queryContent,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Sample ID.
	//
	// example:
	//
	// 5467
	SampleId *int64 `json:"sampleId,omitempty" xml:"sampleId,omitempty"`
	// Scene
	//
	// example:
	//
	// 1
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// Status.
	//
	// example:
	//
	// CREATE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeSampleDataListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleDataListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSampleDataListRequest) GetDeleteTag() *string {
	return s.DeleteTag
}

func (s *DescribeSampleDataListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSampleDataListRequest) GetQueryContent() *string {
	return s.QueryContent
}

func (s *DescribeSampleDataListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleDataListRequest) GetSampleId() *int64 {
	return s.SampleId
}

func (s *DescribeSampleDataListRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeSampleDataListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSampleDataListRequest) SetLang(v string) *DescribeSampleDataListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetCurrentPage(v string) *DescribeSampleDataListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetDeleteTag(v string) *DescribeSampleDataListRequest {
	s.DeleteTag = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetPageSize(v string) *DescribeSampleDataListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetQueryContent(v string) *DescribeSampleDataListRequest {
	s.QueryContent = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetRegId(v string) *DescribeSampleDataListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetSampleId(v int64) *DescribeSampleDataListRequest {
	s.SampleId = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetScene(v string) *DescribeSampleDataListRequest {
	s.Scene = &v
	return s
}

func (s *DescribeSampleDataListRequest) SetStatus(v string) *DescribeSampleDataListRequest {
	s.Status = &v
	return s
}

func (s *DescribeSampleDataListRequest) Validate() error {
	return dara.Validate(s)
}
