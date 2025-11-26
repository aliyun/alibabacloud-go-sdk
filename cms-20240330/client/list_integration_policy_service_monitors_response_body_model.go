// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyServiceMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListIntegrationPolicyServiceMonitorsResponseBody
	GetClusterId() *string
	SetPolicyId(v string) *ListIntegrationPolicyServiceMonitorsResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *ListIntegrationPolicyServiceMonitorsResponseBody
	GetRequestId() *string
	SetServiceMonitors(v []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) *ListIntegrationPolicyServiceMonitorsResponseBody
	GetServiceMonitors() []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors
}

type ListIntegrationPolicyServiceMonitorsResponseBody struct {
	// example:
	//
	// ea119prod-ea119blinkcptssd1
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// policy-ac38a7cb02d14ff48bc9f97d0a75063e
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId       *string                                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceMonitors []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors `json:"serviceMonitors,omitempty" xml:"serviceMonitors,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyServiceMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyServiceMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) GetServiceMonitors() []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	return s.ServiceMonitors
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) SetClusterId(v string) *ListIntegrationPolicyServiceMonitorsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) SetPolicyId(v string) *ListIntegrationPolicyServiceMonitorsResponseBody {
	s.PolicyId = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) SetRequestId(v string) *ListIntegrationPolicyServiceMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) SetServiceMonitors(v []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) *ListIntegrationPolicyServiceMonitorsResponseBody {
	s.ServiceMonitors = v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBody) Validate() error {
	if s.ServiceMonitors != nil {
		for _, item := range s.ServiceMonitors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors struct {
	// example:
	//
	// cloud-acs-ecs
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// example:
	//
	// release-1234567
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// example:
	//
	// 0.0.1
	AddonVersion *string `json:"addonVersion,omitempty" xml:"addonVersion,omitempty"`
	// example:
	//
	// apiVersion: xxxxx
	ConfigYaml *string `json:"configYaml,omitempty" xml:"configYaml,omitempty"`
	// example:
	//
	// run
	EnableStatus *string `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
	// example:
	//
	// YXBpVmVyc2lvbjogeHh4eHgK
	EncryptYaml *bool                                                                       `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	Endpoints   []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MatchedServiceCount *int64 `json:"matchedServiceCount,omitempty" xml:"matchedServiceCount,omitempty"`
	// example:
	//
	// 62a526c5-f6ca-4cfb-b5a4-b76974cffe51
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// arms-prom
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetEndpoints() []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints {
	return s.Endpoints
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetMatchedServiceCount() *int64 {
	return s.MatchedServiceCount
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetAddonName(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetAddonReleaseName(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetAddonVersion(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.AddonVersion = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetConfigYaml(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.ConfigYaml = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetEnableStatus(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.EnableStatus = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetEncryptYaml(v bool) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.EncryptYaml = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetEndpoints(v []*ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.Endpoints = v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetMatchedServiceCount(v int64) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.MatchedServiceCount = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetName(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) SetNamespace(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitors) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints struct {
	// example:
	//
	// 30s
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// 65
	MatchedTargetCount *int64 `json:"matchedTargetCount,omitempty" xml:"matchedTargetCount,omitempty"`
	// example:
	//
	// /metrics
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 9100
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// http
	TargetPort *string `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
}

func (s ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) GetInterval() *string {
	return s.Interval
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) GetMatchedTargetCount() *int64 {
	return s.MatchedTargetCount
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) GetPath() *string {
	return s.Path
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) GetPort() *string {
	return s.Port
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) GetTargetPort() *string {
	return s.TargetPort
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) SetInterval(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints {
	s.Interval = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) SetMatchedTargetCount(v int64) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints {
	s.MatchedTargetCount = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) SetPath(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints {
	s.Path = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) SetPort(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints {
	s.Port = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) SetTargetPort(v string) *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints {
	s.TargetPort = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponseBodyServiceMonitorsEndpoints) Validate() error {
	return dara.Validate(s)
}
