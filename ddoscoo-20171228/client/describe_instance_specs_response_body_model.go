// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody
	GetInstanceSpecs() []*DescribeInstanceSpecsResponseBodyInstanceSpecs
	SetRequestId(v string) *DescribeInstanceSpecsResponseBody
	GetRequestId() *string
}

type DescribeInstanceSpecsResponseBody struct {
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) GetInstanceSpecs() []*DescribeInstanceSpecsResponseBodyInstanceSpecs {
	return s.InstanceSpecs
}

func (s *DescribeInstanceSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	// example:
	//
	// 20000
	BandwidthMbps *int32 `json:"BandwidthMbps,omitempty" xml:"BandwidthMbps,omitempty"`
	// example:
	//
	// 20
	BaseBandwidth *int32 `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	// example:
	//
	// 10
	DefenseCount *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	// example:
	//
	// 50
	DomainLimit *int32 `json:"DomainLimit,omitempty" xml:"DomainLimit,omitempty"`
	// example:
	//
	// 20
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// example:
	//
	// default
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 50
	PortLimit *int32 `json:"PortLimit,omitempty" xml:"PortLimit,omitempty"`
	// example:
	//
	// 1000
	QpsLimit *int32 `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty"`
	// example:
	//
	// 10
	SiteLimit *int32 `json:"SiteLimit,omitempty" xml:"SiteLimit,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetBandwidthMbps() *int32 {
	return s.BandwidthMbps
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetBaseBandwidth() *int32 {
	return s.BaseBandwidth
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetDefenseCount() *int32 {
	return s.DefenseCount
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetDomainLimit() *int32 {
	return s.DomainLimit
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetElasticBandwidth() *int32 {
	return s.ElasticBandwidth
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetFunctionVersion() *string {
	return s.FunctionVersion
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetPortLimit() *int32 {
	return s.PortLimit
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetQpsLimit() *int32 {
	return s.QpsLimit
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetSiteLimit() *int32 {
	return s.SiteLimit
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBandwidthMbps(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BandwidthMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBaseBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDomainLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DomainLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetFunctionVersion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPortLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PortLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetQpsLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.QpsLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetSiteLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.SiteLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) Validate() error {
	return dara.Validate(s)
}
