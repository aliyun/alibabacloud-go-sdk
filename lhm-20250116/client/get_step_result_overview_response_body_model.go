// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStepResultOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetStepResultOverviewResponseBodyData) *GetStepResultOverviewResponseBody
	GetData() *GetStepResultOverviewResponseBodyData
	SetErrCode(v string) *GetStepResultOverviewResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetStepResultOverviewResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetStepResultOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStepResultOverviewResponseBody
	GetSuccess() *bool
}

type GetStepResultOverviewResponseBody struct {
	// The response data.
	Data *GetStepResultOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message code.
	//
	// example:
	//
	// None
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// not supported.pos 10960, line 327, column 26, token IDENTIFIER settings
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 512AF06E-3B95-5932-81D8-717B15143359
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed. Check errCode and errMessage for troubleshooting.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetStepResultOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStepResultOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetStepResultOverviewResponseBody) GetData() *GetStepResultOverviewResponseBodyData {
	return s.Data
}

func (s *GetStepResultOverviewResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetStepResultOverviewResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetStepResultOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStepResultOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStepResultOverviewResponseBody) SetData(v *GetStepResultOverviewResponseBodyData) *GetStepResultOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetStepResultOverviewResponseBody) SetErrCode(v string) *GetStepResultOverviewResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetStepResultOverviewResponseBody) SetErrMessage(v string) *GetStepResultOverviewResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetStepResultOverviewResponseBody) SetRequestId(v string) *GetStepResultOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStepResultOverviewResponseBody) SetSuccess(v bool) *GetStepResultOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetStepResultOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStepResultOverviewResponseBodyData struct {
	// The number of validated fields.
	//
	// example:
	//
	// 10
	CheckColumnCount *int64 `json:"checkColumnCount,omitempty" xml:"checkColumnCount,omitempty"`
	// The metric name of the target.
	//
	// example:
	//
	// amount
	DstMetricName *string `json:"dstMetricName,omitempty" xml:"dstMetricName,omitempty"`
	// Indicates whether the source and target are consistent. Valid values:
	//
	// - 0: Inconsistent.
	//
	// - 1: Consistent.
	//
	// example:
	//
	// 0
	IsConsistent *int32 `json:"isConsistent,omitempty" xml:"isConsistent,omitempty"`
	// The number of validated metrics.
	//
	// example:
	//
	// 6
	MetricColumnCount *int64 `json:"metricColumnCount,omitempty" xml:"metricColumnCount,omitempty"`
	// The number of metrics that passed validation.
	//
	// example:
	//
	// 5
	MetricPassColumnCount *int64 `json:"metricPassColumnCount,omitempty" xml:"metricPassColumnCount,omitempty"`
	// The number of fields that passed validation.
	//
	// example:
	//
	// 8
	PassColumnCount *int64 `json:"passColumnCount,omitempty" xml:"passColumnCount,omitempty"`
	// The unique ID of the validation result.
	//
	// example:
	//
	// 30001
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
	// The partition name of the source.
	//
	// example:
	//
	// ds=20260116
	SourcePtName *string `json:"sourcePtName,omitempty" xml:"sourcePtName,omitempty"`
	// The table name of the source.
	//
	// example:
	//
	// table_demo
	SourceTable *string `json:"sourceTable,omitempty" xml:"sourceTable,omitempty"`
	// The metric name of the source.
	//
	// example:
	//
	// amount
	SrcMetricName *string `json:"srcMetricName,omitempty" xml:"srcMetricName,omitempty"`
	// The task status. Valid values:
	//
	// - 0: Created.
	//
	// - 1: Running.
	//
	// - 2: Completed.
	//
	// - 3: Stopped.
	//
	// - 4: Canceled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The partition name of the target.
	//
	// example:
	//
	// ds=20260116
	TargetPtName *string `json:"targetPtName,omitempty" xml:"targetPtName,omitempty"`
	// The table name of the target.
	//
	// example:
	//
	// table_demo
	TargetTable *string `json:"targetTable,omitempty" xml:"targetTable,omitempty"`
}

func (s GetStepResultOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStepResultOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStepResultOverviewResponseBodyData) GetCheckColumnCount() *int64 {
	return s.CheckColumnCount
}

func (s *GetStepResultOverviewResponseBodyData) GetDstMetricName() *string {
	return s.DstMetricName
}

func (s *GetStepResultOverviewResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *GetStepResultOverviewResponseBodyData) GetMetricColumnCount() *int64 {
	return s.MetricColumnCount
}

func (s *GetStepResultOverviewResponseBodyData) GetMetricPassColumnCount() *int64 {
	return s.MetricPassColumnCount
}

func (s *GetStepResultOverviewResponseBodyData) GetPassColumnCount() *int64 {
	return s.PassColumnCount
}

func (s *GetStepResultOverviewResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *GetStepResultOverviewResponseBodyData) GetSourcePtName() *string {
	return s.SourcePtName
}

func (s *GetStepResultOverviewResponseBodyData) GetSourceTable() *string {
	return s.SourceTable
}

func (s *GetStepResultOverviewResponseBodyData) GetSrcMetricName() *string {
	return s.SrcMetricName
}

func (s *GetStepResultOverviewResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetStepResultOverviewResponseBodyData) GetTargetPtName() *string {
	return s.TargetPtName
}

func (s *GetStepResultOverviewResponseBodyData) GetTargetTable() *string {
	return s.TargetTable
}

func (s *GetStepResultOverviewResponseBodyData) SetCheckColumnCount(v int64) *GetStepResultOverviewResponseBodyData {
	s.CheckColumnCount = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetDstMetricName(v string) *GetStepResultOverviewResponseBodyData {
	s.DstMetricName = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetIsConsistent(v int32) *GetStepResultOverviewResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetMetricColumnCount(v int64) *GetStepResultOverviewResponseBodyData {
	s.MetricColumnCount = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetMetricPassColumnCount(v int64) *GetStepResultOverviewResponseBodyData {
	s.MetricPassColumnCount = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetPassColumnCount(v int64) *GetStepResultOverviewResponseBodyData {
	s.PassColumnCount = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetResultId(v string) *GetStepResultOverviewResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetSourcePtName(v string) *GetStepResultOverviewResponseBodyData {
	s.SourcePtName = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetSourceTable(v string) *GetStepResultOverviewResponseBodyData {
	s.SourceTable = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetSrcMetricName(v string) *GetStepResultOverviewResponseBodyData {
	s.SrcMetricName = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetStatus(v int32) *GetStepResultOverviewResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetTargetPtName(v string) *GetStepResultOverviewResponseBodyData {
	s.TargetPtName = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) SetTargetTable(v string) *GetStepResultOverviewResponseBodyData {
	s.TargetTable = &v
	return s
}

func (s *GetStepResultOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
