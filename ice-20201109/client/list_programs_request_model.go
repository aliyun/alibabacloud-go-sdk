// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProgramsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *ListProgramsRequest
	GetChannelName() *string
	SetPageNo(v string) *ListProgramsRequest
	GetPageNo() *string
	SetPageSize(v string) *ListProgramsRequest
	GetPageSize() *string
	SetProgramName(v string) *ListProgramsRequest
	GetProgramName() *string
	SetSortBy(v string) *ListProgramsRequest
	GetSortBy() *string
}

type ListProgramsRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the program.
	//
	// example:
	//
	// program1
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order. Valid values:
	//
	// 	- asc: ascending order.
	//
	// 	- desc: descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListProgramsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProgramsRequest) GoString() string {
	return s.String()
}

func (s *ListProgramsRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListProgramsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListProgramsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListProgramsRequest) GetProgramName() *string {
	return s.ProgramName
}

func (s *ListProgramsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProgramsRequest) SetChannelName(v string) *ListProgramsRequest {
	s.ChannelName = &v
	return s
}

func (s *ListProgramsRequest) SetPageNo(v string) *ListProgramsRequest {
	s.PageNo = &v
	return s
}

func (s *ListProgramsRequest) SetPageSize(v string) *ListProgramsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProgramsRequest) SetProgramName(v string) *ListProgramsRequest {
	s.ProgramName = &v
	return s
}

func (s *ListProgramsRequest) SetSortBy(v string) *ListProgramsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProgramsRequest) Validate() error {
	return dara.Validate(s)
}
