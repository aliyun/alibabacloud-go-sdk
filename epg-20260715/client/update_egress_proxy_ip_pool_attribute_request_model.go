// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyIpPoolAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateEgressProxyIpPoolAttributeRequest
	GetDryRun() *bool
	SetEgressProxyIpPoolId(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetEgressProxyIpPoolId() *string
	SetEgressProxyIpPoolName(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetEgressProxyIpPoolName() *string
	SetHashFactor(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetHashFactor() *string
	SetHashHeaderName(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetHashHeaderName() *string
	SetScheduler(v string) *UpdateEgressProxyIpPoolAttributeRequest
	GetScheduler() *string
}

type UpdateEgressProxyIpPoolAttributeRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// True
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eppool-gw80606a59ecc101****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyIpPoolName *string `json:"EgressProxyIpPoolName,omitempty" xml:"EgressProxyIpPoolName,omitempty"`
	// example:
	//
	// SourceIP
	HashFactor *string `json:"HashFactor,omitempty" xml:"HashFactor,omitempty"`
	// example:
	//
	// X-Forwarded-For
	HashHeaderName *string `json:"HashHeaderName,omitempty" xml:"HashHeaderName,omitempty"`
	// example:
	//
	// Random
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s UpdateEgressProxyIpPoolAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyIpPoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetEgressProxyIpPoolName() *string {
	return s.EgressProxyIpPoolName
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetHashFactor() *string {
	return s.HashFactor
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetHashHeaderName() *string {
	return s.HashHeaderName
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetClientToken(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetDescription(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetDryRun(v bool) *UpdateEgressProxyIpPoolAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetEgressProxyIpPoolId(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetEgressProxyIpPoolName(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.EgressProxyIpPoolName = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetHashFactor(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.HashFactor = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetHashHeaderName(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.HashHeaderName = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) SetScheduler(v string) *UpdateEgressProxyIpPoolAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeRequest) Validate() error {
	return dara.Validate(s)
}
