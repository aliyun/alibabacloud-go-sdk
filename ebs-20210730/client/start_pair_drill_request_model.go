// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPairDrillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartPairDrillRequest
	GetClientToken() *string
	SetPairId(v string) *StartPairDrillRequest
	GetPairId() *string
	SetRegionId(v string) *StartPairDrillRequest
	GetRegionId() *string
}

type StartPairDrillRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query a list of replication pairs, including replication pair IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-xxxx
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
	// The region ID of the secondary disk in the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query the region in which the secondary disk of the replication pair resides.
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

func (s StartPairDrillRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPairDrillRequest) GoString() string {
	return s.String()
}

func (s *StartPairDrillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartPairDrillRequest) GetPairId() *string {
	return s.PairId
}

func (s *StartPairDrillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartPairDrillRequest) SetClientToken(v string) *StartPairDrillRequest {
	s.ClientToken = &v
	return s
}

func (s *StartPairDrillRequest) SetPairId(v string) *StartPairDrillRequest {
	s.PairId = &v
	return s
}

func (s *StartPairDrillRequest) SetRegionId(v string) *StartPairDrillRequest {
	s.RegionId = &v
	return s
}

func (s *StartPairDrillRequest) Validate() error {
	return dara.Validate(s)
}
