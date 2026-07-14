// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []map[string]*string) *PreviewPipelineResponseBody
	GetData() []map[string]*string
	SetMeta(v *PreviewPipelineResponseBodyMeta) *PreviewPipelineResponseBody
	GetMeta() *PreviewPipelineResponseBodyMeta
	SetRequestId(v string) *PreviewPipelineResponseBody
	GetRequestId() *string
}

type PreviewPipelineResponseBody struct {
	// The `data` field is a collection of sample rows (an array of maps) that contains only the first N rows (up to 5 by default) and does not reflect the complete write plan.
	Data []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The query metadata.
	Meta *PreviewPipelineResponseBodyMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PreviewPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewPipelineResponseBody) GetData() []map[string]*string {
	return s.Data
}

func (s *PreviewPipelineResponseBody) GetMeta() *PreviewPipelineResponseBodyMeta {
	return s.Meta
}

func (s *PreviewPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewPipelineResponseBody) SetData(v []map[string]*string) *PreviewPipelineResponseBody {
	s.Data = v
	return s
}

func (s *PreviewPipelineResponseBody) SetMeta(v *PreviewPipelineResponseBodyMeta) *PreviewPipelineResponseBody {
	s.Meta = v
	return s
}

func (s *PreviewPipelineResponseBody) SetRequestId(v string) *PreviewPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewPipelineResponseBody) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewPipelineResponseBodyMeta struct {
	// The aggregation analysis SPL statement.
	//
	// example:
	//
	// 	- | SELECT status, count(*) AS cnt GROUP BY status
	AggQuery *string `json:"aggQuery,omitempty" xml:"aggQuery,omitempty"`
	// The `meta.columnTypes` field provides a mapping from column names to data types (string / long / double / json).
	ColumnTypes []*string `json:"columnTypes,omitempty" xml:"columnTypes,omitempty" type:"Repeated"`
	// The number of matched log entries.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The number of CPU cores consumed.
	//
	// example:
	//
	// 2
	CpuCores *int32 `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The CPU time consumed, in seconds.
	//
	// example:
	//
	// 0.5
	CpuSec *float64 `json:"cpuSec,omitempty" xml:"cpuSec,omitempty"`
	// The query duration, in milliseconds.
	//
	// example:
	//
	// 1200
	ElapsedMillisecond *int64 `json:"elapsedMillisecond,omitempty" xml:"elapsedMillisecond,omitempty"`
	// Indicates whether the query is an SQL query.
	HasSQL *bool `json:"hasSQL,omitempty" xml:"hasSQL,omitempty"`
	// Indicates whether nanosecond-level ordering is enabled.
	IsAccurate *bool `json:"isAccurate,omitempty" xml:"isAccurate,omitempty"`
	// The list of result column names.
	Keys []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	// The maximum number of result rows returned.
	//
	// example:
	//
	// 5
	Limited *int32 `json:"limited,omitempty" xml:"limited,omitempty"`
	// The query mode identifier.
	//
	// example:
	//
	// 1
	Mode *int32 `json:"mode,omitempty" xml:"mode,omitempty"`
	// The number of processed data bytes.
	//
	// example:
	//
	// 524288
	ProcessedBytes *int64 `json:"processedBytes,omitempty" xml:"processedBytes,omitempty"`
	// The number of processed log rows.
	//
	// example:
	//
	// 10000
	ProcessedRows *int64 `json:"processedRows,omitempty" xml:"processedRows,omitempty"`
	// The SLS query progress. A value of Complete indicates that the query is complete.
	//
	// example:
	//
	// Complete
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
	// The number of raw data bytes scanned.
	//
	// example:
	//
	// 1048576
	ScanBytes *int64 `json:"scanBytes,omitempty" xml:"scanBytes,omitempty"`
	// The column type and aggregation information.
	Terms []map[string]interface{} `json:"terms,omitempty" xml:"terms,omitempty" type:"Repeated"`
	// The filter condition SPL statement.
	//
	// example:
	//
	// status: 200
	WhereQuery *string `json:"whereQuery,omitempty" xml:"whereQuery,omitempty"`
}

func (s PreviewPipelineResponseBodyMeta) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineResponseBodyMeta) GoString() string {
	return s.String()
}

func (s *PreviewPipelineResponseBodyMeta) GetAggQuery() *string {
	return s.AggQuery
}

func (s *PreviewPipelineResponseBodyMeta) GetColumnTypes() []*string {
	return s.ColumnTypes
}

func (s *PreviewPipelineResponseBodyMeta) GetCount() *int32 {
	return s.Count
}

func (s *PreviewPipelineResponseBodyMeta) GetCpuCores() *int32 {
	return s.CpuCores
}

func (s *PreviewPipelineResponseBodyMeta) GetCpuSec() *float64 {
	return s.CpuSec
}

func (s *PreviewPipelineResponseBodyMeta) GetElapsedMillisecond() *int64 {
	return s.ElapsedMillisecond
}

func (s *PreviewPipelineResponseBodyMeta) GetHasSQL() *bool {
	return s.HasSQL
}

func (s *PreviewPipelineResponseBodyMeta) GetIsAccurate() *bool {
	return s.IsAccurate
}

func (s *PreviewPipelineResponseBodyMeta) GetKeys() []*string {
	return s.Keys
}

func (s *PreviewPipelineResponseBodyMeta) GetLimited() *int32 {
	return s.Limited
}

func (s *PreviewPipelineResponseBodyMeta) GetMode() *int32 {
	return s.Mode
}

func (s *PreviewPipelineResponseBodyMeta) GetProcessedBytes() *int64 {
	return s.ProcessedBytes
}

func (s *PreviewPipelineResponseBodyMeta) GetProcessedRows() *int64 {
	return s.ProcessedRows
}

func (s *PreviewPipelineResponseBodyMeta) GetProgress() *string {
	return s.Progress
}

func (s *PreviewPipelineResponseBodyMeta) GetScanBytes() *int64 {
	return s.ScanBytes
}

func (s *PreviewPipelineResponseBodyMeta) GetTerms() []map[string]interface{} {
	return s.Terms
}

func (s *PreviewPipelineResponseBodyMeta) GetWhereQuery() *string {
	return s.WhereQuery
}

func (s *PreviewPipelineResponseBodyMeta) SetAggQuery(v string) *PreviewPipelineResponseBodyMeta {
	s.AggQuery = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetColumnTypes(v []*string) *PreviewPipelineResponseBodyMeta {
	s.ColumnTypes = v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetCount(v int32) *PreviewPipelineResponseBodyMeta {
	s.Count = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetCpuCores(v int32) *PreviewPipelineResponseBodyMeta {
	s.CpuCores = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetCpuSec(v float64) *PreviewPipelineResponseBodyMeta {
	s.CpuSec = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetElapsedMillisecond(v int64) *PreviewPipelineResponseBodyMeta {
	s.ElapsedMillisecond = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetHasSQL(v bool) *PreviewPipelineResponseBodyMeta {
	s.HasSQL = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetIsAccurate(v bool) *PreviewPipelineResponseBodyMeta {
	s.IsAccurate = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetKeys(v []*string) *PreviewPipelineResponseBodyMeta {
	s.Keys = v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetLimited(v int32) *PreviewPipelineResponseBodyMeta {
	s.Limited = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetMode(v int32) *PreviewPipelineResponseBodyMeta {
	s.Mode = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetProcessedBytes(v int64) *PreviewPipelineResponseBodyMeta {
	s.ProcessedBytes = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetProcessedRows(v int64) *PreviewPipelineResponseBodyMeta {
	s.ProcessedRows = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetProgress(v string) *PreviewPipelineResponseBodyMeta {
	s.Progress = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetScanBytes(v int64) *PreviewPipelineResponseBodyMeta {
	s.ScanBytes = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetTerms(v []map[string]interface{}) *PreviewPipelineResponseBodyMeta {
	s.Terms = v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) SetWhereQuery(v string) *PreviewPipelineResponseBodyMeta {
	s.WhereQuery = &v
	return s
}

func (s *PreviewPipelineResponseBodyMeta) Validate() error {
	return dara.Validate(s)
}
