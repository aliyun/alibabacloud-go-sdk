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
	// The ID of the account for which you want to query usage details. If you do not set this parameter, the data of the current Alibaba Cloud account and its RAM users is queried. To query the data of a RAM user, specify the ID of the RAM user.
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
	EndPeriod   *string                                            `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	FilterParam *DescribeSavingsPlansUsageDetailRequestFilterParam `json:"FilterParam,omitempty" xml:"FilterParam,omitempty" type:"Struct"`
	// The maximum number of entries to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity at which usage details are queried. Valid values: MONTH, DAY, and HOUR.
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
	// The token that is used to retrieve the next page of results. You do not need to set this parameter if you query usage details within a specific time range for the first time. The response returns a token that you can use to query usage details that are displayed on the next page. If a null value is returned for the NextToken parameter, no more coverage details can be queried.
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
	Dimensions []*DescribeSavingsPlansUsageDetailRequestFilterParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Tags       []*DescribeSavingsPlansUsageDetailRequestFilterParamTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
