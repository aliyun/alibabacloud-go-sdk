// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTaskGroupRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListTaskGroupRequest
	GetLang() *string
	SetPageSize(v int32) *ListTaskGroupRequest
	GetPageSize() *int32
	SetRegId(v string) *ListTaskGroupRequest
	GetRegId() *string
	SetSampleName(v string) *ListTaskGroupRequest
	GetSampleName() *string
	SetTaskGroupName(v string) *ListTaskGroupRequest
	GetTaskGroupName() *string
	SetType(v string) *ListTaskGroupRequest
	GetType() *string
}

type ListTaskGroupRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the error message returned by the API. Valid values: zh: Chinese; en: English. The default value is en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The area encoding.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Sample name.
	//
	// example:
	//
	// TEst
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// Task group name.
	//
	// example:
	//
	// TeskGroupTest
	TaskGroupName *string `json:"TaskGroupName,omitempty" xml:"TaskGroupName,omitempty"`
	// Access type.
	//
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTaskGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *ListTaskGroupRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTaskGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTaskGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskGroupRequest) GetRegId() *string {
	return s.RegId
}

func (s *ListTaskGroupRequest) GetSampleName() *string {
	return s.SampleName
}

func (s *ListTaskGroupRequest) GetTaskGroupName() *string {
	return s.TaskGroupName
}

func (s *ListTaskGroupRequest) GetType() *string {
	return s.Type
}

func (s *ListTaskGroupRequest) SetCurrentPage(v int32) *ListTaskGroupRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTaskGroupRequest) SetLang(v string) *ListTaskGroupRequest {
	s.Lang = &v
	return s
}

func (s *ListTaskGroupRequest) SetPageSize(v int32) *ListTaskGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskGroupRequest) SetRegId(v string) *ListTaskGroupRequest {
	s.RegId = &v
	return s
}

func (s *ListTaskGroupRequest) SetSampleName(v string) *ListTaskGroupRequest {
	s.SampleName = &v
	return s
}

func (s *ListTaskGroupRequest) SetTaskGroupName(v string) *ListTaskGroupRequest {
	s.TaskGroupName = &v
	return s
}

func (s *ListTaskGroupRequest) SetType(v string) *ListTaskGroupRequest {
	s.Type = &v
	return s
}

func (s *ListTaskGroupRequest) Validate() error {
	return dara.Validate(s)
}
