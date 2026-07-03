// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestions(v []*ListDataIngestionsResponseBodyDataIngestions) *ListDataIngestionsResponseBody
	GetDataIngestions() []*ListDataIngestionsResponseBodyDataIngestions
	SetRequestId(v string) *ListDataIngestionsResponseBody
	GetRequestId() *string
}

type ListDataIngestionsResponseBody struct {
	// The list of data ingestions.
	DataIngestions []*ListDataIngestionsResponseBodyDataIngestions `json:"DataIngestions,omitempty" xml:"DataIngestions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataIngestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataIngestionsResponseBody) GetDataIngestions() []*ListDataIngestionsResponseBodyDataIngestions {
	return s.DataIngestions
}

func (s *ListDataIngestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataIngestionsResponseBody) SetDataIngestions(v []*ListDataIngestionsResponseBodyDataIngestions) *ListDataIngestionsResponseBody {
	s.DataIngestions = v
	return s
}

func (s *ListDataIngestionsResponseBody) SetRequestId(v string) *ListDataIngestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataIngestionsResponseBody) Validate() error {
	if s.DataIngestions != nil {
		for _, item := range s.DataIngestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataIngestionsResponseBodyDataIngestions struct {
	// The time when the data ingestion was enabled.
	//
	// example:
	//
	// 1733269771123
	ActiveTime *int64 `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	// The number of associated security capabilities.
	//
	// example:
	//
	// 3
	CapacityCount *int32 `json:"CapacityCount,omitempty" xml:"CapacityCount,omitempty"`
	// The time when the data ingestion was created.
	//
	// example:
	//
	// 1733269771123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the data ingestion.
	//
	// example:
	//
	// alibaba_cloud_sas_netstat_ingestion_173326*******
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
	// The mode of the data ingestion. Valid values:
	//
	// - realtime
	//
	// - scan
	//
	// example:
	//
	// realtime
	DataIngestionMode *string `json:"DataIngestionMode,omitempty" xml:"DataIngestionMode,omitempty"`
	// Indicates whether the data ingestion mode is editable.
	//
	// example:
	//
	// true
	DataIngestionModeEditable *bool `json:"DataIngestionModeEditable,omitempty" xml:"DataIngestionModeEditable,omitempty"`
	// The state of the data ingestion. Valid values:
	//
	// - ingested
	//
	// - uningested
	//
	// - abnormal
	//
	// example:
	//
	// ingested
	DataIngestionState *string `json:"DataIngestionState,omitempty" xml:"DataIngestionState,omitempty"`
	// The error code for the data ingestion anomaly.
	//
	// example:
	//
	// UserUnauthorized
	DataIngestionStateCode *string `json:"DataIngestionStateCode,omitempty" xml:"DataIngestionStateCode,omitempty"`
	// The status of the data ingestion. Valid values:
	//
	// - enabled: The data ingestion is enabled.
	//
	// - disabled: The data ingestion is disabled.
	//
	// example:
	//
	// enabled
	DataIngestionStatus *string `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	// The ID of the data ingestion template.
	//
	// example:
	//
	// alibaba_cloud_sas_netstat_ingestion
	DataIngestionTemplateId *string `json:"DataIngestionTemplateId,omitempty" xml:"DataIngestionTemplateId,omitempty"`
	// The type of the data ingestion. Valid values:
	//
	// - preset
	//
	// - custom
	//
	// example:
	//
	// preset
	DataIngestionType *string `json:"DataIngestionType,omitempty" xml:"DataIngestionType,omitempty"`
	// Indicates whether the data source is editable.
	//
	// example:
	//
	// true
	DataSourceEditable *bool `json:"DataSourceEditable,omitempty" xml:"DataSourceEditable,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// ds-scpfegri73oyoknbc90c
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Indicates whether the normalization rule is editable.
	//
	// example:
	//
	// true
	NormalizationRuleEditable *bool `json:"NormalizationRuleEditable,omitempty" xml:"NormalizationRuleEditable,omitempty"`
	// The ID of the normalization rule.
	//
	// example:
	//
	// nr-0aywiqdtaqdvwac7xkbjsf3a
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// The ID of the data source for real-time ingestion.
	//
	// example:
	//
	// ds-scpfegri73oyoknbc90c
	RealtimeDataSourceId *string `json:"RealtimeDataSourceId,omitempty" xml:"RealtimeDataSourceId,omitempty"`
	// The ID of the data source for scan-based ingestion.
	//
	// example:
	//
	// ds-scpfegri73oyoknbc90c
	ScanDataSourceId *string `json:"ScanDataSourceId,omitempty" xml:"ScanDataSourceId,omitempty"`
	// The job ID of the stream processing task.
	//
	// example:
	//
	// 73a78aa245e3b1299d6ceed093de7bd8
	StreamJobId *string `json:"StreamJobId,omitempty" xml:"StreamJobId,omitempty"`
	// The time when the data ingestion was last updated.
	//
	// example:
	//
	// 1733269771123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataIngestionsResponseBodyDataIngestions) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionsResponseBodyDataIngestions) GoString() string {
	return s.String()
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetActiveTime() *int64 {
	return s.ActiveTime
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetCapacityCount() *int32 {
	return s.CapacityCount
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionMode() *string {
	return s.DataIngestionMode
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionModeEditable() *bool {
	return s.DataIngestionModeEditable
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionState() *string {
	return s.DataIngestionState
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionStateCode() *string {
	return s.DataIngestionStateCode
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionStatus() *string {
	return s.DataIngestionStatus
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionTemplateId() *string {
	return s.DataIngestionTemplateId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataIngestionType() *string {
	return s.DataIngestionType
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataSourceEditable() *bool {
	return s.DataSourceEditable
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetNormalizationRuleEditable() *bool {
	return s.NormalizationRuleEditable
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetRealtimeDataSourceId() *string {
	return s.RealtimeDataSourceId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetScanDataSourceId() *string {
	return s.ScanDataSourceId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetStreamJobId() *string {
	return s.StreamJobId
}

func (s *ListDataIngestionsResponseBodyDataIngestions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetActiveTime(v int64) *ListDataIngestionsResponseBodyDataIngestions {
	s.ActiveTime = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetCapacityCount(v int32) *ListDataIngestionsResponseBodyDataIngestions {
	s.CapacityCount = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetCreateTime(v int64) *ListDataIngestionsResponseBodyDataIngestions {
	s.CreateTime = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionMode(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionMode = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionModeEditable(v bool) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionModeEditable = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionState(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionState = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionStateCode(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionStateCode = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionStatus(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionStatus = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionTemplateId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionTemplateId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataIngestionType(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataIngestionType = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataSourceEditable(v bool) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataSourceEditable = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetDataSourceId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.DataSourceId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetNormalizationRuleEditable(v bool) *ListDataIngestionsResponseBodyDataIngestions {
	s.NormalizationRuleEditable = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetNormalizationRuleId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetRealtimeDataSourceId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.RealtimeDataSourceId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetScanDataSourceId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.ScanDataSourceId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetStreamJobId(v string) *ListDataIngestionsResponseBodyDataIngestions {
	s.StreamJobId = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) SetUpdateTime(v int64) *ListDataIngestionsResponseBodyDataIngestions {
	s.UpdateTime = &v
	return s
}

func (s *ListDataIngestionsResponseBodyDataIngestions) Validate() error {
	return dara.Validate(s)
}
