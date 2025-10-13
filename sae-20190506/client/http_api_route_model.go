// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiRoute interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *HttpApiRoute
	GetAddressType() *string
	SetDeployStatus(v string) *HttpApiRoute
	GetDeployStatus() *string
	SetDestinationType(v string) *HttpApiRoute
	GetDestinationType() *string
	SetDomains(v []*HttpApiRouteDomains) *HttpApiRoute
	GetDomains() []*HttpApiRouteDomains
	SetEnvironmentId(v string) *HttpApiRoute
	GetEnvironmentId() *string
	SetGatewayId(v string) *HttpApiRoute
	GetGatewayId() *string
	SetHttpApiId(v string) *HttpApiRoute
	GetHttpApiId() *string
	SetHttpApiName(v string) *HttpApiRoute
	GetHttpApiName() *string
	SetHttpApiType(v string) *HttpApiRoute
	GetHttpApiType() *string
	SetIngressId(v int64) *HttpApiRoute
	GetIngressId() *int64
	SetNacosInstanceId(v string) *HttpApiRoute
	GetNacosInstanceId() *string
	SetNacosNamespaceId(v string) *HttpApiRoute
	GetNacosNamespaceId() *string
	SetName(v string) *HttpApiRoute
	GetName() *string
	SetNamespaceId(v string) *HttpApiRoute
	GetNamespaceId() *string
	SetPolicies(v *HttpApiRoutePolicies) *HttpApiRoute
	GetPolicies() *HttpApiRoutePolicies
	SetPredicates(v *HttpApiRoutePredicates) *HttpApiRoute
	GetPredicates() *HttpApiRoutePredicates
	SetRouteId(v string) *HttpApiRoute
	GetRouteId() *string
	SetServices(v []*HttpApiRouteServices) *HttpApiRoute
	GetServices() []*HttpApiRouteServices
	SetSourceType(v string) *HttpApiRoute
	GetSourceType() *string
}

