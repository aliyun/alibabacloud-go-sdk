// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExceptionRate(v []*ListDataDiagnosisReportsResponseBodyExceptionRate) *ListDataDiagnosisReportsResponseBody
	GetExceptionRate() []*ListDataDiagnosisReportsResponseBodyExceptionRate
	SetReportsOfAbnormalBehavior(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) *ListDataDiagnosisReportsResponseBody
	GetReportsOfAbnormalBehavior() [][]*ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior
	SetReportsOfBaseStatistics(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) *ListDataDiagnosisReportsResponseBody
	GetReportsOfBaseStatistics() [][]*ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics
	SetReportsOfChangeRateData(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) *ListDataDiagnosisReportsResponseBody
	GetReportsOfChangeRateData() [][]*ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData
	SetReportsOfJoinTables(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) *ListDataDiagnosisReportsResponseBody
	GetReportsOfJoinTables() [][]*ListDataDiagnosisReportsResponseBodyReportsOfJoinTables
	SetReportsOfPreferenceStatisticsCycle(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) *ListDataDiagnosisReportsResponseBody
	GetReportsOfPreferenceStatisticsCycle() [][]*ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle
	SetRequestId(v string) *ListDataDiagnosisReportsResponseBody
	GetRequestId() *string
	SetType(v string) *ListDataDiagnosisReportsResponseBody
	GetType() *string
}

