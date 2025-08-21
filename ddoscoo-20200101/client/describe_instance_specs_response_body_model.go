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
	// The details of the specifications of the instance.
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4E3A9B5F-5DDB-593D-A1E6-F1F451DB5E0B
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
	// The clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	BandwidthMbps *int32 `json:"BandwidthMbps,omitempty" xml:"BandwidthMbps,omitempty"`
	// The basic protection bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 30
	BaseBandwidth *int32 `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	// The specification of concurrent connections of the instance.
	//
	// example:
	//
	// 100000
	ConnLimit *int64 `json:"ConnLimit,omitempty" xml:"ConnLimit,omitempty"`
	// The specification of new connections of the instance.
	//
	// example:
	//
	// 5000
	CpsLimit *int64 `json:"CpsLimit,omitempty" xml:"CpsLimit,omitempty"`
	// The number of available advanced mitigation sessions for this month. **-1**: unlimited
	//
	// >  This parameter is returned only when the request parameter **RegionId*	- is set to **ap-southeast-1**. If RegionId is set to ap-southeast-1, the specifications of Anti-DDoS Proxy (Outside Chinese Mainland) instances are queried.
	//
	// example:
	//
	// 2
	DefenseCount *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	// The number of domain names that can be protected by the instance.
	//
	// example:
	//
	// 50
	DomainLimit *int32 `json:"DomainLimit,omitempty" xml:"DomainLimit,omitempty"`
	// The burstable protection bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 30
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// The burstable clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 300
	ElasticBw *int32 `json:"ElasticBw,omitempty" xml:"ElasticBw,omitempty"`
	// The metering method of the burstable clean bandwidth. Valid values:
	//
	// 	- **day**: the metering method of daily 95th percentile
	//
	// 	- **month**: the metering method of monthly 95th percentile
	//
	// example:
	//
	// day
	ElasticBwModel *string `json:"ElasticBwModel,omitempty" xml:"ElasticBwModel,omitempty"`
	// The burstable QPS. Unit: QPS
	//
	// example:
	//
	// 10
	ElasticQps *int64 `json:"ElasticQps,omitempty" xml:"ElasticQps,omitempty"`
	// The metering method of the burstable QPS. Valid values:
	//
	// 	- **day**: the metering method of daily 95th percentile
	//
	// 	- **month**: the metering method of monthly 95th percentile
	//
	// example:
	//
	// day
	ElasticQpsMode *string `json:"ElasticQpsMode,omitempty" xml:"ElasticQpsMode,omitempty"`
	// The function plan of the instance. Valid values:
	//
	// 	- **default**: Standard
	//
	// 	- **enhance**: Enhanced
	//
	// 	- **cnhk**: Chinese Mainland Acceleration (CMA)
	//
	// 	- **cnhk_default**: Secure Chinese Mainland Acceleration (Sec-CMA) standard
	//
	// 	- **cnhk_enhance**: Sec-CMA enhanced
	//
	// example:
	//
	// default
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-zvp2eibz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of ports that can be protected by the instance.
	//
	// example:
	//
	// 50
	PortLimit *int32 `json:"PortLimit,omitempty" xml:"PortLimit,omitempty"`
	// The clean QPS.
	//
	// example:
	//
	// 3000
	QpsLimit *int32 `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty"`
	// The threshold of the clean bandwidth. Valid values: 0 to 15360. The value 0 indicates that rate limiting is never triggered. Unit: Mbit/s
	//
	// example:
	//
	// 0
	RealLimitBw *int64 `json:"RealLimitBw,omitempty" xml:"RealLimitBw,omitempty"`
	// The number of sites that can be protected by the instance.
	//
	// example:
	//
	// 5
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetConnLimit() *int64 {
	return s.ConnLimit
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetCpsLimit() *int64 {
	return s.CpsLimit
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetElasticBw() *int32 {
	return s.ElasticBw
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetElasticBwModel() *string {
	return s.ElasticBwModel
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetElasticQps() *int64 {
	return s.ElasticQps
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetElasticQpsMode() *string {
	return s.ElasticQpsMode
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) GetRealLimitBw() *int64 {
	return s.RealLimitBw
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetConnLimit(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ConnLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetCpsLimit(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.CpsLimit = &v
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticBw(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticBw = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticBwModel(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticBwModel = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticQps(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticQps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticQpsMode(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticQpsMode = &v
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetRealLimitBw(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.RealLimitBw = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetSiteLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.SiteLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) Validate() error {
	return dara.Validate(s)
}
