// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalScriptReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListHistoricalScriptReportShrinkRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListHistoricalScriptReportShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListHistoricalScriptReportShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalScriptReportShrinkRequest
	GetPageSize() *int32
	SetScriptIdsShrink(v string) *ListHistoricalScriptReportShrinkRequest
	GetScriptIdsShrink() *string
	SetStartTime(v int64) *ListHistoricalScriptReportShrinkRequest
	GetStartTime() *int64
}

type ListHistoricalScriptReportShrinkRequest struct {
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
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScriptIdsShrink *string `json:"ScriptIds,omitempty" xml:"ScriptIds,omitempty"`
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoricalScriptReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalScriptReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalScriptReportShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHistoricalScriptReportShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalScriptReportShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalScriptReportShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalScriptReportShrinkRequest) GetScriptIdsShrink() *string {
	return s.ScriptIdsShrink
}

func (s *ListHistoricalScriptReportShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalScriptReportShrinkRequest) SetEndTime(v int64) *ListHistoricalScriptReportShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalScriptReportShrinkRequest) SetInstanceId(v string) *ListHistoricalScriptReportShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalScriptReportShrinkRequest) SetPageNumber(v int32) *ListHistoricalScriptReportShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalScriptReportShrinkRequest) SetPageSize(v int32) *ListHistoricalScriptReportShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalScriptReportShrinkRequest) SetScriptIdsShrink(v string) *ListHistoricalScriptReportShrinkRequest {
	s.ScriptIdsShrink = &v
	return s
}

func (s *ListHistoricalScriptReportShrinkRequest) SetStartTime(v int64) *ListHistoricalScriptReportShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalScriptReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
