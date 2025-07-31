// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterModificationHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *DescribeParameterModificationHistoryRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *DescribeParameterModificationHistoryRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeParameterModificationHistoryRequest
	GetEndTime() *string
	SetNodeId(v string) *DescribeParameterModificationHistoryRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeParameterModificationHistoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterModificationHistoryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeParameterModificationHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterModificationHistoryRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeParameterModificationHistoryRequest
	GetStartTime() *string
}

type DescribeParameterModificationHistoryRequest struct {
	// The role of the instance. Valid values:
	//
	// 	- **db**: shard
	//
	// 	- **cs**: Configserver
	//
	// 	- **mongos**: mongos
	//
	// 	- **logic**: sharded cluster instance
	//
	// example:
	//
	// mongos
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2235****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-02T12:10:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the mongos node or shard node whose parameter modification records you want to query in the instance.
	//
	// >  This parameter is valid only when **DBInstanceId*	- is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp1158****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-01T12:10:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeParameterModificationHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterModificationHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeParameterModificationHistoryRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParameterModificationHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeParameterModificationHistoryRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeParameterModificationHistoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterModificationHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterModificationHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterModificationHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterModificationHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeParameterModificationHistoryRequest) SetCharacterType(v string) *DescribeParameterModificationHistoryRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetDBInstanceId(v string) *DescribeParameterModificationHistoryRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetEndTime(v string) *DescribeParameterModificationHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetNodeId(v string) *DescribeParameterModificationHistoryRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetOwnerAccount(v string) *DescribeParameterModificationHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetOwnerId(v int64) *DescribeParameterModificationHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetResourceOwnerAccount(v string) *DescribeParameterModificationHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetResourceOwnerId(v int64) *DescribeParameterModificationHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetStartTime(v string) *DescribeParameterModificationHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) Validate() error {
	return dara.Validate(s)
}
