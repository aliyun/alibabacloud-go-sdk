// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *DeleteConsumerGroupRequest
	GetConsumerId() *string
	SetInstanceId(v string) *DeleteConsumerGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteConsumerGroupRequest
	GetRegionId() *string
}

type DeleteConsumerGroupRequest struct {
	// The name of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// CID-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *DeleteConsumerGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteConsumerGroupRequest) SetConsumerId(v string) *DeleteConsumerGroupRequest {
	s.ConsumerId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetInstanceId(v string) *DeleteConsumerGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetRegionId(v string) *DeleteConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
