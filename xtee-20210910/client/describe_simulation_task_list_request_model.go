// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSimulationTaskListRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeSimulationTaskListRequest
	GetCurrentPage() *string
	SetId(v string) *DescribeSimulationTaskListRequest
	GetId() *string
	SetName(v string) *DescribeSimulationTaskListRequest
	GetName() *string
	SetPageSize(v string) *DescribeSimulationTaskListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeSimulationTaskListRequest
	GetRegId() *string
	SetTitle(v string) *DescribeSimulationTaskListRequest
	GetTitle() *string
}

type DescribeSimulationTaskListRequest struct {
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
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 3144
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name
	//
	// example:
	//
	// SIMULATION
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Title.
	//
	// example:
	//
	// 仿真任务
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeSimulationTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimulationTaskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSimulationTaskListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSimulationTaskListRequest) GetId() *string {
	return s.Id
}

func (s *DescribeSimulationTaskListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSimulationTaskListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSimulationTaskListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSimulationTaskListRequest) GetTitle() *string {
	return s.Title
}

func (s *DescribeSimulationTaskListRequest) SetLang(v string) *DescribeSimulationTaskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) SetCurrentPage(v string) *DescribeSimulationTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) SetId(v string) *DescribeSimulationTaskListRequest {
	s.Id = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) SetName(v string) *DescribeSimulationTaskListRequest {
	s.Name = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) SetPageSize(v string) *DescribeSimulationTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) SetRegId(v string) *DescribeSimulationTaskListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) SetTitle(v string) *DescribeSimulationTaskListRequest {
	s.Title = &v
	return s
}

func (s *DescribeSimulationTaskListRequest) Validate() error {
	return dara.Validate(s)
}
