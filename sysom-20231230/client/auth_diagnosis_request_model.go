// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateRole(v bool) *AuthDiagnosisRequest
	GetAutoCreateRole() *bool
	SetAutoInstallAgent(v bool) *AuthDiagnosisRequest
	GetAutoInstallAgent() *bool
	SetInstances(v []*AuthDiagnosisRequestInstances) *AuthDiagnosisRequest
	GetInstances() []*AuthDiagnosisRequestInstances
}

type AuthDiagnosisRequest struct {
	AutoCreateRole   *bool                            `json:"autoCreateRole,omitempty" xml:"autoCreateRole,omitempty"`
	AutoInstallAgent *bool                            `json:"autoInstallAgent,omitempty" xml:"autoInstallAgent,omitempty"`
	Instances        []*AuthDiagnosisRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s AuthDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequest) GetAutoCreateRole() *bool {
	return s.AutoCreateRole
}

func (s *AuthDiagnosisRequest) GetAutoInstallAgent() *bool {
	return s.AutoInstallAgent
}

func (s *AuthDiagnosisRequest) GetInstances() []*AuthDiagnosisRequestInstances {
	return s.Instances
}

func (s *AuthDiagnosisRequest) SetAutoCreateRole(v bool) *AuthDiagnosisRequest {
	s.AutoCreateRole = &v
	return s
}

func (s *AuthDiagnosisRequest) SetAutoInstallAgent(v bool) *AuthDiagnosisRequest {
	s.AutoInstallAgent = &v
	return s
}

func (s *AuthDiagnosisRequest) SetInstances(v []*AuthDiagnosisRequestInstances) *AuthDiagnosisRequest {
	s.Instances = v
	return s
}

func (s *AuthDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}

type AuthDiagnosisRequestInstances struct {
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s AuthDiagnosisRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s AuthDiagnosisRequestInstances) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequestInstances) GetInstance() *string {
	return s.Instance
}

func (s *AuthDiagnosisRequestInstances) GetRegion() *string {
	return s.Region
}

func (s *AuthDiagnosisRequestInstances) SetInstance(v string) *AuthDiagnosisRequestInstances {
	s.Instance = &v
	return s
}

func (s *AuthDiagnosisRequestInstances) SetRegion(v string) *AuthDiagnosisRequestInstances {
	s.Region = &v
	return s
}

func (s *AuthDiagnosisRequestInstances) Validate() error {
	return dara.Validate(s)
}
