// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonVersion(v string) *DescribeAddonMetricsRequest
	GetAddonVersion() *string
	SetAliyunLang(v string) *DescribeAddonMetricsRequest
	GetAliyunLang() *string
	SetEnvironmentType(v string) *DescribeAddonMetricsRequest
	GetEnvironmentType() *string
	SetName(v string) *DescribeAddonMetricsRequest
	GetName() *string
	SetRegionId(v string) *DescribeAddonMetricsRequest
	GetRegionId() *string
}

type DescribeAddonMetricsRequest struct {
	// The version of the component.
	//
	// example:
	//
	// 0.0.1
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The environment.
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"EnvironmentType,omitempty" xml:"EnvironmentType,omitempty"`
	// The name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAddonMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *DescribeAddonMetricsRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *DescribeAddonMetricsRequest) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *DescribeAddonMetricsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAddonMetricsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAddonMetricsRequest) SetAddonVersion(v string) *DescribeAddonMetricsRequest {
	s.AddonVersion = &v
	return s
}

func (s *DescribeAddonMetricsRequest) SetAliyunLang(v string) *DescribeAddonMetricsRequest {
	s.AliyunLang = &v
	return s
}

func (s *DescribeAddonMetricsRequest) SetEnvironmentType(v string) *DescribeAddonMetricsRequest {
	s.EnvironmentType = &v
	return s
}

func (s *DescribeAddonMetricsRequest) SetName(v string) *DescribeAddonMetricsRequest {
	s.Name = &v
	return s
}

func (s *DescribeAddonMetricsRequest) SetRegionId(v string) *DescribeAddonMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAddonMetricsRequest) Validate() error {
	return dara.Validate(s)
}