type HttpApiRoute struct {
	// example:
	//
	// intranet/internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// Deploying/NotDeployed/Undeploying/Deployed
	DeployStatus *string `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	// example:
	//
	// Single/Multiple/VersionOriented
	DestinationType *string                `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Domains         []*HttpApiRouteDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	EnvironmentId   *string                `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	GatewayId       *string                `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	HttpApiId       *string                `json:"HttpApiId,omitempty" xml:"HttpApiId,omitempty"`
	// example:
	//
	// Http
	HttpApiName *string `json:"HttpApiName,omitempty" xml:"HttpApiName,omitempty"`
	// example:
	//
	// Http
	HttpApiType *string `json:"HttpApiType,omitempty" xml:"HttpApiType,omitempty"`
	// example:
	//
	// 1
	IngressId       *int64  `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
	NacosInstanceId *string `json:"NacosInstanceId,omitempty" xml:"NacosInstanceId,omitempty"`
	// example:
	//
	// test
	NacosNamespaceId *string                 `json:"NacosNamespaceId,omitempty" xml:"NacosNamespaceId,omitempty"`
	Name             *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId      *string                 `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Policies         *HttpApiRoutePolicies   `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	Predicates       *HttpApiRoutePredicates `json:"Predicates,omitempty" xml:"Predicates,omitempty" type:"Struct"`
	RouteId          *string                 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	Services         []*HttpApiRouteServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// example:
	//
	// SAE_NACOS/SAE_K8S_SERVICE/MSE_NACOS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s HttpApiRoute) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoute) GoString() string {
	return s.String()
}

func (s *HttpApiRoute) GetAddressType() *string {
	return s.AddressType
}

func (s *HttpApiRoute) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *HttpApiRoute) GetDestinationType() *string {
	return s.DestinationType
}

func (s *HttpApiRoute) GetDomains() []*HttpApiRouteDomains {
	return s.Domains
}

func (s *HttpApiRoute) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpApiRoute) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiRoute) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *HttpApiRoute) GetHttpApiName() *string {
	return s.HttpApiName
}

func (s *HttpApiRoute) GetHttpApiType() *string {
	return s.HttpApiType
}

func (s *HttpApiRoute) GetIngressId() *int64 {
	return s.IngressId
}

func (s *HttpApiRoute) GetNacosInstanceId() *string {
	return s.NacosInstanceId
}

func (s *HttpApiRoute) GetNacosNamespaceId() *string {
	return s.NacosNamespaceId
}

func (s *HttpApiRoute) GetName() *string {
	return s.Name
}

func (s *HttpApiRoute) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *HttpApiRoute) GetPolicies() *HttpApiRoutePolicies {
	return s.Policies
}

func (s *HttpApiRoute) GetPredicates() *HttpApiRoutePredicates {
	return s.Predicates
}

func (s *HttpApiRoute) GetRouteId() *string {
	return s.RouteId
}

func (s *HttpApiRoute) GetServices() []*HttpApiRouteServices {
	return s.Services
}

func (s *HttpApiRoute) GetSourceType() *string {
	return s.SourceType
}

func (s *HttpApiRoute) SetAddressType(v string) *HttpApiRoute {
	s.AddressType = &v
	return s
}

func (s *HttpApiRoute) SetDeployStatus(v string) *HttpApiRoute {
	s.DeployStatus = &v
	return s
}

func (s *HttpApiRoute) SetDestinationType(v string) *HttpApiRoute {
	s.DestinationType = &v
	return s
}

func (s *HttpApiRoute) SetDomains(v []*HttpApiRouteDomains) *HttpApiRoute {
	s.Domains = v
	return s
}

func (s *HttpApiRoute) SetEnvironmentId(v string) *HttpApiRoute {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiRoute) SetGatewayId(v string) *HttpApiRoute {
	s.GatewayId = &v
	return s
}

func (s *HttpApiRoute) SetHttpApiId(v string) *HttpApiRoute {
	s.HttpApiId = &v
	return s
}

func (s *HttpApiRoute) SetHttpApiName(v string) *HttpApiRoute {
	s.HttpApiName = &v
	return s
}

func (s *HttpApiRoute) SetHttpApiType(v string) *HttpApiRoute {
	s.HttpApiType = &v
	return s
}

func (s *HttpApiRoute) SetIngressId(v int64) *HttpApiRoute {
	s.IngressId = &v
	return s
}

func (s *HttpApiRoute) SetNacosInstanceId(v string) *HttpApiRoute {
	s.NacosInstanceId = &v
	return s
}

func (s *HttpApiRoute) SetNacosNamespaceId(v string) *HttpApiRoute {
	s.NacosNamespaceId = &v
	return s
}

func (s *HttpApiRoute) SetName(v string) *HttpApiRoute {
	s.Name = &v
	return s
}

func (s *HttpApiRoute) SetNamespaceId(v string) *HttpApiRoute {
	s.NamespaceId = &v
	return s
}

func (s *HttpApiRoute) SetPolicies(v *HttpApiRoutePolicies) *HttpApiRoute {
	s.Policies = v
	return s
}

func (s *HttpApiRoute) SetPredicates(v *HttpApiRoutePredicates) *HttpApiRoute {
	s.Predicates = v
	return s
}

func (s *HttpApiRoute) SetRouteId(v string) *HttpApiRoute {
	s.RouteId = &v
	return s
}

func (s *HttpApiRoute) SetServices(v []*HttpApiRouteServices) *HttpApiRoute {
	s.Services = v
	return s
}

func (s *HttpApiRoute) SetSourceType(v string) *HttpApiRoute {
	s.SourceType = &v
	return s
}

func (s *HttpApiRoute) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	if s.Predicates != nil {
		if err := s.Predicates.Validate(); err != nil {
			return err
		}
	}
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiRouteDomains struct {
	DomainId   *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s HttpApiRouteDomains) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRouteDomains) GoString() string {
	return s.String()
}

func (s *HttpApiRouteDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpApiRouteDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *HttpApiRouteDomains) SetDomainId(v string) *HttpApiRouteDomains {
	s.DomainId = &v
	return s
}

func (s *HttpApiRouteDomains) SetDomainName(v string) *HttpApiRouteDomains {
	s.DomainName = &v
	return s
}

func (s *HttpApiRouteDomains) Validate() error {
	return dara.Validate(s)
}

type HttpApiRoutePolicies struct {
	Fallback *HttpApiRoutePoliciesFallback `json:"Fallback,omitempty" xml:"Fallback,omitempty" type:"Struct"`
	Retry    *HttpApiRoutePoliciesRetry    `json:"Retry,omitempty" xml:"Retry,omitempty" type:"Struct"`
	Timeout  *HttpApiRoutePoliciesTimeout  `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
}

func (s HttpApiRoutePolicies) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePolicies) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePolicies) GetFallback() *HttpApiRoutePoliciesFallback {
	return s.Fallback
}

func (s *HttpApiRoutePolicies) GetRetry() *HttpApiRoutePoliciesRetry {
	return s.Retry
}

func (s *HttpApiRoutePolicies) GetTimeout() *HttpApiRoutePoliciesTimeout {
	return s.Timeout
}

