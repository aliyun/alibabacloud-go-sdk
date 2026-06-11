// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyCustomScrapeJobRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody
	GetClusterId() *string
	SetCustomScrapeJobRules(v []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody
	GetCustomScrapeJobRules() []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules
	SetPolicyId(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody
	GetRequestId() *string
}

type ListIntegrationPolicyCustomScrapeJobRulesResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// et15prod-et15storage
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The custom scrape job rules.
	CustomScrapeJobRules []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules `json:"customScrapeJobRules,omitempty" xml:"customScrapeJobRules,omitempty" type:"Repeated"`
	// The policy ID.
	//
	// example:
	//
	// policy-15abcc24c06f4797832b5954198e1ed1
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) GetCustomScrapeJobRules() []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	return s.CustomScrapeJobRules
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) SetClusterId(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) SetCustomScrapeJobRules(v []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody {
	s.CustomScrapeJobRules = v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) SetPolicyId(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody {
	s.PolicyId = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) SetRequestId(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) Validate() error {
	if s.CustomScrapeJobRules != nil {
		for _, item := range s.CustomScrapeJobRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules struct {
	// The add-on name.
	//
	// example:
	//
	// cloud-ecs
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// The name of the add-on release.
	//
	// example:
	//
	// release-12345678
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// The add-on version.
	//
	// example:
	//
	// 0.0.1
	AddonVersion *string `json:"addonVersion,omitempty" xml:"addonVersion,omitempty"`
	// The configuration YAML file.
	//
	// example:
	//
	// scrape_cofnigs:
	//
	// - jobxxxxxx
	ConfigYaml *string `json:"configYaml,omitempty" xml:"configYaml,omitempty"`
	// The enabled status.
	//
	// example:
	//
	// mini
	EnableStatus *string `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
	// Indicates whether the YAML file is encrypted.
	//
	// example:
	//
	// true
	EncryptYaml *bool `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	// The number of matched pods.
	//
	// example:
	//
	// 1
	MatchedPodCount *int64 `json:"matchedPodCount,omitempty" xml:"matchedPodCount,omitempty"`
	// The details.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The service name.
	//
	// example:
	//
	// dlab1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// prod-data
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The custom configurations.
	ScrapeConfigs []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs `json:"scrapeConfigs,omitempty" xml:"scrapeConfigs,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetMatchedPodCount() *int64 {
	return s.MatchedPodCount
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetMessage() *string {
	return s.Message
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) GetScrapeConfigs() []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	return s.ScrapeConfigs
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetAddonName(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetAddonReleaseName(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetAddonVersion(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.AddonVersion = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetConfigYaml(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.ConfigYaml = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetEnableStatus(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.EnableStatus = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetEncryptYaml(v bool) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.EncryptYaml = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetMatchedPodCount(v int64) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.MatchedPodCount = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetMessage(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.Message = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetName(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetNamespace(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) SetScrapeConfigs(v []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules {
	s.ScrapeConfigs = v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules) Validate() error {
	if s.ScrapeConfigs != nil {
		for _, item := range s.ScrapeConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs struct {
	// The scrape job name.
	//
	// example:
	//
	// mysql-exporter
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The details.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The scrape path.
	//
	// example:
	//
	// /metrics
	MetricsPath *string `json:"metricsPath,omitempty" xml:"metricsPath,omitempty"`
	// The invocation method.
	//
	// example:
	//
	// http
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// The scrape interval.
	//
	// example:
	//
	// 30s
	ScrapeInterval *string `json:"scrapeInterval,omitempty" xml:"scrapeInterval,omitempty"`
	// The scrape timeout period.
	//
	// example:
	//
	// 60s
	ScrapeTimeout *string `json:"scrapeTimeout,omitempty" xml:"scrapeTimeout,omitempty"`
	// The service discovery configurations.
	ServiceDiscoveryConfigs []*string `json:"serviceDiscoveryConfigs,omitempty" xml:"serviceDiscoveryConfigs,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetJobName() *string {
	return s.JobName
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetMessage() *string {
	return s.Message
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetMetricsPath() *string {
	return s.MetricsPath
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetScheme() *string {
	return s.Scheme
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetScrapeInterval() *string {
	return s.ScrapeInterval
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetScrapeTimeout() *string {
	return s.ScrapeTimeout
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) GetServiceDiscoveryConfigs() []*string {
	return s.ServiceDiscoveryConfigs
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetJobName(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.JobName = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetMessage(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.Message = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetMetricsPath(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.MetricsPath = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetScheme(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.Scheme = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetScrapeInterval(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.ScrapeInterval = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetScrapeTimeout(v string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.ScrapeTimeout = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) SetServiceDiscoveryConfigs(v []*string) *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs {
	s.ServiceDiscoveryConfigs = v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs) Validate() error {
	return dara.Validate(s)
}
