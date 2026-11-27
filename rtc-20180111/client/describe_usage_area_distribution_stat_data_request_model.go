// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageAreaDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeUsageAreaDistributionStatDataRequest
	GetAppId() *string
	SetEndDate(v string) *DescribeUsageAreaDistributionStatDataRequest
	GetEndDate() *string
	SetParentArea(v string) *DescribeUsageAreaDistributionStatDataRequest
	GetParentArea() *string
	SetStartDate(v string) *DescribeUsageAreaDistributionStatDataRequest
	GetStartDate() *string
}

type DescribeUsageAreaDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query, expressed as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The parent region name. For example, the parent of Hangzhou City is Zhejiang Province, and the parameter should be set to China_Zhejiang Province. An empty parameter indicates statistics at the global scope (country dimension), such as China or the United Kingdom.
	//
	// example:
	//
	// 中国_浙江省
	ParentArea *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
	// The start time of the query, expressed as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeUsageAreaDistributionStatDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeUsageAreaDistributionStatDataRequest) GetParentArea() *string {
	return s.ParentArea
}

func (s *DescribeUsageAreaDistributionStatDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetAppId(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetEndDate(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetStartDate(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