func (s *HttpApiRoutePolicies) SetFallback(v *HttpApiRoutePoliciesFallback) *HttpApiRoutePolicies {
	s.Fallback = v
	return s
}

func (s *HttpApiRoutePolicies) SetRetry(v *HttpApiRoutePoliciesRetry) *HttpApiRoutePolicies {
	s.Retry = v
	return s
}

func (s *HttpApiRoutePolicies) SetTimeout(v *HttpApiRoutePoliciesTimeout) *HttpApiRoutePolicies {
	s.Timeout = v
	return s
}

func (s *HttpApiRoutePolicies) Validate() error {
	if s.Fallback != nil {
		if err := s.Fallback.Validate(); err != nil {
			return err
		}
	}
	if s.Retry != nil {
		if err := s.Retry.Validate(); err != nil {
			return err
		}
	}
	if s.Timeout != nil {
		if err := s.Timeout.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiRoutePoliciesFallback struct {
	Destinations []*HttpApiRoutePoliciesFallbackDestinations `json:"Destinations,omitempty" xml:"Destinations,omitempty" type:"Repeated"`
	Enable       *bool                                       `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s HttpApiRoutePoliciesFallback) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePoliciesFallback) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePoliciesFallback) GetDestinations() []*HttpApiRoutePoliciesFallbackDestinations {
	return s.Destinations
}

func (s *HttpApiRoutePoliciesFallback) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiRoutePoliciesFallback) SetDestinations(v []*HttpApiRoutePoliciesFallbackDestinations) *HttpApiRoutePoliciesFallback {
	s.Destinations = v
	return s
}

func (s *HttpApiRoutePoliciesFallback) SetEnable(v bool) *HttpApiRoutePoliciesFallback {
	s.Enable = &v
	return s
}

func (s *HttpApiRoutePoliciesFallback) Validate() error {
	if s.Destinations != nil {
		for _, item := range s.Destinations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiRoutePoliciesFallbackDestinations struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ServiceId       *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServicePort     *int64  `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
}

func (s HttpApiRoutePoliciesFallbackDestinations) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePoliciesFallbackDestinations) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePoliciesFallbackDestinations) GetAppId() *string {
	return s.AppId
}

func (s *HttpApiRoutePoliciesFallbackDestinations) GetAppName() *string {
	return s.AppName
}

func (s *HttpApiRoutePoliciesFallbackDestinations) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiRoutePoliciesFallbackDestinations) GetServiceName() *string {
	return s.ServiceName
}

func (s *HttpApiRoutePoliciesFallbackDestinations) GetServicePort() *int64 {
	return s.ServicePort
}

func (s *HttpApiRoutePoliciesFallbackDestinations) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *HttpApiRoutePoliciesFallbackDestinations) SetAppId(v string) *HttpApiRoutePoliciesFallbackDestinations {
	s.AppId = &v
	return s
}

func (s *HttpApiRoutePoliciesFallbackDestinations) SetAppName(v string) *HttpApiRoutePoliciesFallbackDestinations {
	s.AppName = &v
	return s
}

func (s *HttpApiRoutePoliciesFallbackDestinations) SetServiceId(v string) *HttpApiRoutePoliciesFallbackDestinations {
	s.ServiceId = &v
	return s
}

func (s *HttpApiRoutePoliciesFallbackDestinations) SetServiceName(v string) *HttpApiRoutePoliciesFallbackDestinations {
	s.ServiceName = &v
	return s
}

func (s *HttpApiRoutePoliciesFallbackDestinations) SetServicePort(v int64) *HttpApiRoutePoliciesFallbackDestinations {
	s.ServicePort = &v
	return s
}

func (s *HttpApiRoutePoliciesFallbackDestinations) SetServiceProtocol(v string) *HttpApiRoutePoliciesFallbackDestinations {
	s.ServiceProtocol = &v
	return s
}

func (s *HttpApiRoutePoliciesFallbackDestinations) Validate() error {
	return dara.Validate(s)
}

type HttpApiRoutePoliciesRetry struct {
	Attempts *int64 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// example:
	//
	// true/false
	Enable    *bool     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HttpCodes []*string `json:"HttpCodes,omitempty" xml:"HttpCodes,omitempty" type:"Repeated"`
	RetryOn   []*string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty" type:"Repeated"`
}

func (s HttpApiRoutePoliciesRetry) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePoliciesRetry) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePoliciesRetry) GetAttempts() *int64 {
	return s.Attempts
}

func (s *HttpApiRoutePoliciesRetry) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiRoutePoliciesRetry) GetHttpCodes() []*string {
	return s.HttpCodes
}

