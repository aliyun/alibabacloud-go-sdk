// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *AuthDiagnosisRequest
	GetXDebugId() *string
	SetAutoCreateRole(v bool) *AuthDiagnosisRequest
	GetAutoCreateRole() *bool
	SetAutoInstallAgent(v bool) *AuthDiagnosisRequest
	GetAutoInstallAgent() *bool
	SetInstances(v []*AuthDiagnosisRequestInstances) *AuthDiagnosisRequest
	GetInstances() []*AuthDiagnosisRequestInstances
	SetXSysomInvokeSource(v string) *AuthDiagnosisRequest
	GetXSysomInvokeSource() *string
}

type AuthDiagnosisRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// Specifies whether to enable automatic creation of the service-linked role.
	AutoCreateRole *bool `json:"autoCreateRole,omitempty" xml:"autoCreateRole,omitempty"`
	// Specifies whether to automatically install the latest version of the agent if it is not installed.
	AutoInstallAgent *bool `json:"autoInstallAgent,omitempty" xml:"autoInstallAgent,omitempty"`
	// The list of instances authorized for diagnosis.
	Instances          []*AuthDiagnosisRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	XSysomInvokeSource *string                          `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s AuthDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *AuthDiagnosisRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *AuthDiagnosisRequest) SetXDebugId(v string) *AuthDiagnosisRequest {
	s.XDebugId = &v
	return s
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

func (s *AuthDiagnosisRequest) SetXSysomInvokeSource(v string) *AuthDiagnosisRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *AuthDiagnosisRequest) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AuthDiagnosisRequestInstances struct {
	// The instance ID.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
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