type ListDataDiagnosisReportsResponseBody struct {
	ExceptionRate                      []*ListDataDiagnosisReportsResponseBodyExceptionRate                        `json:"ExceptionRate,omitempty" xml:"ExceptionRate,omitempty" type:"Repeated"`
	ReportsOfAbnormalBehavior          [][]*ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior          `json:"ReportsOfAbnormalBehavior,omitempty" xml:"ReportsOfAbnormalBehavior,omitempty" type:"Repeated"`
	ReportsOfBaseStatistics            [][]*ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics            `json:"ReportsOfBaseStatistics,omitempty" xml:"ReportsOfBaseStatistics,omitempty" type:"Repeated"`
	ReportsOfChangeRateData            [][]*ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData            `json:"ReportsOfChangeRateData,omitempty" xml:"ReportsOfChangeRateData,omitempty" type:"Repeated"`
	ReportsOfJoinTables                [][]*ListDataDiagnosisReportsResponseBodyReportsOfJoinTables                `json:"ReportsOfJoinTables,omitempty" xml:"ReportsOfJoinTables,omitempty" type:"Repeated"`
	ReportsOfPreferenceStatisticsCycle [][]*ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle `json:"ReportsOfPreferenceStatisticsCycle,omitempty" xml:"ReportsOfPreferenceStatisticsCycle,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ChangeRate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBody) GetExceptionRate() []*ListDataDiagnosisReportsResponseBodyExceptionRate {
	return s.ExceptionRate
}

func (s *ListDataDiagnosisReportsResponseBody) GetReportsOfAbnormalBehavior() [][]*ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	return s.ReportsOfAbnormalBehavior
}

func (s *ListDataDiagnosisReportsResponseBody) GetReportsOfBaseStatistics() [][]*ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	return s.ReportsOfBaseStatistics
}

func (s *ListDataDiagnosisReportsResponseBody) GetReportsOfChangeRateData() [][]*ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData {
	return s.ReportsOfChangeRateData
}

func (s *ListDataDiagnosisReportsResponseBody) GetReportsOfJoinTables() [][]*ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	return s.ReportsOfJoinTables
}

func (s *ListDataDiagnosisReportsResponseBody) GetReportsOfPreferenceStatisticsCycle() [][]*ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	return s.ReportsOfPreferenceStatisticsCycle
}

func (s *ListDataDiagnosisReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataDiagnosisReportsResponseBody) GetType() *string {
	return s.Type
}

func (s *ListDataDiagnosisReportsResponseBody) SetExceptionRate(v []*ListDataDiagnosisReportsResponseBodyExceptionRate) *ListDataDiagnosisReportsResponseBody {
	s.ExceptionRate = v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetReportsOfAbnormalBehavior(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) *ListDataDiagnosisReportsResponseBody {
	s.ReportsOfAbnormalBehavior = v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetReportsOfBaseStatistics(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) *ListDataDiagnosisReportsResponseBody {
	s.ReportsOfBaseStatistics = v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetReportsOfChangeRateData(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) *ListDataDiagnosisReportsResponseBody {
	s.ReportsOfChangeRateData = v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetReportsOfJoinTables(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) *ListDataDiagnosisReportsResponseBody {
	s.ReportsOfJoinTables = v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetReportsOfPreferenceStatisticsCycle(v [][]*ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) *ListDataDiagnosisReportsResponseBody {
	s.ReportsOfPreferenceStatisticsCycle = v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetRequestId(v string) *ListDataDiagnosisReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) SetType(v string) *ListDataDiagnosisReportsResponseBody {
	s.Type = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBody) Validate() error {
	if s.ExceptionRate != nil {
		for _, item := range s.ExceptionRate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataDiagnosisReportsResponseBodyExceptionRate struct {
	// example:
	//
	// add
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// WARN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBodyExceptionRate) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBodyExceptionRate) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) GetGroup() *string {
	return s.Group
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) GetMessage() *string {
	return s.Message
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) GetType() *string {
	return s.Type
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) SetGroup(v string) *ListDataDiagnosisReportsResponseBodyExceptionRate {
	s.Group = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) SetMessage(v string) *ListDataDiagnosisReportsResponseBodyExceptionRate {
	s.Message = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) SetType(v string) *ListDataDiagnosisReportsResponseBodyExceptionRate {
	s.Type = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyExceptionRate) Validate() error {
	return dara.Validate(s)
}

type ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior struct {
	// example:
	//
	// 20250114
	Ds *string `json:"Ds,omitempty" xml:"Ds,omitempty"`
	// example:
	//
	// 1
	RankId *string `json:"RankId,omitempty" xml:"RankId,omitempty"`
	// example:
	//
	// 1.0
	ConversionRate *string `json:"ConversionRate,omitempty" xml:"ConversionRate,omitempty"`
	// example:
	//
	// 100010050+259203779
	ConversionRateIds *string `json:"ConversionRateIds,omitempty" xml:"ConversionRateIds,omitempty"`
	// example:
	//
	// 2.0
	DownStreamCount *string `json:"DownStreamCount,omitempty" xml:"DownStreamCount,omitempty"`
	// example:
	//
	// 189814043+272292277
	DownStreamCountIds *string `json:"DownStreamCountIds,omitempty" xml:"DownStreamCountIds,omitempty"`
	// example:
	//
	// pair
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// 2.0
	UpStreamCount *string `json:"UpStreamCount,omitempty" xml:"UpStreamCount,omitempty"`
	// example:
	//
	// 104684044+249445882
	UpStreamCountIds *string `json:"UpStreamCountIds,omitempty" xml:"UpStreamCountIds,omitempty"`
	// example:
	//
	// {678.8225: 91, 5270.4675: 95}
	Distribution *string `json:"Distribution,omitempty" xml:"Distribution,omitempty"`
	// example:
	//
	// conversion_rate
	IndicatorName *string `json:"IndicatorName,omitempty" xml:"IndicatorName,omitempty"`
	// example:
	//
	// 0.0
	ExceptionRate *string `json:"ExceptionRate,omitempty" xml:"ExceptionRate,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetDs() *string {
	return s.Ds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetRankId() *string {
	return s.RankId
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetConversionRate() *string {
	return s.ConversionRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetConversionRateIds() *string {
	return s.ConversionRateIds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetDownStreamCount() *string {
	return s.DownStreamCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetDownStreamCountIds() *string {
	return s.DownStreamCountIds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetGranularity() *string {
	return s.Granularity
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetUpStreamCount() *string {
	return s.UpStreamCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetUpStreamCountIds() *string {
	return s.UpStreamCountIds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetDistribution() *string {
	return s.Distribution
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetIndicatorName() *string {
	return s.IndicatorName
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) GetExceptionRate() *string {
	return s.ExceptionRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetDs(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.Ds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetRankId(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.RankId = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetConversionRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.ConversionRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetConversionRateIds(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.ConversionRateIds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetDownStreamCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.DownStreamCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetDownStreamCountIds(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.DownStreamCountIds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetGranularity(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.Granularity = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetUpStreamCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.UpStreamCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetUpStreamCountIds(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.UpStreamCountIds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetDistribution(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.Distribution = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetIndicatorName(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.IndicatorName = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) SetExceptionRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior {
	s.ExceptionRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfAbnormalBehavior) Validate() error {
	return dara.Validate(s)
}

type ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics struct {
	// example:
	//
	// \\\\N
	DefaultNullCount *string `json:"DefaultNullCount,omitempty" xml:"DefaultNullCount,omitempty"`
	// example:
	//
	// \\\\N
	DefaultNullRate *string `json:"DefaultNullRate,omitempty" xml:"DefaultNullRate,omitempty"`
	// example:
	//
	// 20230509
	Ds *string `json:"Ds,omitempty" xml:"Ds,omitempty"`
	// example:
	//
	// register_time
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// string
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// 55095
	NullCount *string `json:"NullCount,omitempty" xml:"NullCount,omitempty"`
	// example:
	//
	// 0.5580879448141732
	NullRate *string `json:"NullRate,omitempty" xml:"NullRate,omitempty"`
	// example:
	//
	// 98721
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 29
	UniqueCount *string `json:"UniqueCount,omitempty" xml:"UniqueCount,omitempty"`
	// example:
	//
	// 52.0
	ValueMax *string `json:"ValueMax,omitempty" xml:"ValueMax,omitempty"`
	// example:
	//
	// 35.0
	ValueMedian *string `json:"ValueMedian,omitempty" xml:"ValueMedian,omitempty"`
	// example:
	//
	// 18.0
	ValueMin *string `json:"ValueMin,omitempty" xml:"ValueMin,omitempty"`
	// example:
	//
	// 18.0
	ValueQuantile1 *string `json:"ValueQuantile1,omitempty" xml:"ValueQuantile1,omitempty"`
	// example:
	//
	// 18.0
	ValueQuantile5 *string `json:"ValueQuantile5,omitempty" xml:"ValueQuantile5,omitempty"`
	// example:
	//
	// 18.0
	ValueQuantile25 *string `json:"ValueQuantile25,omitempty" xml:"ValueQuantile25,omitempty"`
	// example:
	//
	// 18.0
	ValueQuantile75 *string `json:"ValueQuantile75,omitempty" xml:"ValueQuantile75,omitempty"`
	// example:
	//
	// 18.0
	ValueQuantile95 *string `json:"ValueQuantile95,omitempty" xml:"ValueQuantile95,omitempty"`
	// example:
	//
	// 18.0
	ValueQuantile99 *string `json:"ValueQuantile99,omitempty" xml:"ValueQuantile99,omitempty"`
	// example:
	//
	// 3
	Rn *string `json:"Rn,omitempty" xml:"Rn,omitempty"`
	// example:
	//
	// 91149.0
	FrequencyMax *string `json:"FrequencyMax,omitempty" xml:"FrequencyMax,omitempty"`
	// example:
	//
	// 1349.0
	FrequencyMedian *string `json:"FrequencyMedian,omitempty" xml:"FrequencyMedian,omitempty"`
	// example:
	//
	// 289.0
	FrequencyMin *string `json:"FrequencyMin,omitempty" xml:"FrequencyMin,omitempty"`
	// example:
	//
	// 289.0
	FrequencyQuantile1 *string `json:"FrequencyQuantile1,omitempty" xml:"FrequencyQuantile1,omitempty"`
	// example:
	//
	// 289.0
	FrequencyQuantile5 *string `json:"FrequencyQuantile5,omitempty" xml:"FrequencyQuantile5,omitempty"`
	// example:
	//
	// 289.0
	FrequencyQuantile25 *string `json:"FrequencyQuantile25,omitempty" xml:"FrequencyQuantile25,omitempty"`
	// example:
	//
	// 289.0
	FrequencyQuantile75 *string `json:"FrequencyQuantile75,omitempty" xml:"FrequencyQuantile75,omitempty"`
	// example:
	//
	// 289.0
	FrequencyQuantile95 *string `json:"FrequencyQuantile95,omitempty" xml:"FrequencyQuantile95,omitempty"`
	// example:
	//
	// 289.0
	FrequencyQuantile99 *string `json:"FrequencyQuantile99,omitempty" xml:"FrequencyQuantile99,omitempty"`
	// example:
	//
	// {678.8225: 91, 5270.4675: 95}
	Distribution *string `json:"Distribution,omitempty" xml:"Distribution,omitempty"`
	// example:
	//
	// 1
	RankId *string `json:"RankId,omitempty" xml:"RankId,omitempty"`
	// example:
	//
	// 1683562246
	FeatureValue *string `json:"FeatureValue,omitempty" xml:"FeatureValue,omitempty"`
	// example:
	//
	// 1
	ValueCount *string `json:"ValueCount,omitempty" xml:"ValueCount,omitempty"`
	// example:
	//
	// 0.000019996000799840032
	ValuePercent *string `json:"ValuePercent,omitempty" xml:"ValuePercent,omitempty"`
	// example:
	//
	// 0.7261657444926671
	ValueQuantile *string `json:"ValueQuantile,omitempty" xml:"ValueQuantile,omitempty"`
	// example:
	//
	// 427
	FeatureFrequency *string `json:"FeatureFrequency,omitempty" xml:"FeatureFrequency,omitempty"`
	// example:
	//
	// 1
	FrequencyCount *string `json:"FrequencyCount,omitempty" xml:"FrequencyCount,omitempty"`
	// example:
	//
	// 0.5
	FrequencyPercent *string `json:"FrequencyPercent,omitempty" xml:"FrequencyPercent,omitempty"`
	// example:
	//
	// 5
	FrequencyQuantile *string `json:"FrequencyQuantile,omitempty" xml:"FrequencyQuantile,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetDefaultNullCount() *string {
	return s.DefaultNullCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetDefaultNullRate() *string {
	return s.DefaultNullRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetDs() *string {
	return s.Ds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFeatureType() *string {
	return s.FeatureType
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetNullCount() *string {
	return s.NullCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetNullRate() *string {
	return s.NullRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetUniqueCount() *string {
	return s.UniqueCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueMax() *string {
	return s.ValueMax
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueMedian() *string {
	return s.ValueMedian
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueMin() *string {
	return s.ValueMin
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile1() *string {
	return s.ValueQuantile1
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile5() *string {
	return s.ValueQuantile5
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile25() *string {
	return s.ValueQuantile25
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile75() *string {
	return s.ValueQuantile75
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile95() *string {
	return s.ValueQuantile95
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile99() *string {
	return s.ValueQuantile99
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetRn() *string {
	return s.Rn
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyMax() *string {
	return s.FrequencyMax
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyMedian() *string {
	return s.FrequencyMedian
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyMin() *string {
	return s.FrequencyMin
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile1() *string {
	return s.FrequencyQuantile1
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile5() *string {
	return s.FrequencyQuantile5
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile25() *string {
	return s.FrequencyQuantile25
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile75() *string {
	return s.FrequencyQuantile75
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile95() *string {
	return s.FrequencyQuantile95
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile99() *string {
	return s.FrequencyQuantile99
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetDistribution() *string {
	return s.Distribution
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetRankId() *string {
	return s.RankId
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFeatureValue() *string {
	return s.FeatureValue
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueCount() *string {
	return s.ValueCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValuePercent() *string {
	return s.ValuePercent
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetValueQuantile() *string {
	return s.ValueQuantile
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFeatureFrequency() *string {
	return s.FeatureFrequency
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyCount() *string {
	return s.FrequencyCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyPercent() *string {
	return s.FrequencyPercent
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) GetFrequencyQuantile() *string {
	return s.FrequencyQuantile
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetDefaultNullCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.DefaultNullCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetDefaultNullRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.DefaultNullRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetDs(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.Ds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFeatureName(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FeatureName = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFeatureType(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FeatureType = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetNullCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.NullCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetNullRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.NullRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetTotalCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.TotalCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetUniqueCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.UniqueCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueMax(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueMax = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueMedian(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueMedian = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueMin(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueMin = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile1(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile1 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile5(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile5 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile25(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile25 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile75(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile75 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile95(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile95 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile99(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile99 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetRn(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.Rn = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyMax(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyMax = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyMedian(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyMedian = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyMin(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyMin = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile1(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile1 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile5(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile5 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile25(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile25 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile75(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile75 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile95(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile95 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile99(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile99 = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetDistribution(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.Distribution = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetRankId(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.RankId = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFeatureValue(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FeatureValue = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValuePercent(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValuePercent = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetValueQuantile(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.ValueQuantile = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFeatureFrequency(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FeatureFrequency = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyPercent(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyPercent = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) SetFrequencyQuantile(v string) *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics {
	s.FrequencyQuantile = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfBaseStatistics) Validate() error {
	return dara.Validate(s)
}

type ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData struct {
	// example:
	//
	// 20230509
	Ds *string `json:"Ds,omitempty" xml:"Ds,omitempty"`
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// 1231
	ChangeCount *string `json:"ChangeCount,omitempty" xml:"ChangeCount,omitempty"`
	// example:
	//
	// 0.1231
	ChangeRate *string `json:"ChangeRate,omitempty" xml:"ChangeRate,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) GetDs() *string {
	return s.Ds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) GetFlag() *string {
	return s.Flag
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) GetChangeCount() *string {
	return s.ChangeCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) GetChangeRate() *string {
	return s.ChangeRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) SetDs(v string) *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData {
	s.Ds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) SetFlag(v string) *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData {
	s.Flag = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) SetChangeCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData {
	s.ChangeCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) SetChangeRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData {
	s.ChangeRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfChangeRateData) Validate() error {
	return dara.Validate(s)
}

type ListDataDiagnosisReportsResponseBodyReportsOfJoinTables struct {
	// example:
	//
	// 20230509
	Ds *string `json:"Ds,omitempty" xml:"Ds,omitempty"`
	// example:
	//
	// user_id
	JoinField *string `json:"JoinField,omitempty" xml:"JoinField,omitempty"`
	// example:
	//
	// 0.53
	LeftExceptRate *string `json:"LeftExceptRate,omitempty" xml:"LeftExceptRate,omitempty"`
	// example:
	//
	// 0.0
	RightExceptRate *string `json:"RightExceptRate,omitempty" xml:"RightExceptRate,omitempty"`
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// register_time
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// 1683562246
	FeatureValue *string `json:"FeatureValue,omitempty" xml:"FeatureValue,omitempty"`
	// example:
	//
	// 1
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// example:
	//
	// 1
	Quantile *string `json:"Quantile,omitempty" xml:"Quantile,omitempty"`
	// example:
	//
	// 1
	ValueCount *string `json:"ValueCount,omitempty" xml:"ValueCount,omitempty"`
	// example:
	//
	// 0.019996000799
	ValuePercent *string `json:"ValuePercent,omitempty" xml:"ValuePercent,omitempty"`
	// example:
	//
	// 0.72616
	ValueQuantile *string `json:"ValueQuantile,omitempty" xml:"ValueQuantile,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetDs() *string {
	return s.Ds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetJoinField() *string {
	return s.JoinField
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetLeftExceptRate() *string {
	return s.LeftExceptRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetRightExceptRate() *string {
	return s.RightExceptRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetFlag() *string {
	return s.Flag
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetFeatureValue() *string {
	return s.FeatureValue
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetPercent() *string {
	return s.Percent
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetQuantile() *string {
	return s.Quantile
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetValueCount() *string {
	return s.ValueCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetValuePercent() *string {
	return s.ValuePercent
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) GetValueQuantile() *string {
	return s.ValueQuantile
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetDs(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.Ds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetJoinField(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.JoinField = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetLeftExceptRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.LeftExceptRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetRightExceptRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.RightExceptRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetFlag(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.Flag = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetFeatureName(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.FeatureName = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetFeatureValue(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.FeatureValue = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetPercent(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.Percent = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetQuantile(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.Quantile = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetValueCount(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.ValueCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetValuePercent(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.ValuePercent = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) SetValueQuantile(v string) *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables {
	s.ValueQuantile = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfJoinTables) Validate() error {
	return dara.Validate(s)
}

type ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle struct {
	// example:
	//
	// 0.73
	CycleRemainRate *string `json:"CycleRemainRate,omitempty" xml:"CycleRemainRate,omitempty"`
	// example:
	//
	// 0.52
	SingleRemainRate *string `json:"SingleRemainRate,omitempty" xml:"SingleRemainRate,omitempty"`
	// example:
	//
	// 20230509
	Ds *string `json:"Ds,omitempty" xml:"Ds,omitempty"`
	// example:
	//
	// 7
	Days *string `json:"Days,omitempty" xml:"Days,omitempty"`
	// example:
	//
	// 0.67
	EverAppearedRate *string `json:"EverAppearedRate,omitempty" xml:"EverAppearedRate,omitempty"`
	// example:
	//
	// week
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 0.33
	PeriodRemainRate *string `json:"PeriodRemainRate,omitempty" xml:"PeriodRemainRate,omitempty"`
	// example:
	//
	// 1
	PeriodRemainCount *int64 `json:"PeriodRemainCount,omitempty" xml:"PeriodRemainCount,omitempty"`
	// example:
	//
	// 1
	PeriodInternal *int64 `json:"PeriodInternal,omitempty" xml:"PeriodInternal,omitempty"`
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetCycleRemainRate() *string {
	return s.CycleRemainRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetSingleRemainRate() *string {
	return s.SingleRemainRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetDs() *string {
	return s.Ds
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetDays() *string {
	return s.Days
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetEverAppearedRate() *string {
	return s.EverAppearedRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetPeriod() *string {
	return s.Period
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetPeriodRemainRate() *string {
	return s.PeriodRemainRate
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetPeriodRemainCount() *int64 {
	return s.PeriodRemainCount
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) GetPeriodInternal() *int64 {
	return s.PeriodInternal
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetCycleRemainRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.CycleRemainRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetSingleRemainRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.SingleRemainRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetDs(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.Ds = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetDays(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.Days = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetEverAppearedRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.EverAppearedRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetPeriod(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.Period = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetPeriodRemainRate(v string) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.PeriodRemainRate = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetPeriodRemainCount(v int64) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.PeriodRemainCount = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) SetPeriodInternal(v int64) *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle {
	s.PeriodInternal = &v
	return s
}

func (s *ListDataDiagnosisReportsResponseBodyReportsOfPreferenceStatisticsCycle) Validate() error {
	return dara.Validate(s)
}
