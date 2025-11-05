// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReprotectDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReprotectDiskReplicaPairRequest
	GetClientToken() *string
	SetRegionId(v string) *ReprotectDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaPairId(v string) *ReprotectDiskReplicaPairRequest
	GetReplicaPairId() *string
	SetReverseReplicate(v bool) *ReprotectDiskReplicaPairRequest
	GetReverseReplicate() *bool
}

type ReprotectDiskReplicaPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query region IDs of secondary disks in replication pairs.
	//
	// >  The reverse replication feature must be enabled from the region where the secondary disk is located.
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
	// Specifies whether to enable the reverse replication sub-feature. Valid values: true and false. Default value: true.
	//
	// example:
	//
	// true
	ReverseReplicate *bool `json:"ReverseReplicate,omitempty" xml:"ReverseReplicate,omitempty"`
}

func (s ReprotectDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s ReprotectDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReprotectDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReprotectDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *ReprotectDiskReplicaPairRequest) GetReverseReplicate() *bool {
	return s.ReverseReplicate
}

func (s *ReprotectDiskReplicaPairRequest) SetClientToken(v string) *ReprotectDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetRegionId(v string) *ReprotectDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetReplicaPairId(v string) *ReprotectDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetReverseReplicate(v bool) *ReprotectDiskReplicaPairRequest {
	s.ReverseReplicate = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
