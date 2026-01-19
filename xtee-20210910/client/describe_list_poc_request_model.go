// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListPocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *DescribeListPocRequest
	GetCurrentPage() *int64
	SetLang(v string) *DescribeListPocRequest
	GetLang() *string
	SetPageSize(v int64) *DescribeListPocRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeListPocRequest
	GetRegId() *string
	SetTaskName(v string) *DescribeListPocRequest
	GetTaskName() *string
	SetTaskStatus(v string) *DescribeListPocRequest
	GetTaskStatus() *string
	SetType(v string) *DescribeListPocRequest
	GetType() *string
}

type DescribeListPocRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Page size.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId      *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Type
	//
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeListPocRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListPocRequest) GoString() string {
	return s.String()
}

func (s *DescribeListPocRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeListPocRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeListPocRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeListPocRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeListPocRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeListPocRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeListPocRequest) GetType() *string {
	return s.Type
}

func (s *DescribeListPocRequest) SetCurrentPage(v int64) *DescribeListPocRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeListPocRequest) SetLang(v string) *DescribeListPocRequest {
	s.Lang = &v
	return s
}

func (s *DescribeListPocRequest) SetPageSize(v int64) *DescribeListPocRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeListPocRequest) SetRegId(v string) *DescribeListPocRequest {
	s.RegId = &v
	return s
}

func (s *DescribeListPocRequest) SetTaskName(v string) *DescribeListPocRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeListPocRequest) SetTaskStatus(v string) *DescribeListPocRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeListPocRequest) SetType(v string) *DescribeListPocRequest {
	s.Type = &v
	return s
}

func (s *DescribeListPocRequest) Validate() error {
	return dara.Validate(s)
}
