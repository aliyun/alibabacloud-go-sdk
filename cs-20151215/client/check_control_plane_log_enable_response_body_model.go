// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckControlPlaneLogEnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliuid(v string) *CheckControlPlaneLogEnableResponseBody
	GetAliuid() *string
	SetComponents(v []*string) *CheckControlPlaneLogEnableResponseBody
	GetComponents() []*string
	SetLogProject(v string) *CheckControlPlaneLogEnableResponseBody
	GetLogProject() *string
	SetLogTtl(v string) *CheckControlPlaneLogEnableResponseBody
	GetLogTtl() *string
}

type CheckControlPlaneLogEnableResponseBody struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 162981*****
	Aliuid *string `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// The list of components for which control plane logging is enabled.
	//
	// This parameter is required.
	Components []*string `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The name of the Simple Log Service (SLS) project that stores the control plane component logs.
	//
	// Default value: k8s-log-$Cluster ID.
	//
	// example:
	//
	// k8s-log-c5b5e80b0b64a4bf6939d2d8fbbc5****
	LogProject *string `json:"log_project,omitempty" xml:"log_project,omitempty"`
	// The data retention period of logs in the SLS Logstore. Unit: days.
	//
	// example:
	//
	// 30
	LogTtl *string `json:"log_ttl,omitempty" xml:"log_ttl,omitempty"`
}

func (s CheckControlPlaneLogEnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckControlPlaneLogEnableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckControlPlaneLogEnableResponseBody) GetAliuid() *string {
	return s.Aliuid
}

func (s *CheckControlPlaneLogEnableResponseBody) GetComponents() []*string {
	return s.Components
}

func (s *CheckControlPlaneLogEnableResponseBody) GetLogProject() *string {
	return s.LogProject
}

func (s *CheckControlPlaneLogEnableResponseBody) GetLogTtl() *string {
	return s.LogTtl
}

func (s *CheckControlPlaneLogEnableResponseBody) SetAliuid(v string) *CheckControlPlaneLogEnableResponseBody {
	s.Aliuid = &v
	return s
}

func (s *CheckControlPlaneLogEnableResponseBody) SetComponents(v []*string) *CheckControlPlaneLogEnableResponseBody {
	s.Components = v
	return s
}

func (s *CheckControlPlaneLogEnableResponseBody) SetLogProject(v string) *CheckControlPlaneLogEnableResponseBody {
	s.LogProject = &v
	return s
}

func (s *CheckControlPlaneLogEnableResponseBody) SetLogTtl(v string) *CheckControlPlaneLogEnableResponseBody {
	s.LogTtl = &v
	return s
}

func (s *CheckControlPlaneLogEnableResponseBody) Validate() error {
	return dara.Validate(s)
}
