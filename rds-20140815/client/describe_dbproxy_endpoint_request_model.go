// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBProxyEndpointRequest
	GetDBInstanceId() *string
	SetDBProxyConnectString(v string) *DescribeDBProxyEndpointRequest
	GetDBProxyConnectString() *string
	SetDBProxyEndpointId(v string) *DescribeDBProxyEndpointRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *DescribeDBProxyEndpointRequest
	GetDBProxyEngineType() *string
	SetOwnerId(v int64) *DescribeDBProxyEndpointRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBProxyEndpointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBProxyEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBProxyEndpointRequest
	GetResourceOwnerId() *int64
}

type DescribeDBProxyEndpointRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1ja4f56s7us****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The proxy endpoint that you want to query. You can call the DescribeDBProxy interface to query the proxy endpoint.
	//
	// example:
	//
	// testproxy****.rwlb.rds.aliyuncs.com
	DBProxyConnectString *string `json:"DBProxyConnectString,omitempty" xml:"DBProxyConnectString,omitempty"`
	// The name of the proxy terminal. You can call the DescribeDBProxy interface to query the name of the proxy terminal.
	//
	// example:
	//
	// keaxncrjluwu0gue****
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBProxyEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBProxyEndpointRequest) GetDBProxyConnectString() *string {
	return s.DBProxyConnectString
}

func (s *DescribeDBProxyEndpointRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *DescribeDBProxyEndpointRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DescribeDBProxyEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBProxyEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBProxyEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBProxyEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBProxyEndpointRequest) SetDBInstanceId(v string) *DescribeDBProxyEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetDBProxyConnectString(v string) *DescribeDBProxyEndpointRequest {
	s.DBProxyConnectString = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetDBProxyEndpointId(v string) *DescribeDBProxyEndpointRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetDBProxyEngineType(v string) *DescribeDBProxyEndpointRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetOwnerId(v int64) *DescribeDBProxyEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetRegionId(v string) *DescribeDBProxyEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetResourceOwnerAccount(v string) *DescribeDBProxyEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) SetResourceOwnerId(v int64) *DescribeDBProxyEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBProxyEndpointRequest) Validate() error {
	return dara.Validate(s)
}
