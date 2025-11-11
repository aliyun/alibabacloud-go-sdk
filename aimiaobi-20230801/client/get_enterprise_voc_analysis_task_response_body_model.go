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
	// NoData
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetEnterpriseVocAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// error
	ErrorMessage       *string                                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	StatisticsOverview *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverview `json:"StatisticsOverview,omitempty" xml:"StatisticsOverview,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESSED
	Status *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Usage  *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
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
	// 100
	Count                     *int32                                                                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	FilterDimensionStatistics *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatistics `json:"FilterDimensionStatistics,omitempty" xml:"FilterDimensionStatistics,omitempty" type:"Struct"`
	TagDimensionStatistics    *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatistics    `json:"TagDimensionStatistics,omitempty" xml:"TagDimensionStatistics,omitempty" type:"Struct"`
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
	TagValueCountStatistic []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic `json:"TagValueCountStatistic,omitempty" xml:"TagValueCountStatistic,omitempty" type:"Repeated"`
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
	// example:
	//
	// xxx
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// multiTagValues
	TagTaskType *string `json:"TagTaskType,omitempty" xml:"TagTaskType,omitempty"`
	// example:
	//
	// 100
	ValueCount *int32 `json:"ValueCount,omitempty" xml:"ValueCount,omitempty"`
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

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) GetTagTaskType() *string {
	return s.TagTaskType
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) GetValueCount() *int32 {
	return s.ValueCount
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) SetTagName(v string) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic {
	s.TagName = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic) SetTagTaskType(v string) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewFilterDimensionStatisticsTagValueCountStatistic {
	s.TagTaskType = &v
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
	TagValueCountStatistic []*GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic `json:"TagValueCountStatistic,omitempty" xml:"TagValueCountStatistic,omitempty" type:"Repeated"`
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
	// example:
	//
	// xxx
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// multiTagValues
	TagTaskType *string `json:"TagTaskType,omitempty" xml:"TagTaskType,omitempty"`
	// example:
	//
	// 100
	ValueCount *int32 `json:"ValueCount,omitempty" xml:"ValueCount,omitempty"`
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

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) GetTagTaskType() *string {
	return s.TagTaskType
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) GetValueCount() *int32 {
	return s.ValueCount
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) SetTagName(v string) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic {
	s.TagName = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic) SetTagTaskType(v string) *GetEnterpriseVocAnalysisTaskResponseBodyDataStatisticsOverviewTagDimensionStatisticsTagValueCountStatistic {
	s.TagTaskType = &v
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
	// 200
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) SetInputTokens(v int64) *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) SetOutputTokens(v int64) *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
