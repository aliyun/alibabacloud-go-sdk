// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPlaneLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliuid(v string) *UpdateControlPlaneLogRequest
	GetAliuid() *string
	SetComponents(v []*string) *UpdateControlPlaneLogRequest
	GetComponents() []*string
	SetLogProject(v string) *UpdateControlPlaneLogRequest
	GetLogProject() *string
	SetLogTtl(v string) *UpdateControlPlaneLogRequest
	GetLogTtl() *string
}

type UpdateControlPlaneLogRequest struct {
	// The Alibaba Cloud account ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 162981*****
	Aliuid *string `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// The list of components for which control plane logging is enabled.
	Components []*string `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The name of the SLS project used to store control plane component logs.
	//
	// Default value: k8s-log-$Cluster ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// k8s-log-c5b5e80b0b64a4bf6939d2d8fbbc5****
	LogProject *string `json:"log_project,omitempty" xml:"log_project,omitempty"`
	// The retention period of logs in the SLS Logstore. Valid values: 1 to 3000. Unit: days.
	//
	// Default value: 30 days.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	LogTtl *string `json:"log_ttl,omitempty" xml:"log_ttl,omitempty"`
}

func (s UpdateControlPlaneLogRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPlaneLogRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogRequest) GetAliuid() *string {
	return s.Aliuid
}

func (s *UpdateControlPlaneLogRequest) GetComponents() []*string {
	return s.Components
}

func (s *UpdateControlPlaneLogRequest) GetLogProject() *string {
	return s.LogProject
}

func (s *UpdateControlPlaneLogRequest) GetLogTtl() *string {
	return s.LogTtl
}

func (s *UpdateControlPlaneLogRequest) SetAliuid(v string) *UpdateControlPlaneLogRequest {
	s.Aliuid = &v
	return s
}

func (s *UpdateControlPlaneLogRequest) SetComponents(v []*string) *UpdateControlPlaneLogRequest {
	s.Components = v
	return s
}

func (s *UpdateControlPlaneLogRequest) SetLogProject(v string) *UpdateControlPlaneLogRequest {
	s.LogProject = &v
	return s
}

func (s *UpdateControlPlaneLogRequest) SetLogTtl(v string) *UpdateControlPlaneLogRequest {
	s.LogTtl = &v
	return s
}

func (s *UpdateControlPlaneLogRequest) Validate() error {
	return dara.Validate(s)
}
