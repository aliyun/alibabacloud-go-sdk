// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeFunctionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEdgeFunctionShrinkRequest
	GetClientToken() *string
	SetCodeShrink(v string) *UpdateEdgeFunctionShrinkRequest
	GetCodeShrink() *string
	SetCustomConfigShrink(v string) *UpdateEdgeFunctionShrinkRequest
	GetCustomConfigShrink() *string
	SetEdgeFunctionName(v string) *UpdateEdgeFunctionShrinkRequest
	GetEdgeFunctionName() *string
	SetEnvsShrink(v string) *UpdateEdgeFunctionShrinkRequest
	GetEnvsShrink() *string
	SetInstanceName(v string) *UpdateEdgeFunctionShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *UpdateEdgeFunctionShrinkRequest
	GetRegionId() *string
}

type UpdateEdgeFunctionShrinkRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CodeShrink  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	CustomConfigShrink *string `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// fc-xxxx。
	//
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	EnvsShrink       *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEdgeFunctionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeFunctionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeFunctionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEdgeFunctionShrinkRequest) GetCodeShrink() *string {
	return s.CodeShrink
}

func (s *UpdateEdgeFunctionShrinkRequest) GetCustomConfigShrink() *string {
	return s.CustomConfigShrink
}

func (s *UpdateEdgeFunctionShrinkRequest) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *UpdateEdgeFunctionShrinkRequest) GetEnvsShrink() *string {
	return s.EnvsShrink
}

func (s *UpdateEdgeFunctionShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateEdgeFunctionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEdgeFunctionShrinkRequest) SetClientToken(v string) *UpdateEdgeFunctionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) SetCodeShrink(v string) *UpdateEdgeFunctionShrinkRequest {
	s.CodeShrink = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) SetCustomConfigShrink(v string) *UpdateEdgeFunctionShrinkRequest {
	s.CustomConfigShrink = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) SetEdgeFunctionName(v string) *UpdateEdgeFunctionShrinkRequest {
	s.EdgeFunctionName = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) SetEnvsShrink(v string) *UpdateEdgeFunctionShrinkRequest {
	s.EnvsShrink = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) SetInstanceName(v string) *UpdateEdgeFunctionShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) SetRegionId(v string) *UpdateEdgeFunctionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEdgeFunctionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
