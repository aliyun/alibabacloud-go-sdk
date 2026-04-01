// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DeleteDBInstanceReplicationRequest
	GetChannelName() *string
	SetDbInstanceId(v string) *DeleteDBInstanceReplicationRequest
	GetDbInstanceId() *string
	SetOwnerId(v int64) *DeleteDBInstanceReplicationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDBInstanceReplicationRequest
	GetRegionId() *string
}

type DeleteDBInstanceReplicationRequest struct {
	// 复制通道名称，用于标识需要删除的复制链路
	//
	// example:
	//
	// replication-channel-001
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// 目标RDS实例ID，复制链路将从此实例上删除
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1234567890abcdef
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// 阿里云账号ID，用于指定资源的所有者
	//
	// example:
	//
	// 1234567890123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 地域ID，表示RDS实例所在的地域
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBInstanceReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceReplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceReplicationRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DeleteDBInstanceReplicationRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DeleteDBInstanceReplicationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBInstanceReplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBInstanceReplicationRequest) SetChannelName(v string) *DeleteDBInstanceReplicationRequest {
	s.ChannelName = &v
	return s
}

func (s *DeleteDBInstanceReplicationRequest) SetDbInstanceId(v string) *DeleteDBInstanceReplicationRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DeleteDBInstanceReplicationRequest) SetOwnerId(v int64) *DeleteDBInstanceReplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceReplicationRequest) SetRegionId(v string) *DeleteDBInstanceReplicationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceReplicationRequest) Validate() error {
	return dara.Validate(s)
}
