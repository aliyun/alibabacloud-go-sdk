// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectString(v string) *ModifyDBClusterVipRequest
	GetConnectString() *string
	SetDBClusterId(v string) *ModifyDBClusterVipRequest
	GetDBClusterId() *string
	SetVPCId(v string) *ModifyDBClusterVipRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBClusterVipRequest
	GetVSwitchId() *string
}

type ModifyDBClusterVipRequest struct {
	// The connection string.
	//
	// example:
	//
	// amv-2ze8mbuai974s4y2500000169.ads.aliyuncs.com
	ConnectString *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster ID of an AnalyticDB for MySQL Data Lakehouse Edition cluster in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2ze8mbuai97*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The VPC ID.
	//
	// > - The **VPC*	- must be in the same region as the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1at5ze0t5u3xtqn****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// > - The new vSwitch must be in the same zone as the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1aadw9k19x6cis9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBClusterVipRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVipRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVipRequest) GetConnectString() *string {
	return s.ConnectString
}

func (s *ModifyDBClusterVipRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterVipRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBClusterVipRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBClusterVipRequest) SetConnectString(v string) *ModifyDBClusterVipRequest {
	s.ConnectString = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetDBClusterId(v string) *ModifyDBClusterVipRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetVPCId(v string) *ModifyDBClusterVipRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetVSwitchId(v string) *ModifyDBClusterVipRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) Validate() error {
	return dara.Validate(s)
}