func (s *HttpApiRoutePoliciesRetry) GetRetryOn() []*string {
	return s.RetryOn
}

func (s *HttpApiRoutePoliciesRetry) SetAttempts(v int64) *HttpApiRoutePoliciesRetry {
	s.Attempts = &v
	return s
}

func (s *HttpApiRoutePoliciesRetry) SetEnable(v bool) *HttpApiRoutePoliciesRetry {
	s.Enable = &v
	return s
}

func (s *HttpApiRoutePoliciesRetry) SetHttpCodes(v []*string) *HttpApiRoutePoliciesRetry {
	s.HttpCodes = v
	return s
}

func (s *HttpApiRoutePoliciesRetry) SetRetryOn(v []*string) *HttpApiRoutePoliciesRetry {
	s.RetryOn = v
	return s
}

func (s *HttpApiRoutePoliciesRetry) Validate() error {
	return dara.Validate(s)
}

type HttpApiRoutePoliciesTimeout struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// s
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	UnitNum  *int64  `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s HttpApiRoutePoliciesTimeout) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePoliciesTimeout) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePoliciesTimeout) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiRoutePoliciesTimeout) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *HttpApiRoutePoliciesTimeout) GetUnitNum() *int64 {
	return s.UnitNum
}

func (s *HttpApiRoutePoliciesTimeout) SetEnable(v bool) *HttpApiRoutePoliciesTimeout {
	s.Enable = &v
	return s
}

func (s *HttpApiRoutePoliciesTimeout) SetTimeUnit(v string) *HttpApiRoutePoliciesTimeout {
	s.TimeUnit = &v
	return s
}

func (s *HttpApiRoutePoliciesTimeout) SetUnitNum(v int64) *HttpApiRoutePoliciesTimeout {
	s.UnitNum = &v
	return s
}

func (s *HttpApiRoutePoliciesTimeout) Validate() error {
	return dara.Validate(s)
}

type HttpApiRoutePredicates struct {
	HeaderPredicates []*HttpApiRoutePredicatesHeaderPredicates `json:"HeaderPredicates,omitempty" xml:"HeaderPredicates,omitempty" type:"Repeated"`
	MethodPredicates []*string                                 `json:"MethodPredicates,omitempty" xml:"MethodPredicates,omitempty" type:"Repeated"`
	PathPredicates   *HttpApiRoutePredicatesPathPredicates     `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
	QueryPredicates  []*HttpApiRoutePredicatesQueryPredicates  `json:"QueryPredicates,omitempty" xml:"QueryPredicates,omitempty" type:"Repeated"`
}

func (s HttpApiRoutePredicates) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePredicates) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePredicates) GetHeaderPredicates() []*HttpApiRoutePredicatesHeaderPredicates {
	return s.HeaderPredicates
}

func (s *HttpApiRoutePredicates) GetMethodPredicates() []*string {
	return s.MethodPredicates
}

func (s *HttpApiRoutePredicates) GetPathPredicates() *HttpApiRoutePredicatesPathPredicates {
	return s.PathPredicates
}

func (s *HttpApiRoutePredicates) GetQueryPredicates() []*HttpApiRoutePredicatesQueryPredicates {
	return s.QueryPredicates
}

func (s *HttpApiRoutePredicates) SetHeaderPredicates(v []*HttpApiRoutePredicatesHeaderPredicates) *HttpApiRoutePredicates {
	s.HeaderPredicates = v
	return s
}

func (s *HttpApiRoutePredicates) SetMethodPredicates(v []*string) *HttpApiRoutePredicates {
	s.MethodPredicates = v
	return s
}

func (s *HttpApiRoutePredicates) SetPathPredicates(v *HttpApiRoutePredicatesPathPredicates) *HttpApiRoutePredicates {
	s.PathPredicates = v
	return s
}

func (s *HttpApiRoutePredicates) SetQueryPredicates(v []*HttpApiRoutePredicatesQueryPredicates) *HttpApiRoutePredicates {
	s.QueryPredicates = v
	return s
}

func (s *HttpApiRoutePredicates) Validate() error {
	if s.HeaderPredicates != nil {
		for _, item := range s.HeaderPredicates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PathPredicates != nil {
		if err := s.PathPredicates.Validate(); err != nil {
			return err
		}
	}
	if s.QueryPredicates != nil {
		for _, item := range s.QueryPredicates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiRoutePredicatesHeaderPredicates struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Prefix/Exact/Regex
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s HttpApiRoutePredicatesHeaderPredicates) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePredicatesHeaderPredicates) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePredicatesHeaderPredicates) GetName() *string {
	return s.Name
}

