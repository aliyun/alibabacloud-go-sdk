// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDSEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeDSEntityRequest
	GetAgentKey() *string
	SetEntityId(v int64) *DescribeDSEntityRequest
	GetEntityId() *int64
	SetInstanceId(v string) *DescribeDSEntityRequest
	GetInstanceId() *string
}

type DescribeDSEntityRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDSEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDSEntityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDSEntityRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeDSEntityRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeDSEntityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDSEntityRequest) SetAgentKey(v string) *DescribeDSEntityRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDSEntityRequest) SetEntityId(v int64) *DescribeDSEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeDSEntityRequest) SetInstanceId(v string) *DescribeDSEntityRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDSEntityRequest) Validate() error {
	return dara.Validate(s)
}
