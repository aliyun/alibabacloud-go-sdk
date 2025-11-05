// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaPairProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDiskReplicaPairProgressRequest
	GetRegionId() *string
	SetReplicaPairId(v string) *DescribeDiskReplicaPairProgressRequest
	GetReplicaPairId() *string
}

type DescribeDiskReplicaPairProgressRequest struct {
	// The region ID of the replication pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html)operation to query the IDs of existing replication pairs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-cn-tl32ribst0z
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s DescribeDiskReplicaPairProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairProgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskReplicaPairProgressRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *DescribeDiskReplicaPairProgressRequest) SetRegionId(v string) *DescribeDiskReplicaPairProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressRequest) SetReplicaPairId(v string) *DescribeDiskReplicaPairProgressRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressRequest) Validate() error {
	return dara.Validate(s)
}
