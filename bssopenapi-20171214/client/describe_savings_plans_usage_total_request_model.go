// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansUsageTotalRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansUsageTotalRequest
	GetEndPeriod() *string
	SetFilterParam(v *DescribeSavingsPlansUsageTotalRequestFilterParam) *DescribeSavingsPlansUsageTotalRequest
	GetFilterParam() *DescribeSavingsPlansUsageTotalRequestFilterParam
	SetPeriodType(v string) *DescribeSavingsPlansUsageTotalRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansUsageTotalRequest
	GetStartPeriod() *string
}

type DescribeSavingsPlansUsageTotalRequest struct {
	// The ID of the account for which you want to query usage summary. If you do not set this parameter, the data of the current Alibaba Cloud account and its RAM users is queried. To query the data of a RAM user, specify the ID of the RAM user.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end is excluded from the time range. If you do not set this parameter, the end time is the current time. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod   *string                                           `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParam *DescribeSavingsPlansUsageTotalRequestFilterParam `json:"FilterParam,omitempty" xml:"FilterParam,omitempty" type:"Struct"`
	// The time granularity at which usage summary are queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The beginning of the time range to query. The beginning is included in the time range. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
}

func (s DescribeSavingsPlansUsageTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansUsageTotalRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansUsageTotalRequest) GetFilterParam() *DescribeSavingsPlansUsageTotalRequestFilterParam {
	return s.FilterParam
}

func (s *DescribeSavingsPlansUsageTotalRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansUsageTotalRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansUsageTotalRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansUsageTotalRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequest) SetEndPeriod(v string) *DescribeSavingsPlansUsageTotalRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequest) SetFilterParam(v *DescribeSavingsPlansUsageTotalRequestFilterParam) *DescribeSavingsPlansUsageTotalRequest {
	s.FilterParam = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequest) SetPeriodType(v string) *DescribeSavingsPlansUsageTotalRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequest) SetStartPeriod(v string) *DescribeSavingsPlansUsageTotalRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequest) Validate() error {
	if s.FilterParam != nil {
		if err := s.FilterParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlansUsageTotalRequestFilterParam struct {
	Dimensions []*DescribeSavingsPlansUsageTotalRequestFilterParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Tags       []*DescribeSavingsPlansUsageTotalRequestFilterParamTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansUsageTotalRequestFilterParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalRequestFilterParam) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParam) GetDimensions() []*DescribeSavingsPlansUsageTotalRequestFilterParamDimensions {
	return s.Dimensions
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParam) GetTags() []*DescribeSavingsPlansUsageTotalRequestFilterParamTags {
	return s.Tags
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParam) SetDimensions(v []*DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) *DescribeSavingsPlansUsageTotalRequestFilterParam {
	s.Dimensions = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParam) SetTags(v []*DescribeSavingsPlansUsageTotalRequestFilterParamTags) *DescribeSavingsPlansUsageTotalRequestFilterParam {
	s.Tags = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParam) Validate() error {
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

type DescribeSavingsPlansUsageTotalRequestFilterParamDimensions struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) SetCode(v string) *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) SetSelectType(v string) *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) SetValues(v []*string) *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansUsageTotalRequestFilterParamTags struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansUsageTotalRequestFilterParamTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageTotalRequestFilterParamTags) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) SetCode(v string) *DescribeSavingsPlansUsageTotalRequestFilterParamTags {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) SetSelectType(v string) *DescribeSavingsPlansUsageTotalRequestFilterParamTags {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) SetValues(v []*string) *DescribeSavingsPlansUsageTotalRequestFilterParamTags {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansUsageTotalRequestFilterParamTags) Validate() error {
	return dara.Validate(s)
}
