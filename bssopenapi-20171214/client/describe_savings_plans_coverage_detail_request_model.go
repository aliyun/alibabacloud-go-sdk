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
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end time is excluded from the time range. If you do not set this parameter, the end time is the current time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	EndPeriod   *string                                               `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParam *DescribeSavingsPlansCoverageDetailRequestFilterParam `json:"FilterParam,omitempty" xml:"FilterParam,omitempty" type:"Struct"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity for querying the usage summary. Valid values: MONTH, DAY, HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The beginning of the time range to query. The start time is included in the time range. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The pagination token. Leave this parameter empty for the first page. The response returns the token for the next page. An empty token in the response indicates that all data has been returned.
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
