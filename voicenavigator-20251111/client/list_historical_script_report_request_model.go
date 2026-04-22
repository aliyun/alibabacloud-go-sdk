// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalScriptReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListHistoricalScriptReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListHistoricalScriptReportRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListHistoricalScriptReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalScriptReportRequest
	GetPageSize() *int32
	SetScriptIds(v []*string) *ListHistoricalScriptReportRequest
	GetScriptIds() []*string
	SetStartTime(v int64) *ListHistoricalScriptReportRequest
	GetStartTime() *int64
}

type ListHistoricalScriptReportRequest struct {
	// example:
	//
	// 1582103299434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScriptIds []*string `json:"ScriptIds,omitempty" xml:"ScriptIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoricalScriptReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalScriptReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalScriptReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHistoricalScriptReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalScriptReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalScriptReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalScriptReportRequest) GetScriptIds() []*string {
	return s.ScriptIds
}

func (s *ListHistoricalScriptReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalScriptReportRequest) SetEndTime(v int64) *ListHistoricalScriptReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalScriptReportRequest) SetInstanceId(v string) *ListHistoricalScriptReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalScriptReportRequest) SetPageNumber(v int32) *ListHistoricalScriptReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalScriptReportRequest) SetPageSize(v int32) *ListHistoricalScriptReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalScriptReportRequest) SetScriptIds(v []*string) *ListHistoricalScriptReportRequest {
	s.ScriptIds = v
	return s
}

func (s *ListHistoricalScriptReportRequest) SetStartTime(v int64) *ListHistoricalScriptReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalScriptReportRequest) Validate() error {
	return dara.Validate(s)
}
