// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnSslServerLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeVpnSslServerLogsResponseBody
	GetCount() *int32
	SetData(v *DescribeVpnSslServerLogsResponseBodyData) *DescribeVpnSslServerLogsResponseBody
	GetData() *DescribeVpnSslServerLogsResponseBodyData
	SetIsCompleted(v bool) *DescribeVpnSslServerLogsResponseBody
	GetIsCompleted() *bool
	SetPageNumber(v int32) *DescribeVpnSslServerLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnSslServerLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnSslServerLogsResponseBody
	GetRequestId() *string
}

type DescribeVpnSslServerLogsResponseBody struct {
	// The number of log entries.
	//
	// example:
	//
	// 10
	Count *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *DescribeVpnSslServerLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the log is accurate. Valid values:
	//
	// 	- **true**: accurate
	//
	// 	- **false**: inaccurate
	//
	// example:
	//
	// true
	IsCompleted *bool `json:"IsCompleted,omitempty" xml:"IsCompleted,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DEDAC5B1-9292-5BF7-BDDF-61BA58CFB2FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpnSslServerLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnSslServerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnSslServerLogsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVpnSslServerLogsResponseBody) GetData() *DescribeVpnSslServerLogsResponseBodyData {
	return s.Data
}

func (s *DescribeVpnSslServerLogsResponseBody) GetIsCompleted() *bool {
	return s.IsCompleted
}

func (s *DescribeVpnSslServerLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnSslServerLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnSslServerLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnSslServerLogsResponseBody) SetCount(v int32) *DescribeVpnSslServerLogsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBody) SetData(v *DescribeVpnSslServerLogsResponseBodyData) *DescribeVpnSslServerLogsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBody) SetIsCompleted(v bool) *DescribeVpnSslServerLogsResponseBody {
	s.IsCompleted = &v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBody) SetPageNumber(v int32) *DescribeVpnSslServerLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBody) SetPageSize(v int32) *DescribeVpnSslServerLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBody) SetRequestId(v string) *DescribeVpnSslServerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnSslServerLogsResponseBodyData struct {
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s DescribeVpnSslServerLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnSslServerLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVpnSslServerLogsResponseBodyData) GetLogs() []*string {
	return s.Logs
}

func (s *DescribeVpnSslServerLogsResponseBodyData) SetLogs(v []*string) *DescribeVpnSslServerLogsResponseBodyData {
	s.Logs = v
	return s
}

func (s *DescribeVpnSslServerLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
