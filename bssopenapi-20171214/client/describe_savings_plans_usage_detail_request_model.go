// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansUsageDetailRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailRequest
	GetEndPeriod() *string
	SetFilterParam(v *DescribeSavingsPlansUsageDetailRequestFilterParam) *DescribeSavingsPlansUsageDetailRequest
	GetFilterParam() *DescribeSavingsPlansUsageDetailRequestFilterParam
	SetMaxResults(v int32) *DescribeSavingsPlansUsageDetailRequest
	GetMaxResults() *int32
	SetPeriodType(v string) *DescribeSavingsPlansUsageDetailRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailRequest
	GetStartPeriod() *string
	SetToken(v string) *DescribeSavingsPlansUsageDetailRequest
	GetToken() *string
}

type DescribeSavingsPlansUsageDetailRequest struct {
	// The ID of the account whose data you want to query. If this parameter is left empty, the data of the current account and its trusteeship linked accounts is queried. You can specify the ID of a trusteeship linked account to query only the data of that account.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end time of the query. Format: yyyy-MM-dd HH:mm:ss. The specified time is excluded from the query. If this parameter is left empty, the current time is used.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The filter parameters.
	FilterParam *DescribeSavingsPlansUsageDetailRequestFilterParam `json:"FilterParam,omitempty" xml:"FilterParam,omitempty" type:"Struct"`
	// The maximum number of data records to return in the current request. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity for querying the usage summary. Valid values:
	//
	// - MONTH: monthly
	//
	// - DAY: daily
	//
	// - HOUR: hourly
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The start time of the query. Format: yyyy-MM-dd HH:mm:ss. The specified time is included in the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The pagination token. Leave this parameter empty when you query the first page. The token for the next page is returned in the response. If the returned token is empty, all results are returned.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetFilterParam() *DescribeSavingsPlansUsageDetailRequestFilterParam {
	return s.FilterParam
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansUsageDetailRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetFilterParam(v *DescribeSavingsPlansUsageDetailRequestFilterParam) *DescribeSavingsPlansUsageDetailRequest {
	s.FilterParam = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetMaxResults(v int32) *DescribeSavingsPlansUsageDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetPeriodType(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetToken(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.Token = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) Validate() error {
	if s.FilterParam != nil {
		if err := s.FilterParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlansUsageDetailRequestFilterParam struct {
	// The list of dimension filter conditions.
	Dimensions []*DescribeSavingsPlansUsageDetailRequestFilterParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The list of tag filter conditions.
	Tags []*DescribeSavingsPlansUsageDetailRequestFilterParamTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansUsageDetailRequestFilterParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailRequestFilterParam) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParam) GetDimensions() []*DescribeSavingsPlansUsageDetailRequestFilterParamDimensions {
	return s.Dimensions
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParam) GetTags() []*DescribeSavingsPlansUsageDetailRequestFilterParamTags {
	return s.Tags
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParam) SetDimensions(v []*DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) *DescribeSavingsPlansUsageDetailRequestFilterParam {
	s.Dimensions = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParam) SetTags(v []*DescribeSavingsPlansUsageDetailRequestFilterParamTags) *DescribeSavingsPlansUsageDetailRequestFilterParam {
	s.Tags = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParam) Validate() error {
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

type DescribeSavingsPlansUsageDetailRequestFilterParamDimensions struct {
	// The dimension code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The selection method for the dimension filter.
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The list of dimension filter values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) SetCode(v string) *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) SetSelectType(v string) *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) SetValues(v []*string) *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansUsageDetailRequestFilterParamTags struct {
	// The tag code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The selection method for the tag filter.
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The list of tag filter values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlansUsageDetailRequestFilterParamTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailRequestFilterParamTags) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) GetValues() []*string {
	return s.Values
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) SetCode(v string) *DescribeSavingsPlansUsageDetailRequestFilterParamTags {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) SetSelectType(v string) *DescribeSavingsPlansUsageDetailRequestFilterParamTags {
	s.SelectType = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) SetValues(v []*string) *DescribeSavingsPlansUsageDetailRequestFilterParamTags {
	s.Values = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequestFilterParamTags) Validate() error {
	return dara.Validate(s)
}
