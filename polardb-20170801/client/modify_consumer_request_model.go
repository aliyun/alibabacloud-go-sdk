// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupName(v string) *ModifyConsumerRequest
	GetConsumerGroupName() *string
	SetConsumerId(v string) *ModifyConsumerRequest
	GetConsumerId() *string
	SetGwClusterId(v string) *ModifyConsumerRequest
	GetGwClusterId() *string
	SetIsDefault(v string) *ModifyConsumerRequest
	GetIsDefault() *string
	SetName(v string) *ModifyConsumerRequest
	GetName() *string
	SetRegionId(v string) *ModifyConsumerRequest
	GetRegionId() *string
}

type ModifyConsumerRequest struct {
	// example:
	//
	// cg-xxxxxx
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
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
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 0
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumerRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *ModifyConsumerRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ModifyConsumerRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyConsumerRequest) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ModifyConsumerRequest) GetName() *string {
	return s.Name
}

func (s *ModifyConsumerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyConsumerRequest) SetConsumerGroupName(v string) *ModifyConsumerRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *ModifyConsumerRequest) SetConsumerId(v string) *ModifyConsumerRequest {
	s.ConsumerId = &v
	return s
}

func (s *ModifyConsumerRequest) SetGwClusterId(v string) *ModifyConsumerRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyConsumerRequest) SetIsDefault(v string) *ModifyConsumerRequest {
	s.IsDefault = &v
	return s
}

func (s *ModifyConsumerRequest) SetName(v string) *ModifyConsumerRequest {
	s.Name = &v
	return s
}

func (s *ModifyConsumerRequest) SetRegionId(v string) *ModifyConsumerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumerRequest) Validate() error {
	return dara.Validate(s)
}
