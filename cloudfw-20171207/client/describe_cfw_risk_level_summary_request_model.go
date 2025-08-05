// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCfwRiskLevelSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *DescribeCfwRiskLevelSummaryRequest
	GetInstanceType() *string
	SetLang(v string) *DescribeCfwRiskLevelSummaryRequest
	GetLang() *string
	SetRegionId(v string) *DescribeCfwRiskLevelSummaryRequest
	GetRegionId() *string
}

type DescribeCfwRiskLevelSummaryRequest struct {
	// The instance type.
	//
	// example:
	//
	// EcsEIP
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of your Cloud Firewall.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCfwRiskLevelSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCfwRiskLevelSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCfwRiskLevelSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCfwRiskLevelSummaryRequest) SetInstanceType(v string) *DescribeCfwRiskLevelSummaryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryRequest) SetLang(v string) *DescribeCfwRiskLevelSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryRequest) SetRegionId(v string) *DescribeCfwRiskLevelSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryRequest) Validate() error {
	return dara.Validate(s)
}
