// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupName(v string) *CreateConsumerGroupRequest
	GetConsumerGroupName() *string
	SetGwClusterId(v string) *CreateConsumerGroupRequest
	GetGwClusterId() *string
	SetIsDefault(v string) *CreateConsumerGroupRequest
	GetIsDefault() *string
	SetNickName(v string) *CreateConsumerGroupRequest
	GetNickName() *string
	SetRegionId(v string) *CreateConsumerGroupRequest
	GetRegionId() *string
}

type CreateConsumerGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 0
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *CreateConsumerGroupRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateConsumerGroupRequest) GetIsDefault() *string {
	return s.IsDefault
}

func (s *CreateConsumerGroupRequest) GetNickName() *string {
	return s.NickName
}

func (s *CreateConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupName(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetGwClusterId(v string) *CreateConsumerGroupRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetIsDefault(v string) *CreateConsumerGroupRequest {
	s.IsDefault = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetNickName(v string) *CreateConsumerGroupRequest {
	s.NickName = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRegionId(v string) *CreateConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}
