// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListSchedulesResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListSchedulesResponseBody
	GetPageSize() *int32
	SetPrograms(v []*ChannelAssemblyScheduleData) *ListSchedulesResponseBody
	GetPrograms() []*ChannelAssemblyScheduleData
	SetRequestId(v string) *ListSchedulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSchedulesResponseBody
	GetTotalCount() *int32
}

type ListSchedulesResponseBody struct {
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
	// The program schedule.
	Programs []*ChannelAssemblyScheduleData `json:"Programs,omitempty" xml:"Programs,omitempty" type:"Repeated"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSchedulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchedulesResponseBody) GetPrograms() []*ChannelAssemblyScheduleData {
	return s.Programs
}

func (s *ListSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchedulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSchedulesResponseBody) SetPageNo(v int32) *ListSchedulesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListSchedulesResponseBody) SetPageSize(v int32) *ListSchedulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSchedulesResponseBody) SetPrograms(v []*ChannelAssemblyScheduleData) *ListSchedulesResponseBody {
	s.Programs = v
	return s
}

func (s *ListSchedulesResponseBody) SetRequestId(v string) *ListSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchedulesResponseBody) SetTotalCount(v int32) *ListSchedulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}
