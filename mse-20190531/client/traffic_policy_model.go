// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrafficPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerSettings(v *TrafficPolicyLoadBalancerSettings) *TrafficPolicy
	GetLoadBalancerSettings() *TrafficPolicyLoadBalancerSettings
	SetTlsSetting(v *TrafficPolicyTlsSetting) *TrafficPolicy
	GetTlsSetting() *TrafficPolicyTlsSetting
}

type TrafficPolicy struct {
	// The load balancing settings.
	LoadBalancerSettings *TrafficPolicyLoadBalancerSettings `json:"LoadBalancerSettings,omitempty" xml:"LoadBalancerSettings,omitempty" type:"Struct"`
	// The data structure.
	TlsSetting *TrafficPolicyTlsSetting `json:"TlsSetting,omitempty" xml:"TlsSetting,omitempty" type:"Struct"`
}

func (s TrafficPolicy) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicy) GoString() string {
	return s.String()
}

func (s *TrafficPolicy) GetLoadBalancerSettings() *TrafficPolicyLoadBalancerSettings {
	return s.LoadBalancerSettings
}

func (s *TrafficPolicy) GetTlsSetting() *TrafficPolicyTlsSetting {
	return s.TlsSetting
}

func (s *TrafficPolicy) SetLoadBalancerSettings(v *TrafficPolicyLoadBalancerSettings) *TrafficPolicy {
	s.LoadBalancerSettings = v
	return s
}

func (s *TrafficPolicy) SetTlsSetting(v *TrafficPolicyTlsSetting) *TrafficPolicy {
	s.TlsSetting = v
	return s
}

