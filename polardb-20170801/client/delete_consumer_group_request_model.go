// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupName(v string) *DeleteConsumerGroupRequest
	GetConsumerGroupName() *string
	SetGwClusterId(v string) *DeleteConsumerGroupRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DeleteConsumerGroupRequest
	GetRegionId() *string
}

type DeleteConsumerGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cg-xxxxxxx
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *DeleteConsumerGroupRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteConsumerGroupRequest) SetConsumerGroupName(v string) *DeleteConsumerGroupRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetGwClusterId(v string) *DeleteConsumerGroupRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetRegionId(v string) *DeleteConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
