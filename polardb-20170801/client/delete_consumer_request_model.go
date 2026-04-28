// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *DeleteConsumerRequest
	GetConsumerId() *string
	SetGwClusterId(v string) *DeleteConsumerRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DeleteConsumerRequest
	GetRegionId() *string
}

type DeleteConsumerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *DeleteConsumerRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteConsumerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteConsumerRequest) SetConsumerId(v string) *DeleteConsumerRequest {
	s.ConsumerId = &v
	return s
}

func (s *DeleteConsumerRequest) SetGwClusterId(v string) *DeleteConsumerRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteConsumerRequest) SetRegionId(v string) *DeleteConsumerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerRequest) Validate() error {
	return dara.Validate(s)
}
