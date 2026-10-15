// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageTotalRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest
	GetEndPeriod() *string
	SetFilterParam(v *DescribeSavingsPlansCoverageTotalRequestFilterParam) *DescribeSavingsPlansCoverageTotalRequest
	GetFilterParam() *DescribeSavingsPlansCoverageTotalRequestFilterParam
	SetPeriodType(v string) *DescribeSavingsPlansCoverageTotalRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest
	GetStartPeriod() *string
}

type DescribeSavingsPlansCoverageTotalRequest struct {
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// Format: yyyy-MM-dd HH:mm:ss. The query excludes this time. If this parameter is left empty, the query covers data up to the current time.
	//
	// example:
	//
	// 2021-07-20 00:00:00
	EndPeriod   *string                                              `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParam *DescribeSavingsPlansCoverageTotalRequestFilterParam `json:"FilterParam,omitempty" xml:"FilterParam,omitempty" type:"Struct"`
	// Valid values:
	//
	// - MONTH: month granularity.
	//
	// - DAY: day granularity.
	//
	// - HOUR: hour granularity.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// Format: yyyy-MM-dd HH:mm:ss. The query includes this time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-15 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeSavingsPlansCoverageTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetFilterParam() *DescribeSavingsPlansCoverageTotalRequestFilterParam {
	return s.FilterParam
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansCoverageTotalRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageTotalRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetEndPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetFilterParam(v *DescribeSavingsPlansCoverageTotalRequestFilterParam) *DescribeSavingsPlansCoverageTotalRequest {
	s.FilterParam = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetPeriodType(v string) *DescribeSavingsPlansCoverageTotalRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) SetStartPeriod(v string) *DescribeSavingsPlansCoverageTotalRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequest) Validate() error {
	if s.FilterParam != nil {
		if err := s.FilterParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlansCoverageTotalRequestFilterParam struct {
	Dimensions []*DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Tags       []*DescribeSavingsPlansCoverageTotalRequestFilterParamTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansCoverageTotalRequestFilterParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalRequestFilterParam) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParam) GetDimensions() []*DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions {
	return s.Dimensions
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParam) GetTags() []*DescribeSavingsPlansCoverageTotalRequestFilterParamTags {
	return s.Tags
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParam) SetDimensions(v []*DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) *DescribeSavingsPlansCoverageTotalRequestFilterParam {
	s.Dimensions = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParam) SetTags(v []*DescribeSavingsPlansCoverageTotalRequestFilterParamTags) *DescribeSavingsPlansCoverageTotalRequestFilterParam {
	s.Tags = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParam) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) SetCode(v string) *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) SetSelectType(v string) *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) SetValues(v []*string) *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageTotalRequestFilterParamTags struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansCoverageTotalRequestFilterParamTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalRequestFilterParamTags) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) SetCode(v string) *DescribeSavingsPlansCoverageTotalRequestFilterParamTags {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) SetSelectType(v string) *DescribeSavingsPlansCoverageTotalRequestFilterParamTags {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) SetValues(v []*string) *DescribeSavingsPlansCoverageTotalRequestFilterParamTags {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalRequestFilterParamTags) Validate() error {
	return dara.Validate(s)
}
