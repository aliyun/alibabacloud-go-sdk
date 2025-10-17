// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseVocAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEnterpriseVocAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *GetEnterpriseVocAnalysisTaskResponseBodyData) *GetEnterpriseVocAnalysisTaskResponseBody
	GetData() *GetEnterpriseVocAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetEnterpriseVocAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetEnterpriseVocAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEnterpriseVocAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEnterpriseVocAnalysisTaskResponseBody
	GetSuccess() *bool
}

type GetEnterpriseVocAnalysisTaskResponseBody struct {
	// example:
	//
	// NoPermission
	Code *string                                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetEnterpriseVocAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 403
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) GetData() *GetEnterpriseVocAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) SetCode(v string) *GetEnterpriseVocAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) SetData(v *GetEnterpriseVocAnalysisTaskResponseBodyData) *GetEnterpriseVocAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetEnterpriseVocAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) SetMessage(v string) *GetEnterpriseVocAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) SetRequestId(v string) *GetEnterpriseVocAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) SetSuccess(v bool) *GetEnterpriseVocAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEnterpriseVocAnalysisTaskResponseBodyData struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// qwen-max
	ModelId            *string                                                         `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelName          *string                                                         `json:"modelName,omitempty" xml:"modelName,omitempty"`
	StatisticsOverview *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview `json:"statisticsOverview,omitempty" xml:"statisticsOverview,omitempty" type:"Struct"`
	// example:
	//
	// PENDING
	Status *string                                            `json:"status,omitempty" xml:"status,omitempty"`
	Usage  *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) GetStatisticsOverview() *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview {
	return s.StatisticsOverview
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) GetUsage() *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage {
	return s.Usage
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetEnterpriseVocAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) SetModelId(v string) *GetEnterpriseVocAnalysisTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) SetModelName(v string) *GetEnterpriseVocAnalysisTaskResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) SetStatisticsOverview(v *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) *GetEnterpriseVocAnalysisTaskResponseBodyData {
	s.StatisticsOverview = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) SetStatus(v string) *GetEnterpriseVocAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) SetUsage(v *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) *GetEnterpriseVocAnalysisTaskResponseBodyData {
	s.Usage = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyData) Validate() error {
	if s.StatisticsOverview != nil {
		if err := s.StatisticsOverview.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview struct {
	// example:
	//
	// 17
	Count                     *int32                                                                                   `json:"count,omitempty" xml:"count,omitempty"`
	FilterDimensionStatistics *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics `json:"filterDimensionStatistics,omitempty" xml:"filterDimensionStatistics,omitempty" type:"Struct"`
	TagDimensionStatistics    *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics    `json:"tagDimensionStatistics,omitempty" xml:"tagDimensionStatistics,omitempty" type:"Struct"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) GetCount() *int32 {
	return s.Count
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) GetFilterDimensionStatistics() *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics {
	return s.FilterDimensionStatistics
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) GetTagDimensionStatistics() *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics {
	return s.TagDimensionStatistics
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) SetCount(v int32) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview {
	s.Count = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) SetFilterDimensionStatistics(v *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview {
	s.FilterDimensionStatistics = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) SetTagDimensionStatistics(v *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview {
	s.TagDimensionStatistics = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview) Validate() error {
	if s.FilterDimensionStatistics != nil {
		if err := s.FilterDimensionStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.TagDimensionStatistics != nil {
		if err := s.TagDimensionStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics struct {
	TagValueCountStatistic []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic `json:"tagValueCountStatistic,omitempty" xml:"tagValueCountStatistic,omitempty" type:"Repeated"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics) GetTagValueCountStatistic() []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic {
	return s.TagValueCountStatistic
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics) SetTagValueCountStatistic(v []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics {
	s.TagValueCountStatistic = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics) Validate() error {
	if s.TagValueCountStatistic != nil {
		for _, item := range s.TagValueCountStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic struct {
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// example:
	//
	// 10
	ValueCount *int32 `json:"valueCount,omitempty" xml:"valueCount,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) GetTagName() *string {
	return s.TagName
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) GetValueCount() *int32 {
	return s.ValueCount
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) SetTagName(v string) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic {
	s.TagName = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) SetValueCount(v int32) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic {
	s.ValueCount = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) Validate() error {
	return dara.Validate(s)
}

type GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics struct {
	TagValueCountStatistic []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic `json:"tagValueCountStatistic,omitempty" xml:"tagValueCountStatistic,omitempty" type:"Repeated"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics) GetTagValueCountStatistic() []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic {
	return s.TagValueCountStatistic
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics) SetTagValueCountStatistic(v []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics {
	s.TagValueCountStatistic = v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics) Validate() error {
	if s.TagValueCountStatistic != nil {
		for _, item := range s.TagValueCountStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic struct {
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// example:
	//
	// 10
	ValueCount *int32 `json:"valueCount,omitempty" xml:"valueCount,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) GetTagName() *string {
	return s.TagName
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) GetValueCount() *int32 {
	return s.ValueCount
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) SetTagName(v string) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic {
	s.TagName = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) SetValueCount(v int32) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic {
	s.ValueCount = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) Validate() error {
	return dara.Validate(s)
}

type GetEnterpriseVocAnalysisTaskResponseBodyDataUsage struct {
	// example:
	//
	// 1
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 2
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) SetInputTokens(v int32) *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) SetOutputTokens(v int32) *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
