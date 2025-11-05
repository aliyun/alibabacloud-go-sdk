// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReplicaGroupDrillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartReplicaGroupDrillRequest
	GetClientToken() *string
	SetGroupId(v string) *StartReplicaGroupDrillRequest
	GetGroupId() *string
	SetRegionId(v string) *StartReplicaGroupDrillRequest
	GetRegionId() *string
}

type StartReplicaGroupDrillRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the replication pair-consistent group ID. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation the most recent list of async replication pair-consistent groups, including group IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region where the secondary site in the replication pair-consistent group is located. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query the region where the secondary site in the replication pair-consistent group is located.
	//
	// >  You must enable the disaster recovery drill feature in the region in which the secondary site resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartReplicaGroupDrillRequest) String() string {
	return dara.Prettify(s)
}

func (s StartReplicaGroupDrillRequest) GoString() string {
	return s.String()
}

func (s *StartReplicaGroupDrillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartReplicaGroupDrillRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *StartReplicaGroupDrillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartReplicaGroupDrillRequest) SetClientToken(v string) *StartReplicaGroupDrillRequest {
	s.ClientToken = &v
	return s
}

func (s *StartReplicaGroupDrillRequest) SetGroupId(v string) *StartReplicaGroupDrillRequest {
	s.GroupId = &v
	return s
}

func (s *StartReplicaGroupDrillRequest) SetRegionId(v string) *StartReplicaGroupDrillRequest {
	s.RegionId = &v
	return s
}

func (s *StartReplicaGroupDrillRequest) Validate() error {
	return dara.Validate(s)
}
