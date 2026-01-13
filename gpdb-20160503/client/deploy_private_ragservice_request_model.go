// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPrivateRAGServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeployPrivateRAGServiceRequest
	GetDBInstanceId() *string
	SetVSwitchId(v string) *DeployPrivateRAGServiceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *DeployPrivateRAGServiceRequest
	GetZoneId() *string
}

type DeployPrivateRAGServiceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The vSwitch ID.
	//
	// >   The zone where the **vSwitch*	- resides must be the same as the zone that is specified by **ZoneId**.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available regions and zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeployPrivateRAGServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployPrivateRAGServiceRequest) GoString() string {
	return s.String()
}

func (s *DeployPrivateRAGServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeployPrivateRAGServiceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeployPrivateRAGServiceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeployPrivateRAGServiceRequest) SetDBInstanceId(v string) *DeployPrivateRAGServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeployPrivateRAGServiceRequest) SetVSwitchId(v string) *DeployPrivateRAGServiceRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeployPrivateRAGServiceRequest) SetZoneId(v string) *DeployPrivateRAGServiceRequest {
	s.ZoneId = &v
	return s
}

func (s *DeployPrivateRAGServiceRequest) Validate() error {
	return dara.Validate(s)
}
