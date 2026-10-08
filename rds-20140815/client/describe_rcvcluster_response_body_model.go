// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCVClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRCVClusterResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeRCVClusterResponseBody
	GetClusterName() *string
	SetMysqlOperator(v *DescribeRCVClusterResponseBodyMysqlOperator) *DescribeRCVClusterResponseBody
	GetMysqlOperator() *DescribeRCVClusterResponseBodyMysqlOperator
	SetRegion(v string) *DescribeRCVClusterResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeRCVClusterResponseBody
	GetRequestId() *string
	SetSupportDiskPerformanceLevel(v []*string) *DescribeRCVClusterResponseBody
	GetSupportDiskPerformanceLevel() []*string
	SetVClusterStatus(v string) *DescribeRCVClusterResponseBody
	GetVClusterStatus() *string
	SetVpcId(v string) *DescribeRCVClusterResponseBody
	GetVpcId() *string
}

type DescribeRCVClusterResponseBody struct {
	ClusterId                   *string                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName                 *string                                      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	MysqlOperator               *DescribeRCVClusterResponseBodyMysqlOperator `json:"MysqlOperator,omitempty" xml:"MysqlOperator,omitempty" type:"Struct"`
	Region                      *string                                      `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId                   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportDiskPerformanceLevel []*string                                    `json:"SupportDiskPerformanceLevel,omitempty" xml:"SupportDiskPerformanceLevel,omitempty" type:"Repeated"`
	VClusterStatus              *string                                      `json:"VClusterStatus,omitempty" xml:"VClusterStatus,omitempty"`
	VpcId                       *string                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCVClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCVClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRCVClusterResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeRCVClusterResponseBody) GetMysqlOperator() *DescribeRCVClusterResponseBodyMysqlOperator {
	return s.MysqlOperator
}

func (s *DescribeRCVClusterResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeRCVClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCVClusterResponseBody) GetSupportDiskPerformanceLevel() []*string {
	return s.SupportDiskPerformanceLevel
}

func (s *DescribeRCVClusterResponseBody) GetVClusterStatus() *string {
	return s.VClusterStatus
}

func (s *DescribeRCVClusterResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCVClusterResponseBody) SetClusterId(v string) *DescribeRCVClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetClusterName(v string) *DescribeRCVClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetMysqlOperator(v *DescribeRCVClusterResponseBodyMysqlOperator) *DescribeRCVClusterResponseBody {
	s.MysqlOperator = v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetRegion(v string) *DescribeRCVClusterResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetRequestId(v string) *DescribeRCVClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetSupportDiskPerformanceLevel(v []*string) *DescribeRCVClusterResponseBody {
	s.SupportDiskPerformanceLevel = v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetVClusterStatus(v string) *DescribeRCVClusterResponseBody {
	s.VClusterStatus = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) SetVpcId(v string) *DescribeRCVClusterResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeRCVClusterResponseBody) Validate() error {
	if s.MysqlOperator != nil {
		if err := s.MysqlOperator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCVClusterResponseBodyMysqlOperator struct {
	DashboardPublicEndpoint *string `json:"DashboardPublicEndpoint,omitempty" xml:"DashboardPublicEndpoint,omitempty"`
	DashboardUsername       *string `json:"DashboardUsername,omitempty" xml:"DashboardUsername,omitempty"`
	DashboardVpcEndpoint    *string `json:"DashboardVpcEndpoint,omitempty" xml:"DashboardVpcEndpoint,omitempty"`
	DeployTime              *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRCVClusterResponseBodyMysqlOperator) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCVClusterResponseBodyMysqlOperator) GoString() string {
	return s.String()
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) GetDashboardPublicEndpoint() *string {
	return s.DashboardPublicEndpoint
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) GetDashboardVpcEndpoint() *string {
	return s.DashboardVpcEndpoint
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) GetDeployTime() *string {
	return s.DeployTime
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) SetDashboardPublicEndpoint(v string) *DescribeRCVClusterResponseBodyMysqlOperator {
	s.DashboardPublicEndpoint = &v
	return s
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) SetDashboardUsername(v string) *DescribeRCVClusterResponseBodyMysqlOperator {
	s.DashboardUsername = &v
	return s
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) SetDashboardVpcEndpoint(v string) *DescribeRCVClusterResponseBodyMysqlOperator {
	s.DashboardVpcEndpoint = &v
	return s
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) SetDeployTime(v string) *DescribeRCVClusterResponseBodyMysqlOperator {
	s.DeployTime = &v
	return s
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) SetStatus(v string) *DescribeRCVClusterResponseBodyMysqlOperator {
	s.Status = &v
	return s
}

func (s *DescribeRCVClusterResponseBodyMysqlOperator) Validate() error {
	return dara.Validate(s)
}
