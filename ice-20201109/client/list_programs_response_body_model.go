// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProgramsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListProgramsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListProgramsResponseBody
	GetPageSize() *int32
	SetPrograms(v []*ChannelAssemblyProgram) *ListProgramsResponseBody
	GetPrograms() []*ChannelAssemblyProgram
	SetRequestId(v string) *ListProgramsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProgramsResponseBody
	GetTotalCount() *int32
}

type ListProgramsResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The programs.
	Programs []*ChannelAssemblyProgram `json:"Programs,omitempty" xml:"Programs,omitempty" type:"Repeated"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of programs returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProgramsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProgramsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProgramsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListProgramsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProgramsResponseBody) GetPrograms() []*ChannelAssemblyProgram {
	return s.Programs
}

func (s *ListProgramsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProgramsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProgramsResponseBody) SetPageNo(v int32) *ListProgramsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListProgramsResponseBody) SetPageSize(v int32) *ListProgramsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProgramsResponseBody) SetPrograms(v []*ChannelAssemblyProgram) *ListProgramsResponseBody {
	s.Programs = v
	return s
}

func (s *ListProgramsResponseBody) SetRequestId(v string) *ListProgramsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProgramsResponseBody) SetTotalCount(v int32) *ListProgramsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProgramsResponseBody) Validate() error {
	return dara.Validate(s)
}
