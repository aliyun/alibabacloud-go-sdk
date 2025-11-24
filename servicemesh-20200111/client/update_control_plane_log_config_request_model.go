// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPlaneLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateControlPlaneLogConfigRequest
	GetEnabled() *bool
	SetLogTTLInDay(v int32) *UpdateControlPlaneLogConfigRequest
	GetLogTTLInDay() *int32
	SetProject(v string) *UpdateControlPlaneLogConfigRequest
	GetProject() *string
	SetServiceMeshId(v string) *UpdateControlPlaneLogConfigRequest
	GetServiceMeshId() *string
}

type UpdateControlPlaneLogConfigRequest struct {
	// Specifies whether to collect control plane logs to Simple Log Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time to live (TTL) period of the collected logs. Unit: day.
	//
	// example:
	//
	// 30
	LogTTLInDay *int32 `json:"LogTTLInDay,omitempty" xml:"LogTTLInDay,omitempty"`
	// The name of the Simple Log Service project to which control plane logs are collected.
	//
	// example:
	//
	// aia-asm-deva-sh
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c20667db760fe4ee6910220136624****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateControlPlaneLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPlaneLogConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateControlPlaneLogConfigRequest) GetLogTTLInDay() *int32 {
	return s.LogTTLInDay
}

func (s *UpdateControlPlaneLogConfigRequest) GetProject() *string {
	return s.Project
}

func (s *UpdateControlPlaneLogConfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateControlPlaneLogConfigRequest) SetEnabled(v bool) *UpdateControlPlaneLogConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) SetLogTTLInDay(v int32) *UpdateControlPlaneLogConfigRequest {
	s.LogTTLInDay = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) SetProject(v string) *UpdateControlPlaneLogConfigRequest {
	s.Project = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) SetServiceMeshId(v string) *UpdateControlPlaneLogConfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
