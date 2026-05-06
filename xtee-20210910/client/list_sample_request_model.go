// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListSampleRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListSampleRequest
	GetLang() *string
	SetPageSize(v int32) *ListSampleRequest
	GetPageSize() *int32
	SetRegId(v string) *ListSampleRequest
	GetRegId() *string
	SetSampleName(v string) *ListSampleRequest
	GetSampleName() *string
	SetTab(v string) *ListSampleRequest
	GetTab() *string
	SetType(v string) *ListSampleRequest
	GetType() *string
	SetUploadTimeEnd(v string) *ListSampleRequest
	GetUploadTimeEnd() *string
	SetUploadTimeStart(v string) *ListSampleRequest
	GetUploadTimeStart() *string
}

type ListSampleRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// SampleTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2023-03-09 14:45:26
	UploadTimeEnd *string `json:"UploadTimeEnd,omitempty" xml:"UploadTimeEnd,omitempty"`
	// example:
	//
	// 2023-03-09 14:45:23
	UploadTimeStart *string `json:"UploadTimeStart,omitempty" xml:"UploadTimeStart,omitempty"`
}

func (s ListSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSampleRequest) GoString() string {
	return s.String()
}

func (s *ListSampleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSampleRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSampleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSampleRequest) GetRegId() *string {
	return s.RegId
}

func (s *ListSampleRequest) GetSampleName() *string {
	return s.SampleName
}

func (s *ListSampleRequest) GetTab() *string {
	return s.Tab
}

func (s *ListSampleRequest) GetType() *string {
	return s.Type
}

func (s *ListSampleRequest) GetUploadTimeEnd() *string {
	return s.UploadTimeEnd
}

func (s *ListSampleRequest) GetUploadTimeStart() *string {
	return s.UploadTimeStart
}

func (s *ListSampleRequest) SetCurrentPage(v int32) *ListSampleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSampleRequest) SetLang(v string) *ListSampleRequest {
	s.Lang = &v
	return s
}

func (s *ListSampleRequest) SetPageSize(v int32) *ListSampleRequest {
	s.PageSize = &v
	return s
}

func (s *ListSampleRequest) SetRegId(v string) *ListSampleRequest {
	s.RegId = &v
	return s
}

func (s *ListSampleRequest) SetSampleName(v string) *ListSampleRequest {
	s.SampleName = &v
	return s
}

func (s *ListSampleRequest) SetTab(v string) *ListSampleRequest {
	s.Tab = &v
	return s
}

func (s *ListSampleRequest) SetType(v string) *ListSampleRequest {
	s.Type = &v
	return s
}

func (s *ListSampleRequest) SetUploadTimeEnd(v string) *ListSampleRequest {
	s.UploadTimeEnd = &v
	return s
}

func (s *ListSampleRequest) SetUploadTimeStart(v string) *ListSampleRequest {
	s.UploadTimeStart = &v
	return s
}

func (s *ListSampleRequest) Validate() error {
	return dara.Validate(s)
}
