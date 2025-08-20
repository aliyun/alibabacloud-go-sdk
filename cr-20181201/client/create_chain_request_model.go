// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChainConfig(v string) *CreateChainRequest
	GetChainConfig() *string
	SetDescription(v string) *CreateChainRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateChainRequest
	GetInstanceId() *string
	SetName(v string) *CreateChainRequest
	GetName() *string
	SetRepoName(v string) *CreateChainRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *CreateChainRequest
	GetRepoNamespaceName() *string
	SetScopeExclude(v []*string) *CreateChainRequest
	GetScopeExclude() []*string
}

type CreateChainRequest struct {
	// The configuration of the delivery chain in the JSON format.
	//
	// example:
	//
	// chainconfig
	ChainConfig *string `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty"`
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
	// The name of the repository.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// Repositories in which the delivery chain does not take effect.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
}

func (s CreateChainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChainRequest) GoString() string {
	return s.String()
}

func (s *CreateChainRequest) GetChainConfig() *string {
	return s.ChainConfig
}

func (s *CreateChainRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateChainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateChainRequest) GetName() *string {
	return s.Name
}

func (s *CreateChainRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateChainRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *CreateChainRequest) GetScopeExclude() []*string {
	return s.ScopeExclude
}

func (s *CreateChainRequest) SetChainConfig(v string) *CreateChainRequest {
	s.ChainConfig = &v
	return s
}

func (s *CreateChainRequest) SetDescription(v string) *CreateChainRequest {
	s.Description = &v
	return s
}

func (s *CreateChainRequest) SetInstanceId(v string) *CreateChainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChainRequest) SetName(v string) *CreateChainRequest {
	s.Name = &v
	return s
}

func (s *CreateChainRequest) SetRepoName(v string) *CreateChainRequest {
	s.RepoName = &v
	return s
}

func (s *CreateChainRequest) SetRepoNamespaceName(v string) *CreateChainRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *CreateChainRequest) SetScopeExclude(v []*string) *CreateChainRequest {
	s.ScopeExclude = v
	return s
}

func (s *CreateChainRequest) Validate() error {
	return dara.Validate(s)
}
