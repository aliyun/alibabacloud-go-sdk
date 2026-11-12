// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupName(v string) *CreateConsumerRequest
	GetConsumerGroupName() *string
	SetGwClusterId(v string) *CreateConsumerRequest
	GetGwClusterId() *string
	SetKeyType(v string) *CreateConsumerRequest
	GetKeyType() *string
	SetName(v string) *CreateConsumerRequest
	GetName() *string
	SetNickName(v string) *CreateConsumerRequest
	GetNickName() *string
	SetRegionId(v string) *CreateConsumerRequest
	GetRegionId() *string
}

type CreateConsumerRequest struct {
	// The ID of the user group. You can obtain the ID from the ConsumerGroupId parameter returned by CreateConsumerGroup. This parameter is actually required. If it is not provided, the API returns an InvalidParameters.Format (400) error. Although the parameter is named ConsumerGroupName, you must pass the ConsumerGroupId in the cg-xxxx format. You can obtain the ID from the ConsumerGroupId field returned by CreateConsumerGroup or DescribeConsumerGroups.
	//
	// example:
	//
	// cg-p3gk2oh55c**
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The ID of the gateway instance. You can obtain the ID from the GwClusterId parameter returned by CreateGateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The key type. Default value: ApiKey. Currently, only ApiKey is supported.
	//
	// example:
	//
	// ApiKey
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerRequest) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *CreateConsumerRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateConsumerRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *CreateConsumerRequest) GetName() *string {
	return s.Name
}

func (s *CreateConsumerRequest) GetNickName() *string {
	return s.NickName
}

func (s *CreateConsumerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConsumerRequest) SetConsumerGroupName(v string) *CreateConsumerRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *CreateConsumerRequest) SetGwClusterId(v string) *CreateConsumerRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateConsumerRequest) SetKeyType(v string) *CreateConsumerRequest {
	s.KeyType = &v
	return s
}

func (s *CreateConsumerRequest) SetName(v string) *CreateConsumerRequest {
	s.Name = &v
	return s
}

func (s *CreateConsumerRequest) SetNickName(v string) *CreateConsumerRequest {
	s.NickName = &v
	return s
}

func (s *CreateConsumerRequest) SetRegionId(v string) *CreateConsumerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerRequest) Validate() error {
	return dara.Validate(s)
}
