// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExposedAsapVulCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedAsapVulCount() *int32
	SetExposedComponentCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedComponentCount() *int32
	SetExposedDdsCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedDdsCount() *int32
	SetExposedEcsCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedEcsCount() *int32
	SetExposedInstanceCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedInstanceCount() *int32
	SetExposedIpCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedIpCount() *int32
	SetExposedKvstoreCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedKvstoreCount() *int32
	SetExposedLaterVulCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedLaterVulCount() *int32
	SetExposedNntfVulCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedNntfVulCount() *int32
	SetExposedPortCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedPortCount() *int32
	SetExposedRdsCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedRdsCount() *int32
	SetExposedWeekPasswordMachineCount(v int32) *DescribeExposedStatisticsResponseBody
	GetExposedWeekPasswordMachineCount() *int32
	SetGatewayAssetCount(v int32) *DescribeExposedStatisticsResponseBody
	GetGatewayAssetCount() *int32
	SetRequestId(v string) *DescribeExposedStatisticsResponseBody
	GetRequestId() *string
}

type DescribeExposedStatisticsResponseBody struct {
	// The total number of high-risk vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 1
	ExposedAsapVulCount *int32 `json:"ExposedAsapVulCount,omitempty" xml:"ExposedAsapVulCount,omitempty"`
	// The total number of system components that are exposed on the Internet. The components include OpenSSL and OpenSSH.
	//
	// example:
	//
	// 7
	ExposedComponentCount *int32 `json:"ExposedComponentCount,omitempty" xml:"ExposedComponentCount,omitempty"`
	// The number of ApsaraDB for MongoDB instances that are exposed on the Internet.
	//
	// example:
	//
	// 1
	ExposedDdsCount *int32 `json:"ExposedDdsCount,omitempty" xml:"ExposedDdsCount,omitempty"`
	// The number of Elastic Compute Service (ECS) instances that are exposed on the Internet.
	//
	// example:
	//
	// 5
	ExposedEcsCount *int32 `json:"ExposedEcsCount,omitempty" xml:"ExposedEcsCount,omitempty"`
	// The total number of assets that are exposed on the Internet.
	//
	// example:
	//
	// 100
	ExposedInstanceCount *int32 `json:"ExposedInstanceCount,omitempty" xml:"ExposedInstanceCount,omitempty"`
	// The total number of IP addresses that are exposed on the Internet.
	//
	// example:
	//
	// 100
	ExposedIpCount *int32 `json:"ExposedIpCount,omitempty" xml:"ExposedIpCount,omitempty"`
	// The number of ApsaraDB for Redis instances that are exposed on the Internet.
	//
	// example:
	//
	// 3
	ExposedKvstoreCount *int32 `json:"ExposedKvstoreCount,omitempty" xml:"ExposedKvstoreCount,omitempty"`
	// The total number of medium-risk vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 5
	ExposedLaterVulCount *int32 `json:"ExposedLaterVulCount,omitempty" xml:"ExposedLaterVulCount,omitempty"`
	// The total number of low-risk vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 0
	ExposedNntfVulCount *int32 `json:"ExposedNntfVulCount,omitempty" xml:"ExposedNntfVulCount,omitempty"`
	// The total number of ports that are exposed on the Internet.
	//
	// example:
	//
	// 6
	ExposedPortCount *int32 `json:"ExposedPortCount,omitempty" xml:"ExposedPortCount,omitempty"`
	// The number of ApsaraDB RDS instances that are exposed on the Internet.
	//
	// example:
	//
	// 1
	ExposedRdsCount *int32 `json:"ExposedRdsCount,omitempty" xml:"ExposedRdsCount,omitempty"`
	// The total number of system keys that are detected on your servers and are exposed on the Internet.
	//
	// example:
	//
	// 20
	ExposedWeekPasswordMachineCount *int32 `json:"ExposedWeekPasswordMachineCount,omitempty" xml:"ExposedWeekPasswordMachineCount,omitempty"`
	// The total number of gateway assets that are exposed on the Internet. The gateway assets include NAT gateways and Server Load Balancer (SLB) instances.
	//
	// example:
	//
	// 3
	GatewayAssetCount *int32 `json:"GatewayAssetCount,omitempty" xml:"GatewayAssetCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4B897D10-B3CD-4A93-A5FA-591F3ED12A86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExposedStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedAsapVulCount() *int32 {
	return s.ExposedAsapVulCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedComponentCount() *int32 {
	return s.ExposedComponentCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedDdsCount() *int32 {
	return s.ExposedDdsCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedEcsCount() *int32 {
	return s.ExposedEcsCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedInstanceCount() *int32 {
	return s.ExposedInstanceCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedIpCount() *int32 {
	return s.ExposedIpCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedKvstoreCount() *int32 {
	return s.ExposedKvstoreCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedLaterVulCount() *int32 {
	return s.ExposedLaterVulCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedNntfVulCount() *int32 {
	return s.ExposedNntfVulCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedPortCount() *int32 {
	return s.ExposedPortCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedRdsCount() *int32 {
	return s.ExposedRdsCount
}

func (s *DescribeExposedStatisticsResponseBody) GetExposedWeekPasswordMachineCount() *int32 {
	return s.ExposedWeekPasswordMachineCount
}

func (s *DescribeExposedStatisticsResponseBody) GetGatewayAssetCount() *int32 {
	return s.GatewayAssetCount
}

func (s *DescribeExposedStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedAsapVulCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedAsapVulCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedComponentCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedComponentCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedDdsCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedDdsCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedEcsCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedEcsCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedInstanceCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedInstanceCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedIpCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedIpCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedKvstoreCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedKvstoreCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedLaterVulCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedLaterVulCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedNntfVulCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedNntfVulCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedPortCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedPortCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedRdsCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedRdsCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedWeekPasswordMachineCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedWeekPasswordMachineCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetGatewayAssetCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.GatewayAssetCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetRequestId(v string) *DescribeExposedStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
