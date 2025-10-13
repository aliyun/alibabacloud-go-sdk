// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyPodMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListIntegrationPolicyPodMonitorsResponseBody
	GetClusterId() *string
	SetPodMonitors(v []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) *ListIntegrationPolicyPodMonitorsResponseBody
	GetPodMonitors() []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors
	SetPolicyId(v string) *ListIntegrationPolicyPodMonitorsResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *ListIntegrationPolicyPodMonitorsResponseBody
	GetRequestId() *string
}

type ListIntegrationPolicyPodMonitorsResponseBody struct {
	ClusterId   *string                                                    `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	PodMonitors []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors `json:"podMonitors,omitempty" xml:"podMonitors,omitempty" type:"Repeated"`
	// example:
	//
	// policy-c9efed2b99c348d49e589c5f780fc074
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIntegrationPolicyPodMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyPodMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) GetPodMonitors() []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	return s.PodMonitors
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) SetClusterId(v string) *ListIntegrationPolicyPodMonitorsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) SetPodMonitors(v []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) *ListIntegrationPolicyPodMonitorsResponseBody {
	s.PodMonitors = v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) SetPolicyId(v string) *ListIntegrationPolicyPodMonitorsResponseBody {
	s.PolicyId = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) SetRequestId(v string) *ListIntegrationPolicyPodMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBody) Validate() error {
	if s.PodMonitors != nil {
		for _, item := range s.PodMonitors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors struct {
	// example:
	//
	// cloud-ecs
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// example:
	//
	// release-2345678
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
	EncryptYaml *bool                                                               `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	Endpoints   []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	MatchedPodCount *int64 `json:"matchedPodCount,omitempty" xml:"matchedPodCount,omitempty"`
	// example:
	//
	// znzmo_entity_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// sla-ns-d5aeb2b4f91b47
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetEndpoints() []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints {
	return s.Endpoints
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetMatchedPodCount() *int64 {
	return s.MatchedPodCount
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetAddonName(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetAddonReleaseName(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetAddonVersion(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.AddonVersion = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetConfigYaml(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.ConfigYaml = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetEnableStatus(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.EnableStatus = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetEncryptYaml(v bool) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.EncryptYaml = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetEndpoints(v []*ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.Endpoints = v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetMatchedPodCount(v int64) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.MatchedPodCount = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetName(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) SetNamespace(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitors) Validate() error {
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

type ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints struct {
	// example:
	//
	// 30s
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// 1
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
	// https
	TargetPort *string `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
}

func (s ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) GetInterval() *string {
	return s.Interval
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) GetMatchedTargetCount() *int64 {
	return s.MatchedTargetCount
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) GetPath() *string {
	return s.Path
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) GetPort() *string {
	return s.Port
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) GetTargetPort() *string {
	return s.TargetPort
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) SetInterval(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints {
	s.Interval = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) SetMatchedTargetCount(v int64) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints {
	s.MatchedTargetCount = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) SetPath(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints {
	s.Path = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) SetPort(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints {
	s.Port = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) SetTargetPort(v string) *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints {
	s.TargetPort = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponseBodyPodMonitorsEndpoints) Validate() error {
	return dara.Validate(s)
}
