// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDiskReplicaPairRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaPairId(v string) *DeleteDiskReplicaPairRequest
	GetReplicaPairId() *string
}

type DeleteDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the primary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query the region information of replication pairs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
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

func (s DeleteDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *DeleteDiskReplicaPairRequest) SetClientToken(v string) *DeleteDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetRegionId(v string) *DeleteDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetReplicaPairId(v string) *DeleteDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
