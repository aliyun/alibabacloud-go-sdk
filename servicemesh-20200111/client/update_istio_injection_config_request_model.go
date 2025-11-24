// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioInjectionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPlaneMode(v string) *UpdateIstioInjectionConfigRequest
	GetDataPlaneMode() *string
	SetEnableIstioInjection(v bool) *UpdateIstioInjectionConfigRequest
	GetEnableIstioInjection() *bool
	SetEnableSidecarSetInjection(v bool) *UpdateIstioInjectionConfigRequest
	GetEnableSidecarSetInjection() *bool
	SetIstioRev(v string) *UpdateIstioInjectionConfigRequest
	GetIstioRev() *string
	SetNamespace(v string) *UpdateIstioInjectionConfigRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *UpdateIstioInjectionConfigRequest
	GetServiceMeshId() *string
}

type UpdateIstioInjectionConfigRequest struct {
	// The data plane mode of the namespace. This parameter is valid only when the Ambient Mesh mode is enabled for the current Service Mesh (ASM) instance. Valid values:
	//
	// 	- ambient: sets the data plane mode of the namespace to the Ambient Mesh mode.
	//
	// 	- sidecar: sets the data plane mode of the namespace to the Sidecar mode.
	//
	// example:
	//
	// ambient
	DataPlaneMode *string `json:"DataPlaneMode,omitempty" xml:"DataPlaneMode,omitempty"`
	// Specifies whether to enable Istio automatic sidecar injection.
	//
	// example:
	//
	// true
	EnableIstioInjection *bool `json:"EnableIstioInjection,omitempty" xml:"EnableIstioInjection,omitempty"`
	// Specifies whether to enable automatic sidecar injection by using SidecarSet.
	//
	// example:
	//
	// false
	EnableSidecarSetInjection *bool `json:"EnableSidecarSetInjection,omitempty" xml:"EnableSidecarSetInjection,omitempty"`
	// Specifies the version to be injected into the namespace. This parameter is valid only when the ASM instance performs a canary release. When IstioRev is not empty, you must not specify EnableIstioInjection and EnableSidecarSetInjection.
	//
	// example:
	//
	// canary
	IstioRev *string `json:"IstioRev,omitempty" xml:"IstioRev,omitempty"`
	// The namespace for which you want to modify the sidecar injection setting.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce2cdbb9d013f447180cf5ca8bb******
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateIstioInjectionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioInjectionConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioInjectionConfigRequest) GetDataPlaneMode() *string {
	return s.DataPlaneMode
}

func (s *UpdateIstioInjectionConfigRequest) GetEnableIstioInjection() *bool {
	return s.EnableIstioInjection
}

func (s *UpdateIstioInjectionConfigRequest) GetEnableSidecarSetInjection() *bool {
	return s.EnableSidecarSetInjection
}

func (s *UpdateIstioInjectionConfigRequest) GetIstioRev() *string {
	return s.IstioRev
}

func (s *UpdateIstioInjectionConfigRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateIstioInjectionConfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateIstioInjectionConfigRequest) SetDataPlaneMode(v string) *UpdateIstioInjectionConfigRequest {
	s.DataPlaneMode = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetEnableIstioInjection(v bool) *UpdateIstioInjectionConfigRequest {
	s.EnableIstioInjection = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetEnableSidecarSetInjection(v bool) *UpdateIstioInjectionConfigRequest {
	s.EnableSidecarSetInjection = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetIstioRev(v string) *UpdateIstioInjectionConfigRequest {
	s.IstioRev = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetNamespace(v string) *UpdateIstioInjectionConfigRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetServiceMeshId(v string) *UpdateIstioInjectionConfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) Validate() error {
	return dara.Validate(s)
}
