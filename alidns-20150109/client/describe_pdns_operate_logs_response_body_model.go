// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsOperateLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*DescribePdnsOperateLogsResponseBodyLogs) *DescribePdnsOperateLogsResponseBody
	GetLogs() []*DescribePdnsOperateLogsResponseBodyLogs
	SetPageNumber(v int64) *DescribePdnsOperateLogsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsOperateLogsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePdnsOperateLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePdnsOperateLogsResponseBody
	GetTotalCount() *int64
}

type DescribePdnsOperateLogsResponseBody struct {
	Logs       []*DescribePdnsOperateLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePdnsOperateLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsOperateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsOperateLogsResponseBody) GetLogs() []*DescribePdnsOperateLogsResponseBodyLogs {
	return s.Logs
}

func (s *DescribePdnsOperateLogsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsOperateLogsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsOperateLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsOperateLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsOperateLogsResponseBody) SetLogs(v []*DescribePdnsOperateLogsResponseBodyLogs) *DescribePdnsOperateLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribePdnsOperateLogsResponseBody) SetPageNumber(v int64) *DescribePdnsOperateLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBody) SetPageSize(v int64) *DescribePdnsOperateLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBody) SetRequestId(v string) *DescribePdnsOperateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBody) SetTotalCount(v int64) *DescribePdnsOperateLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePdnsOperateLogsResponseBodyLogs struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	OperateTime *string `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s DescribePdnsOperateLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsOperateLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) GetAction() *string {
	return s.Action
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) GetOperateTime() *string {
	return s.OperateTime
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) GetType() *string {
	return s.Type
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) GetContent() *string {
	return s.Content
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) SetAction(v string) *DescribePdnsOperateLogsResponseBodyLogs {
	s.Action = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) SetOperateTime(v string) *DescribePdnsOperateLogsResponseBodyLogs {
	s.OperateTime = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) SetType(v string) *DescribePdnsOperateLogsResponseBodyLogs {
	s.Type = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) SetContent(v string) *DescribePdnsOperateLogsResponseBodyLogs {
	s.Content = &v
	return s
}

func (s *DescribePdnsOperateLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
