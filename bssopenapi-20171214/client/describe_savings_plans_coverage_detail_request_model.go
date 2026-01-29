// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageDetailRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansCoverageDetailRequest
	GetEndPeriod() *string
	SetFilterParam(v *DescribeSavingsPlansCoverageDetailRequestFilterParam) *DescribeSavingsPlansCoverageDetailRequest
	GetFilterParam() *DescribeSavingsPlansCoverageDetailRequestFilterParam
	SetMaxResults(v int32) *DescribeSavingsPlansCoverageDetailRequest
	GetMaxResults() *int32
	SetPeriodType(v string) *DescribeSavingsPlansCoverageDetailRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansCoverageDetailRequest
	GetStartPeriod() *string
	SetToken(v string) *DescribeSavingsPlansCoverageDetailRequest
	GetToken() *string
}

type DescribeSavingsPlansCoverageDetailRequest struct {
	// The ID of the account for which you want to query coverage details.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end is excluded from the time range. If you do not set this parameter, the end time is the current time. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-08-09 00:00:00
	EndPeriod   *string                                               `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParam *DescribeSavingsPlansCoverageDetailRequestFilterParam `json:"FilterParam,omitempty" xml:"FilterParam,omitempty" type:"Struct"`
	// The maximum number of entries to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity at which coverage details are queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAY
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The beginning of the time range to query. The beginning is included in the time range. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-15 13:40:45
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The token that is used to retrieve the next page of results. You do not need to set this parameter if you query coverage details within a specific time range for the first time. The response returns a token that you can use to query coverage details that are displayed on the next page. If a null value is returned for the NextToken parameter, no more coverage details can be queried.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeSavingsPlansCoverageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetFilterParam() *DescribeSavingsPlansCoverageDetailRequestFilterParam {
	return s.FilterParam
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansCoverageDetailRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansCoverageDetailRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetEndPeriod(v string) *DescribeSavingsPlansCoverageDetailRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetFilterParam(v *DescribeSavingsPlansCoverageDetailRequestFilterParam) *DescribeSavingsPlansCoverageDetailRequest {
	s.FilterParam = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetMaxResults(v int32) *DescribeSavingsPlansCoverageDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetPeriodType(v string) *DescribeSavingsPlansCoverageDetailRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetStartPeriod(v string) *DescribeSavingsPlansCoverageDetailRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) SetToken(v string) *DescribeSavingsPlansCoverageDetailRequest {
	s.Token = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequest) Validate() error {
	if s.FilterParam != nil {
		if err := s.FilterParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlansCoverageDetailRequestFilterParam struct {
	Dimensions []*DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Tags       []*DescribeSavingsPlansCoverageDetailRequestFilterParamTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansCoverageDetailRequestFilterParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailRequestFilterParam) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParam) GetDimensions() []*DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions {
	return s.Dimensions
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParam) GetTags() []*DescribeSavingsPlansCoverageDetailRequestFilterParamTags {
	return s.Tags
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParam) SetDimensions(v []*DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) *DescribeSavingsPlansCoverageDetailRequestFilterParam {
	s.Dimensions = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParam) SetTags(v []*DescribeSavingsPlansCoverageDetailRequestFilterParamTags) *DescribeSavingsPlansCoverageDetailRequestFilterParam {
	s.Tags = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParam) Validate() error {
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

type DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) SetCode(v string) *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) SetSelectType(v string) *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) SetValues(v []*string) *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageDetailRequestFilterParamTags struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansCoverageDetailRequestFilterParamTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailRequestFilterParamTags) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) SetCode(v string) *DescribeSavingsPlansCoverageDetailRequestFilterParamTags {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) SetSelectType(v string) *DescribeSavingsPlansCoverageDetailRequestFilterParamTags {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) SetValues(v []*string) *DescribeSavingsPlansCoverageDetailRequestFilterParamTags {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailRequestFilterParamTags) Validate() error {
	return dara.Validate(s)
}
