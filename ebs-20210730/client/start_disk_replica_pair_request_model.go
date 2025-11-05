// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartDiskReplicaPairRequest
	GetClientToken() *string
	SetOneShot(v bool) *StartDiskReplicaPairRequest
	GetOneShot() *bool
	SetRegionId(v string) *StartDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaPairId(v string) *StartDiskReplicaPairRequest
	GetReplicaPairId() *string
}

type StartDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to immediately synchronize data. Valid values:
	//
	// 	- true: immediately synchronizes data.
	//
	// 	- false: synchronizes data based on the recovery point objective (RPO).
	//
	// Default value: false.
	//
	// example:
	//
	// false
	OneShot *bool `json:"OneShot,omitempty" xml:"OneShot,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query the region information of replication pairs.
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

func (s StartDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDiskReplicaPairRequest) GetOneShot() *bool {
	return s.OneShot
}

func (s *StartDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *StartDiskReplicaPairRequest) SetClientToken(v string) *StartDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetOneShot(v bool) *StartDiskReplicaPairRequest {
	s.OneShot = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetRegionId(v string) *StartDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetReplicaPairId(v string) *StartDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *StartDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
