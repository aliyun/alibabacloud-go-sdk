// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeVpnConnectionLogsResponseBody
	GetCount() *int32
	SetData(v *DescribeVpnConnectionLogsResponseBodyData) *DescribeVpnConnectionLogsResponseBody
	GetData() *DescribeVpnConnectionLogsResponseBodyData
	SetIsCompleted(v bool) *DescribeVpnConnectionLogsResponseBody
	GetIsCompleted() *bool
	SetPageNumber(v int32) *DescribeVpnConnectionLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnConnectionLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnConnectionLogsResponseBody
	GetRequestId() *string
}

type DescribeVpnConnectionLogsResponseBody struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The log list.
	Data *DescribeVpnConnectionLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the log is accurate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsCompleted *bool `json:"IsCompleted,omitempty" xml:"IsCompleted,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CF161502-4959-5C3B-B499-09B87BA931D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpnConnectionLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionLogsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVpnConnectionLogsResponseBody) GetData() *DescribeVpnConnectionLogsResponseBodyData {
	return s.Data
}

func (s *DescribeVpnConnectionLogsResponseBody) GetIsCompleted() *bool {
	return s.IsCompleted
}

func (s *DescribeVpnConnectionLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnConnectionLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnConnectionLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnConnectionLogsResponseBody) SetCount(v int32) *DescribeVpnConnectionLogsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBody) SetData(v *DescribeVpnConnectionLogsResponseBodyData) *DescribeVpnConnectionLogsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBody) SetIsCompleted(v bool) *DescribeVpnConnectionLogsResponseBody {
	s.IsCompleted = &v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBody) SetPageNumber(v int32) *DescribeVpnConnectionLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBody) SetPageSize(v int32) *DescribeVpnConnectionLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBody) SetRequestId(v string) *DescribeVpnConnectionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionLogsResponseBodyData struct {
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s DescribeVpnConnectionLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionLogsResponseBodyData) GetLogs() []*string {
	return s.Logs
}

func (s *DescribeVpnConnectionLogsResponseBodyData) SetLogs(v []*string) *DescribeVpnConnectionLogsResponseBodyData {
	s.Logs = v
	return s
}

func (s *DescribeVpnConnectionLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
