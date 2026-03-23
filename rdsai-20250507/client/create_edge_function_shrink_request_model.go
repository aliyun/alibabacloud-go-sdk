// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeFunctionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEdgeFunctionShrinkRequest
	GetClientToken() *string
	SetCodeShrink(v string) *CreateEdgeFunctionShrinkRequest
	GetCodeShrink() *string
	SetCustomConfigShrink(v string) *CreateEdgeFunctionShrinkRequest
	GetCustomConfigShrink() *string
	SetEdgeFunctionName(v string) *CreateEdgeFunctionShrinkRequest
	GetEdgeFunctionName() *string
	SetEnvsShrink(v string) *CreateEdgeFunctionShrinkRequest
	GetEnvsShrink() *string
	SetInstanceName(v string) *CreateEdgeFunctionShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *CreateEdgeFunctionShrinkRequest
	GetRegionId() *string
}

type CreateEdgeFunctionShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The code parameters.
	CodeShrink *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration parameters of the edge function.
	CustomConfigShrink *string `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// ef-*****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// The environment variables.
	EnvsShrink *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The ID of the RDS Supabase instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEdgeFunctionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeFunctionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeFunctionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEdgeFunctionShrinkRequest) GetCodeShrink() *string {
	return s.CodeShrink
}

func (s *CreateEdgeFunctionShrinkRequest) GetCustomConfigShrink() *string {
	return s.CustomConfigShrink
}

func (s *CreateEdgeFunctionShrinkRequest) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *CreateEdgeFunctionShrinkRequest) GetEnvsShrink() *string {
	return s.EnvsShrink
}

func (s *CreateEdgeFunctionShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateEdgeFunctionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEdgeFunctionShrinkRequest) SetClientToken(v string) *CreateEdgeFunctionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) SetCodeShrink(v string) *CreateEdgeFunctionShrinkRequest {
	s.CodeShrink = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) SetCustomConfigShrink(v string) *CreateEdgeFunctionShrinkRequest {
	s.CustomConfigShrink = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) SetEdgeFunctionName(v string) *CreateEdgeFunctionShrinkRequest {
	s.EdgeFunctionName = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) SetEnvsShrink(v string) *CreateEdgeFunctionShrinkRequest {
	s.EnvsShrink = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) SetInstanceName(v string) *CreateEdgeFunctionShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) SetRegionId(v string) *CreateEdgeFunctionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEdgeFunctionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