func (s *TrafficPolicy) Validate() error {
	if s.LoadBalancerSettings != nil {
		if err := s.LoadBalancerSettings.Validate(); err != nil {
			return err
		}
	}
	if s.TlsSetting != nil {
		if err := s.TlsSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrafficPolicyLoadBalancerSettings struct {
	// The data structure.
	ConsistentHashLBConfig *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig `json:"ConsistentHashLBConfig,omitempty" xml:"ConsistentHashLBConfig,omitempty" type:"Struct"`
	// The load balancing type. Valid values:
	//
	// 	- ROUND_ROBIN: round robin
	//
	// 	- LEAST_CONN: least connection load balancing
	//
	// 	- RANDOM: random load balancing
	//
	// 	- CONSISTENT_HASH: consistent hashing load balancing
	//
	// example:
	//
	// RANDOM
	LoadbalancerType *string `json:"LoadbalancerType,omitempty" xml:"LoadbalancerType,omitempty"`
	// The prefetch duration. Unit: seconds.
	WarmupDuration *int64 `json:"WarmupDuration,omitempty" xml:"WarmupDuration,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettings) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettings) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettings) GetConsistentHashLBConfig() *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	return s.ConsistentHashLBConfig
}

func (s *TrafficPolicyLoadBalancerSettings) GetLoadbalancerType() *string {
	return s.LoadbalancerType
}

func (s *TrafficPolicyLoadBalancerSettings) GetWarmupDuration() *int64 {
	return s.WarmupDuration
}

func (s *TrafficPolicyLoadBalancerSettings) SetConsistentHashLBConfig(v *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) *TrafficPolicyLoadBalancerSettings {
	s.ConsistentHashLBConfig = v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) SetLoadbalancerType(v string) *TrafficPolicyLoadBalancerSettings {
	s.LoadbalancerType = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) SetWarmupDuration(v int64) *TrafficPolicyLoadBalancerSettings {
	s.WarmupDuration = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) Validate() error {
	if s.ConsistentHashLBConfig != nil {
		if err := s.ConsistentHashLBConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig struct {
	// The type of the object based on which consistent hashing is performed. Valid values:
	//
	// 	- QUERY_PARAMETER: request parameter
	//
	// 	- COOKIE: cookie
	//
	// 	- SOURCE_IP: the source IP address
	//
	// 	- HEADER: request header
	//
	// example:
	//
	// QUERY_PARAMETER
	ConsistentHashLBType *string `json:"ConsistentHashLBType,omitempty" xml:"ConsistentHashLBType,omitempty"`
	// You must specify this parameter only if ConsistentHashLBType is set to COOKIE.
	HttpCookie *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie `json:"HttpCookie,omitempty" xml:"HttpCookie,omitempty" type:"Struct"`
	// The name of the object based on which consistent hashing is performed. If consistent hashing is performed based on a parameter, set the value to the parameter name. If consistent hashing is performed based on a header, set the value to the header name.
	//
	// example:
	//
	// test
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetConsistentHashLBType() *string {
	return s.ConsistentHashLBType
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetHttpCookie() *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	return s.HttpCookie
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetParameterName() *string {
	return s.ParameterName
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetConsistentHashLBType(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ConsistentHashLBType = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetHttpCookie(v *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.HttpCookie = v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetParameterName(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ParameterName = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) Validate() error {
	if s.HttpCookie != nil {
		if err := s.HttpCookie.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie struct {
	// The name of the cookie.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the cookie.
	//
	// example:
	//
	// /path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The lifecycle of the cookie.
	//
	// example:
	//
	// 10s
	TTL *string `json:"TTL,omitempty" xml:"TTL,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetName() *string {
	return s.Name
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetPath() *string {
	return s.Path
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetTTL() *string {
	return s.TTL
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetName(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Name = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetPath(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Path = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetTTL(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.TTL = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) Validate() error {
	return dara.Validate(s)
}

type TrafficPolicyTlsSetting struct {
	// The trusted CA certificate chain. In mTLS, if the backend service certificate is issued by a private certificate authority (CA), you must add its CA certificate to the trusted CA certificate chain.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIH0zCCBbugAwIBAgIIXsO3pkN/pOAwDQYJKoZIhvcNAQEFBQAwQjESMBAGA1UE AwwJQUNDVlJBSVoxMRAwDgYDVQQLDAdQS0lBQ0NWMQ0wCwYDVQQKDARBQ0NWMQsw CQYDVQQGEwJFUzAeFw0xMTA1MDUwOTM3MzdaFw0zMDEyMzEwOTM3MzdaMEIxEjAQ BgNVBAMMCUFDQ1ZSQUlaMTEQMA4GA1UECwwHUEtJQUNDVjENMAsGA1UECgwEQUND VjELMAkGA1UEBhMCRVMwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCb qau/YUqXry+XZpp0X9DZlv3P4uRm7x8fRzPCRKPfmt4ftVTdFXxpNRFvu8gMjmoY HtiP2Ra8EEg2XPBjs5BaXCQ316PWywlxufEBcoSwfdtNgM3802/J+Nq2DoLSRYWo G2ioPej0RGy9ocLLA76MPhMAhN9KSMDjIgro6TenGEyxCQ0jVn8ETdkXhBilyNpA lHPrzg5XPAOBOp0KoVdDaaxXbXmQeOW1tDvYvEyNKKGno6e6Ak4l0Squ7a4DIrhr IA8wKFSVf+DuzgpmndFALW4ir50awQUZ0m/A8p/4e7MCQvtQqR0tkw8jq8bBD5L/ 0KIV9VMJcRz/RROE5iZe+OCIHAr8Fraocwa48GOEAqDGWuzndN9wrqODJerWx5eH k6fGioozl2A3ED6XPm4pFdahD9GILBKfb6qkxkLrQaLjlUPTAYVtjrs78yM2x/47 4KElB0iryYl0/wiPgL/AlmXz7uxLaL2diMMxs0Dx6M/2OLuc5NF/1OVYm3z61PMO m3WR5LpSLhl+0fXNWhn8ugb2+1KoS5kE3fj5tItQo05iifCHJPqDQsGH+tUtKSpa cXpkatcnYGMN285J9Y0fkIkyF/hzQ7jSWpOGYdbhdQrqeWZ2iE9x6wQl1gpaepPl uUsXQA+xtrn13k/c4LOsOxFwYIRKQ26ZIMApcQrAZQIDAQABo4ICyzCCAscwfQYI KwYBBQUHAQEEcTBvMEwGCCsGAQUFBzAChkBodHRwOi8vd3d3LmFjY3YuZXMvZmls ZWFkbWluL0FyY2hpdm9zL2NlcnRpZmljYWRvcy9yYWl6YWNjdjEuY3J0MB8GCCsG AQUFBzABhhNodHRwOi8vb2NzcC5hY2N2LmVzMB0GA1UdDgQWBBTSh7Tj3zcnk1X2 VuqB5TbMjB4/vTAPBgNVHRMBAf8EBTADAQH/MB8GA1UdIwQYMBaAFNKHtOPfNyeT VfZW6oHlNsyMHj+9MIIBcwYDVR0gBIIBajCCAWYwggFiBgRVHSAAMIIBWDCCASIG CCsGAQUFBwICMIIBFB6CARAAQQB1AHQAbwByAGkAZABhAGQAIABkAGUAIABDAGUA cgB0AGkAZgBpAGMAYQBjAGkA8wBuACAAUgBhAO0AegAgAGQAZQAgAGwAYQAgAEEA QwBDAFYAIAAoAEEAZwBlAG4AYwBpAGEAIABkAGUAIABUAGUAYwBuAG8AbABvAGcA 7QBhACAAeQAgAEMAZQByAHQAaQBmAGkAYwBhAGMAaQDzAG4AIABFAGwAZQBjAHQA cgDzAG4AaQBjAGEALAAgAEMASQBGACAAUQA0ADYAMAAxADEANQA2AEUAKQAuACAA QwBQAFMAIABlAG4AIABoAHQAdABwADoALwAvAHcAdwB3AC4AYQBjAGMAdgAuAGUA czAwBggrBgEFBQcCARYkaHR0cDovL3d3dy5hY2N2LmVzL2xlZ2lzbGFjaW9uX2Mu aHRtMFUGA1UdHwROMEwwSqBIoEaGRGh0dHA6Ly93d3cuYWNjdi5lcy9maWxlYWRt aW4vQXJjaGl2b3MvY2VydGlmaWNhZG9zL3JhaXphY2N2MV9kZXIuY3JsMA4GA1Ud DwEB/wQEAwIBBjAXBgNVHREEEDAOgQxhY2N2QGFjY3YuZXMwDQYJKoZIhvcNAQEF BQADggIBAJcxAp/n/UNnSEQU5CmH7UwoZtCPNdpNYbdKl02125DgBS4OxnnQ8pdp D70ER9m+27Up2pvZrqmZ1dM8MJP1jaGo/AaNRPTKFpV8M9xii6g3+CfYCS0b78gU JyCpZET/LtZ1qmxNYEAZSUNUY9rizLpm5U9EelvZaoErQNV/+QEnWCzI7UiRfD+m AM/EKXMRNt6GGT6d7hmKG9Ww7Y49nCrADdg9ZuM8Db3VlFzi4qc1GwQA9j9ajepD vV+JHanBsMyZ4k0ACtrJJ1vnE5Bc5PUzolVt3OAJTS+xJlsndQAJxGJ3KQhfnlms tn6tn1QwIgPBHnFk/vk4CpYY3QIUrCPLBhwepH2NDd4nQeit2hW3sCPdK6jT2iWH 7ehVRE2I9DZ+hJp4rPcOVkkO1jMl1oRQQmwgEh0q1b688nCBpHBgvgW1m54ERL5h I6zppSSMEYCUWqKiuUnSwdzRp+0xESyeGabu4VXhwOrPDYTkF7eifKXeVSUG7szA h1xA2syVP1XgNce4hL60Xc16gwFy7ofmXx2utYXGJt/mwZrpHgJHnyqobalbz+xF d3+YJ5oyXSrjhO7FmGYvliAd3djDJ9ew+f7Zfc3Qn48LFFhRny+Lwzgt3uiP1o2H pPVWQxaZLPSkVrQ0uGE3ycJYgBugl6H8WY3pEfbRD0tVNEYqi4Y7 -----END CERTIFICATE-----
	CaCertContent *string `json:"CaCertContent,omitempty" xml:"CaCertContent,omitempty"`
	// The ID of the certificate that is managed in Alibaba Cloud Security.
	//
	// example:
	//
	// 6456988-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The server name indication (SNI) that is used to establish TLS links.
	//
	// example:
	//
	// www.aliyun.com
	Sni *string `json:"Sni,omitempty" xml:"Sni,omitempty"`
	// The Transport Layer Security (TLS) mode that is used to distribute traffic to backend services. Valid values:
	//
	// 	- DISABLE: TLS is disabled. Plaintext is used.
	//
	// 	- SIMPLE: TLS is enabled.
	//
	// 	- MUTUAL: Mutual Transport Layer Security (mTLS) is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// SIMPLE
	TlsMode *string `json:"TlsMode,omitempty" xml:"TlsMode,omitempty"`
}

func (s TrafficPolicyTlsSetting) String() string {
	return dara.Prettify(s)
}

func (s TrafficPolicyTlsSetting) GoString() string {
	return s.String()
}

func (s *TrafficPolicyTlsSetting) GetCaCertContent() *string {
	return s.CaCertContent
}

func (s *TrafficPolicyTlsSetting) GetCertId() *string {
	return s.CertId
}

func (s *TrafficPolicyTlsSetting) GetSni() *string {
	return s.Sni
}

func (s *TrafficPolicyTlsSetting) GetTlsMode() *string {
	return s.TlsMode
}

func (s *TrafficPolicyTlsSetting) SetCaCertContent(v string) *TrafficPolicyTlsSetting {
	s.CaCertContent = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetCertId(v string) *TrafficPolicyTlsSetting {
	s.CertId = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetSni(v string) *TrafficPolicyTlsSetting {
	s.Sni = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetTlsMode(v string) *TrafficPolicyTlsSetting {
	s.TlsMode = &v
	return s
}

func (s *TrafficPolicyTlsSetting) Validate() error {
	return dara.Validate(s)
}
