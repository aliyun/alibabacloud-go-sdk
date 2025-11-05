// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReprotectDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReprotectDiskReplicaGroupRequest
	GetClientToken() *string
	SetRegionId(v string) *ReprotectDiskReplicaGroupRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *ReprotectDiskReplicaGroupRequest
	GetReplicaGroupId() *string
	SetReverseReplicate(v bool) *ReprotectDiskReplicaGroupRequest
	GetReverseReplicate() *bool
}

type ReprotectDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query the IDs of replication pair-consistent groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query the IDs of replication pair-consistent groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-myreplica****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// Specifies whether to enable the reverse replication sub-feature. Valid values: true and false. Default value: true.
	//
	// example:
	//
	// true
	ReverseReplicate *bool `json:"ReverseReplicate,omitempty" xml:"ReverseReplicate,omitempty"`
}

func (s ReprotectDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ReprotectDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReprotectDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReprotectDiskReplicaGroupRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *ReprotectDiskReplicaGroupRequest) GetReverseReplicate() *bool {
	return s.ReverseReplicate
}

func (s *ReprotectDiskReplicaGroupRequest) SetClientToken(v string) *ReprotectDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetRegionId(v string) *ReprotectDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetReplicaGroupId(v string) *ReprotectDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetReverseReplicate(v bool) *ReprotectDiskReplicaGroupRequest {
	s.ReverseReplicate = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) Validate() error {
	return dara.Validate(s)
}