func (s *HttpApiRoutePredicatesHeaderPredicates) GetType() *string {
	return s.Type
}

func (s *HttpApiRoutePredicatesHeaderPredicates) GetValue() *string {
	return s.Value
}

func (s *HttpApiRoutePredicatesHeaderPredicates) SetName(v string) *HttpApiRoutePredicatesHeaderPredicates {
	s.Name = &v
	return s
}

func (s *HttpApiRoutePredicatesHeaderPredicates) SetType(v string) *HttpApiRoutePredicatesHeaderPredicates {
	s.Type = &v
	return s
}

func (s *HttpApiRoutePredicatesHeaderPredicates) SetValue(v string) *HttpApiRoutePredicatesHeaderPredicates {
	s.Value = &v
	return s
}

func (s *HttpApiRoutePredicatesHeaderPredicates) Validate() error {
	return dara.Validate(s)
}

type HttpApiRoutePredicatesPathPredicates struct {
	IgnoreCase *bool   `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// Prefix/Exact/Regex
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s HttpApiRoutePredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePredicatesPathPredicates) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *HttpApiRoutePredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *HttpApiRoutePredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *HttpApiRoutePredicatesPathPredicates) SetIgnoreCase(v bool) *HttpApiRoutePredicatesPathPredicates {
	s.IgnoreCase = &v
	return s
}

func (s *HttpApiRoutePredicatesPathPredicates) SetPath(v string) *HttpApiRoutePredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *HttpApiRoutePredicatesPathPredicates) SetType(v string) *HttpApiRoutePredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *HttpApiRoutePredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}

type HttpApiRoutePredicatesQueryPredicates struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Prefix/Exact/Regex
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s HttpApiRoutePredicatesQueryPredicates) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRoutePredicatesQueryPredicates) GoString() string {
	return s.String()
}

func (s *HttpApiRoutePredicatesQueryPredicates) GetName() *string {
	return s.Name
}

func (s *HttpApiRoutePredicatesQueryPredicates) GetType() *string {
	return s.Type
}

func (s *HttpApiRoutePredicatesQueryPredicates) GetValue() *string {
	return s.Value
}

func (s *HttpApiRoutePredicatesQueryPredicates) SetName(v string) *HttpApiRoutePredicatesQueryPredicates {
	s.Name = &v
	return s
}

func (s *HttpApiRoutePredicatesQueryPredicates) SetType(v string) *HttpApiRoutePredicatesQueryPredicates {
	s.Type = &v
	return s
}

func (s *HttpApiRoutePredicatesQueryPredicates) SetValue(v string) *HttpApiRoutePredicatesQueryPredicates {
	s.Value = &v
	return s
}

func (s *HttpApiRoutePredicatesQueryPredicates) Validate() error {
	return dara.Validate(s)
}

type HttpApiRouteServices struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServicePort *int64  `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// example:
	//
	// 90
	ServiceWeight *int64 `json:"ServiceWeight,omitempty" xml:"ServiceWeight,omitempty"`
}

func (s HttpApiRouteServices) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRouteServices) GoString() string {
	return s.String()
}

func (s *HttpApiRouteServices) GetAppId() *string {
	return s.AppId
}

func (s *HttpApiRouteServices) GetAppName() *string {
	return s.AppName
}

func (s *HttpApiRouteServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiRouteServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *HttpApiRouteServices) GetServicePort() *int64 {
	return s.ServicePort
}

func (s *HttpApiRouteServices) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *HttpApiRouteServices) GetServiceWeight() *int64 {
	return s.ServiceWeight
}

func (s *HttpApiRouteServices) SetAppId(v string) *HttpApiRouteServices {
	s.AppId = &v
	return s
}

func (s *HttpApiRouteServices) SetAppName(v string) *HttpApiRouteServices {
	s.AppName = &v
	return s
}

func (s *HttpApiRouteServices) SetServiceId(v string) *HttpApiRouteServices {
	s.ServiceId = &v
	return s
}

func (s *HttpApiRouteServices) SetServiceName(v string) *HttpApiRouteServices {
	s.ServiceName = &v
	return s
}

func (s *HttpApiRouteServices) SetServicePort(v int64) *HttpApiRouteServices {
	s.ServicePort = &v
	return s
}

func (s *HttpApiRouteServices) SetServiceProtocol(v string) *HttpApiRouteServices {
	s.ServiceProtocol = &v
	return s
}

func (s *HttpApiRouteServices) SetServiceWeight(v int64) *HttpApiRouteServices {
	s.ServiceWeight = &v
	return s
}

func (s *HttpApiRouteServices) Validate() error {
	return dara.Validate(s)
}
