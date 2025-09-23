// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructMvRecommendSubTaskModel interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *OpenStructMvRecommendSubTaskModel
	GetEndTime() *string
	SetMinRewriteQueryCount(v int32) *OpenStructMvRecommendSubTaskModel
	GetMinRewriteQueryCount() *int32
	SetMinRewriteQueryPattern(v int32) *OpenStructMvRecommendSubTaskModel
	GetMinRewriteQueryPattern() *int32
	SetScanQueriesCount(v int64) *OpenStructMvRecommendSubTaskModel
	GetScanQueriesCount() *int64
	SetSlowQueryThreshold(v int32) *OpenStructMvRecommendSubTaskModel
	GetSlowQueryThreshold() *int32
	SetStartTime(v string) *OpenStructMvRecommendSubTaskModel
	GetStartTime() *string
	SetStatus(v string) *OpenStructMvRecommendSubTaskModel
	GetStatus() *string
	SetSubQueriesCount(v int64) *OpenStructMvRecommendSubTaskModel
	GetSubQueriesCount() *int64
	SetSubtaskId(v int64) *OpenStructMvRecommendSubTaskModel
	GetSubtaskId() *int64
}

type OpenStructMvRecommendSubTaskModel struct {
	EndTime                *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MinRewriteQueryCount   *int32  `json:"MinRewriteQueryCount,omitempty" xml:"MinRewriteQueryCount,omitempty"`
	MinRewriteQueryPattern *int32  `json:"MinRewriteQueryPattern,omitempty" xml:"MinRewriteQueryPattern,omitempty"`
	ScanQueriesCount       *int64  `json:"ScanQueriesCount,omitempty" xml:"ScanQueriesCount,omitempty"`
	SlowQueryThreshold     *int32  `json:"SlowQueryThreshold,omitempty" xml:"SlowQueryThreshold,omitempty"`
	StartTime              *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubQueriesCount        *int64  `json:"SubQueriesCount,omitempty" xml:"SubQueriesCount,omitempty"`
	SubtaskId              *int64  `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
}

func (s OpenStructMvRecommendSubTaskModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructMvRecommendSubTaskModel) GoString() string {
	return s.String()
}

func (s *OpenStructMvRecommendSubTaskModel) GetEndTime() *string {
	return s.EndTime
}

func (s *OpenStructMvRecommendSubTaskModel) GetMinRewriteQueryCount() *int32 {
	return s.MinRewriteQueryCount
}

func (s *OpenStructMvRecommendSubTaskModel) GetMinRewriteQueryPattern() *int32 {
	return s.MinRewriteQueryPattern
}

func (s *OpenStructMvRecommendSubTaskModel) GetScanQueriesCount() *int64 {
	return s.ScanQueriesCount
}

func (s *OpenStructMvRecommendSubTaskModel) GetSlowQueryThreshold() *int32 {
	return s.SlowQueryThreshold
}

func (s *OpenStructMvRecommendSubTaskModel) GetStartTime() *string {
	return s.StartTime
}

func (s *OpenStructMvRecommendSubTaskModel) GetStatus() *string {
	return s.Status
}

func (s *OpenStructMvRecommendSubTaskModel) GetSubQueriesCount() *int64 {
	return s.SubQueriesCount
}

func (s *OpenStructMvRecommendSubTaskModel) GetSubtaskId() *int64 {
	return s.SubtaskId
}

func (s *OpenStructMvRecommendSubTaskModel) SetEndTime(v string) *OpenStructMvRecommendSubTaskModel {
	s.EndTime = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetMinRewriteQueryCount(v int32) *OpenStructMvRecommendSubTaskModel {
	s.MinRewriteQueryCount = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetMinRewriteQueryPattern(v int32) *OpenStructMvRecommendSubTaskModel {
	s.MinRewriteQueryPattern = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetScanQueriesCount(v int64) *OpenStructMvRecommendSubTaskModel {
	s.ScanQueriesCount = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetSlowQueryThreshold(v int32) *OpenStructMvRecommendSubTaskModel {
	s.SlowQueryThreshold = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetStartTime(v string) *OpenStructMvRecommendSubTaskModel {
	s.StartTime = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetStatus(v string) *OpenStructMvRecommendSubTaskModel {
	s.Status = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetSubQueriesCount(v int64) *OpenStructMvRecommendSubTaskModel {
	s.SubQueriesCount = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) SetSubtaskId(v int64) *OpenStructMvRecommendSubTaskModel {
	s.SubtaskId = &v
	return s
}

func (s *OpenStructMvRecommendSubTaskModel) Validate() error {
	return dara.Validate(s)
}
