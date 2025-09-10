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
	// example:
	//
	// et15prod-et15storage
	ClusterId            *string                                                                      `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	CustomScrapeJobRules []*ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules `json:"customScrapeJobRules,omitempty" xml:"customScrapeJobRules,omitempty" type:"Repeated"`
	// example:
	//
	// policy-15abcc24c06f4797832b5954198e1ed1
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// Id of the request
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
	return dara.Validate(s)
}

type ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRules struct {
	// example:
	//
	// cloud-ecs
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// example:
	//
	// release-12345678
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// example:
	//
	// 0.0.1
	AddonVersion *string `json:"addonVersion,omitempty" xml:"addonVersion,omitempty"`
	// example:
	//
	// scrape_cofnigs:
	//
	// - jobxxxxxx
	ConfigYaml *string `json:"configYaml,omitempty" xml:"configYaml,omitempty"`
	// example:
	//
	// mini
	EnableStatus *string `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
	// example:
	//
	// true
	EncryptYaml *bool `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	// example:
	//
	// 1
	MatchedPodCount *int64 `json:"matchedPodCount,omitempty" xml:"matchedPodCount,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// dlab1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// prod-data
	Namespace     *string                                                                                   `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
	return dara.Validate(s)
}

type ListIntegrationPolicyCustomScrapeJobRulesResponseBodyCustomScrapeJobRulesScrapeConfigs struct {
	// example:
	//
	// mysql-exporter
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// /metrics
	MetricsPath *string `json:"metricsPath,omitempty" xml:"metricsPath,omitempty"`
	// example:
	//
	// http
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// example:
	//
	// 30s
	ScrapeInterval *string `json:"scrapeInterval,omitempty" xml:"scrapeInterval,omitempty"`
	// example:
	//
	// 60s
	ScrapeTimeout           *string   `json:"scrapeTimeout,omitempty" xml:"scrapeTimeout,omitempty"`
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
