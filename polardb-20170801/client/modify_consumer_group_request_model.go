// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupName(v string) *ModifyConsumerGroupRequest
	GetConsumerGroupName() *string
	SetGwClusterId(v string) *ModifyConsumerGroupRequest
	GetGwClusterId() *string
	SetIsDefault(v string) *ModifyConsumerGroupRequest
	GetIsDefault() *string
	SetNickName(v string) *ModifyConsumerGroupRequest
	GetNickName() *string
	SetRegionId(v string) *ModifyConsumerGroupRequest
	GetRegionId() *string
}

type ModifyConsumerGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cg-xxxxxx
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
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
	// example:
	//
	// yonghu
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *ModifyConsumerGroupRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyConsumerGroupRequest) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ModifyConsumerGroupRequest) GetNickName() *string {
	return s.NickName
}

func (s *ModifyConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyConsumerGroupRequest) SetConsumerGroupName(v string) *ModifyConsumerGroupRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *ModifyConsumerGroupRequest) SetGwClusterId(v string) *ModifyConsumerGroupRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyConsumerGroupRequest) SetIsDefault(v string) *ModifyConsumerGroupRequest {
	s.IsDefault = &v
	return s
}

func (s *ModifyConsumerGroupRequest) SetNickName(v string) *ModifyConsumerGroupRequest {
	s.NickName = &v
	return s
}

func (s *ModifyConsumerGroupRequest) SetRegionId(v string) *ModifyConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
