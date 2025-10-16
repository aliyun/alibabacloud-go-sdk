// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskSecurityGroupDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRiskSecurityGroupDetailRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeRiskSecurityGroupDetailRequest
	GetInstanceName() *string
	SetLang(v string) *DescribeRiskSecurityGroupDetailRequest
	GetLang() *string
	SetPageNo(v string) *DescribeRiskSecurityGroupDetailRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeRiskSecurityGroupDetailRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeRiskSecurityGroupDetailRequest
	GetRegionId() *string
	SetRuleUuid(v string) *DescribeRiskSecurityGroupDetailRequest
	GetRuleUuid() *string
	SetSourceIp(v string) *DescribeRiskSecurityGroupDetailRequest
	GetSourceIp() *string
}

type DescribeRiskSecurityGroupDetailRequest struct {
	// example:
	//
	// sg-2vc0p803vgxumn6r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test-instance-name
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3b1adf8d-1949-4c8e-809b-fb92ee11****
	RuleUuid *string `json:"RuleUuid,omitempty" xml:"RuleUuid,omitempty"`
	// example:
	//
	// 220.189.117.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRiskSecurityGroupDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskSecurityGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetRuleUuid() *string {
	return s.RuleUuid
}

func (s *DescribeRiskSecurityGroupDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetInstanceId(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetInstanceName(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetLang(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetPageNo(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetPageSize(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetRegionId(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetRuleUuid(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.RuleUuid = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) SetSourceIp(v string) *DescribeRiskSecurityGroupDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailRequest) Validate() error {
	return dara.Validate(s)
}
