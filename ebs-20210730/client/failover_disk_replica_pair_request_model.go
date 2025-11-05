// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *FailoverDiskReplicaPairRequest
	GetClientToken() *string
	SetRegionId(v string) *FailoverDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaPairId(v string) *FailoverDiskReplicaPairRequest
	GetReplicaPairId() *string
}

type FailoverDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query region IDs of secondary disks in replication pairs.
	//
	// >  The failover feature must be enabled for the region where the secondary disk is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-cn-dsa****
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s FailoverDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s FailoverDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FailoverDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *FailoverDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *FailoverDiskReplicaPairRequest) SetClientToken(v string) *FailoverDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetRegionId(v string) *FailoverDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetReplicaPairId(v string) *FailoverDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
