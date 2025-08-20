// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChainConfig(v string) *UpdateChainRequest
	GetChainConfig() *string
	SetChainId(v string) *UpdateChainRequest
	GetChainId() *string
	SetDescription(v string) *UpdateChainRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateChainRequest
	GetInstanceId() *string
	SetName(v string) *UpdateChainRequest
	GetName() *string
	SetScopeExclude(v []*string) *UpdateChainRequest
	GetScopeExclude() []*string
}

type UpdateChainRequest struct {
	// The configuration of the delivery chain in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// chainconfig
	ChainConfig *string `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty"`
	// The ID of the delivery chain.
	//
	// This parameter is required.
	//
	// example:
	//
	// chi-02ymhtwl3cq8****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The description of the delivery chain.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the delivery chain.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Repositories in which the delivery chain does not take effect.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
}

func (s UpdateChainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChainRequest) GoString() string {
	return s.String()
}

func (s *UpdateChainRequest) GetChainConfig() *string {
	return s.ChainConfig
}

func (s *UpdateChainRequest) GetChainId() *string {
	return s.ChainId
}

func (s *UpdateChainRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateChainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateChainRequest) GetName() *string {
	return s.Name
}

func (s *UpdateChainRequest) GetScopeExclude() []*string {
	return s.ScopeExclude
}

func (s *UpdateChainRequest) SetChainConfig(v string) *UpdateChainRequest {
	s.ChainConfig = &v
	return s
}

func (s *UpdateChainRequest) SetChainId(v string) *UpdateChainRequest {
	s.ChainId = &v
	return s
}

func (s *UpdateChainRequest) SetDescription(v string) *UpdateChainRequest {
	s.Description = &v
	return s
}

func (s *UpdateChainRequest) SetInstanceId(v string) *UpdateChainRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChainRequest) SetName(v string) *UpdateChainRequest {
	s.Name = &v
	return s
}

func (s *UpdateChainRequest) SetScopeExclude(v []*string) *UpdateChainRequest {
	s.ScopeExclude = v
	return s
}

func (s *UpdateChainRequest) Validate() error {
	return dara.Validate(s)
}
