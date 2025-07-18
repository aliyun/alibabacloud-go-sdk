// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAcceleratePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationType(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetAccelerationType() *string
	SetDescription(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetDescription() *string
	SetName(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetName() *string
	SetPriority(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetPriority() *string
	SetShowInClient(v int32) *CreateEnterpriseAcceleratePolicyRequest
	GetShowInClient() *int32
	SetUpstreamHost(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetUpstreamHost() *string
	SetUpstreamPort(v int32) *CreateEnterpriseAcceleratePolicyRequest
	GetUpstreamPort() *int32
	SetUpstreamType(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetUpstreamType() *string
	SetUserAttributeGroup(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetUserAttributeGroup() *string
}

type CreateEnterpriseAcceleratePolicyRequest struct {
	// example:
	//
	// whitelist
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 99
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	ShowInClient *int32 `json:"ShowInClient,omitempty" xml:"ShowInClient,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12.34.56.XX
	UpstreamHost *string `json:"UpstreamHost,omitempty" xml:"UpstreamHost,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	UpstreamPort *int32 `json:"UpstreamPort,omitempty" xml:"UpstreamPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ga
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	// This parameter is required.
	UserAttributeGroup *string `json:"UserAttributeGroup,omitempty" xml:"UserAttributeGroup,omitempty"`
}

func (s CreateEnterpriseAcceleratePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAcceleratePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetShowInClient() *int32 {
	return s.ShowInClient
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUpstreamHost() *string {
	return s.UpstreamHost
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUpstreamPort() *int32 {
	return s.UpstreamPort
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUserAttributeGroup() *string {
	return s.UserAttributeGroup
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetAccelerationType(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.AccelerationType = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetDescription(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetName(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetPriority(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetShowInClient(v int32) *CreateEnterpriseAcceleratePolicyRequest {
	s.ShowInClient = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUpstreamHost(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.UpstreamHost = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUpstreamPort(v int32) *CreateEnterpriseAcceleratePolicyRequest {
	s.UpstreamPort = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUpstreamType(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.UpstreamType = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUserAttributeGroup(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.UserAttributeGroup = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) Validate() error {
	return dara.Validate(s)
}
