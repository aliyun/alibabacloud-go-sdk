// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioGatewayRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateIstioGatewayRoutesRequest
	GetDescription() *string
	SetGatewayRoute(v *UpdateIstioGatewayRoutesRequestGatewayRoute) *UpdateIstioGatewayRoutesRequest
	GetGatewayRoute() *UpdateIstioGatewayRoutesRequestGatewayRoute
	SetIstioGatewayName(v string) *UpdateIstioGatewayRoutesRequest
	GetIstioGatewayName() *string
	SetPriority(v int32) *UpdateIstioGatewayRoutesRequest
	GetPriority() *int32
	SetServiceMeshId(v string) *UpdateIstioGatewayRoutesRequest
	GetServiceMeshId() *string
	SetStatus(v int32) *UpdateIstioGatewayRoutesRequest
	GetStatus() *int32
}

type UpdateIstioGatewayRoutesRequest struct {
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be updated for the ASM gateway.
	GatewayRoute *UpdateIstioGatewayRoutesRequestGatewayRoute `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty" type:"Struct"`
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// 	- `0`: The routing rule is valid.
	//
	// 	- `1`: The routing rule is invalid.
	//
	// 	- `2`: An error occurs during the creation or update of the routing rule.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIstioGatewayRoutesRequest) GetGatewayRoute() *UpdateIstioGatewayRoutesRequestGatewayRoute {
	return s.GatewayRoute
}

func (s *UpdateIstioGatewayRoutesRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *UpdateIstioGatewayRoutesRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateIstioGatewayRoutesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateIstioGatewayRoutesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateIstioGatewayRoutesRequest) SetDescription(v string) *UpdateIstioGatewayRoutesRequest {
	s.Description = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetGatewayRoute(v *UpdateIstioGatewayRoutesRequestGatewayRoute) *UpdateIstioGatewayRoutesRequest {
	s.GatewayRoute = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetIstioGatewayName(v string) *UpdateIstioGatewayRoutesRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetPriority(v int32) *UpdateIstioGatewayRoutesRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetServiceMeshId(v string) *UpdateIstioGatewayRoutesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetStatus(v int32) *UpdateIstioGatewayRoutesRequest {
	s.Status = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) Validate() error {
	if s.GatewayRoute != nil {
		if err := s.GatewayRoute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRoute struct {
	// The list of requested domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The advanced settings for routing HTTP traffic.
	HTTPAdvancedOptions *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions `json:"HTTPAdvancedOptions,omitempty" xml:"HTTPAdvancedOptions,omitempty" type:"Struct"`
	// The matching rules for traffic routing.
	MatchRequest *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest `json:"MatchRequest,omitempty" xml:"MatchRequest,omitempty" type:"Struct"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The original YAML file of the virtual service that is serialized in a JSON string.
	//
	// example:
	//
	// {}
	RawVSRoute interface{} `json:"RawVSRoute,omitempty" xml:"RawVSRoute,omitempty"`
	// The endpoints of destination services for Layer 4 weighted routing.
	RouteDestinations []*UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations `json:"RouteDestinations,omitempty" xml:"RouteDestinations,omitempty" type:"Repeated"`
	// The name of the routing rule.
	//
	// example:
	//
	// reviews-v2-routes
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The type of the traffic to be routed. Valid values: `HTTP`, `TLS`, and `TCP`.
	//
	// example:
	//
	// HTTP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRoute) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRoute) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetDomains() []*string {
	return s.Domains
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetHTTPAdvancedOptions() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	return s.HTTPAdvancedOptions
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetMatchRequest() *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	return s.MatchRequest
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetRawVSRoute() interface{} {
	return s.RawVSRoute
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetRouteDestinations() []*UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	return s.RouteDestinations
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) GetRouteType() *string {
	return s.RouteType
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetDomains(v []*string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.Domains = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetHTTPAdvancedOptions(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.HTTPAdvancedOptions = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetMatchRequest(v *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.MatchRequest = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetNamespace(v string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.Namespace = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRawVSRoute(v interface{}) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RawVSRoute = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRouteDestinations(v []*UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteDestinations = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRouteName(v string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteName = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRouteType(v string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteType = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) Validate() error {
	if s.HTTPAdvancedOptions != nil {
		if err := s.HTTPAdvancedOptions.Validate(); err != nil {
			return err
		}
	}
	if s.MatchRequest != nil {
		if err := s.MatchRequest.Validate(); err != nil {
			return err
		}
	}
	if s.RouteDestinations != nil {
		for _, item := range s.RouteDestinations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions struct {
	// The virtual service that defines traffic routing.
	Delegate *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate `json:"Delegate,omitempty" xml:"Delegate,omitempty" type:"Struct"`
	// The configurations of fault injection.
	Fault *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault `json:"Fault,omitempty" xml:"Fault,omitempty" type:"Struct"`
	// The HTTP redirection rule.
	HTTPRedirect *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect `json:"HTTPRedirect,omitempty" xml:"HTTPRedirect,omitempty" type:"Struct"`
	// The configurations for mirroring HTTP traffic to another destination in addition to forwarding requests to the specified destination.
	Mirror *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror `json:"Mirror,omitempty" xml:"Mirror,omitempty" type:"Struct"`
	// The percentage of requests that are mirrored to another destination except for the original destination.
	MirrorPercentage *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage `json:"MirrorPercentage,omitempty" xml:"MirrorPercentage,omitempty" type:"Struct"`
	// The configurations of retries for failed requests.
	Retries *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries `json:"Retries,omitempty" xml:"Retries,omitempty" type:"Struct"`
	// The configurations for rewriting the virtual service.
	Rewrite *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Struct"`
	// The timeout period for requests.
	//
	// example:
	//
	// 5s
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetDelegate() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	return s.Delegate
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetFault() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	return s.Fault
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetHTTPRedirect() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	return s.HTTPRedirect
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetMirror() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	return s.Mirror
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetMirrorPercentage() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage {
	return s.MirrorPercentage
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetRetries() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	return s.Retries
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetRewrite() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	return s.Rewrite
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetTimeout() *string {
	return s.Timeout
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetDelegate(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Delegate = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetFault(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Fault = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetHTTPRedirect(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.HTTPRedirect = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirror(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Mirror = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirrorPercentage(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.MirrorPercentage = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRetries(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Retries = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRewrite(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Rewrite = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetTimeout(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Timeout = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) Validate() error {
	if s.Delegate != nil {
		if err := s.Delegate.Validate(); err != nil {
			return err
		}
	}
	if s.Fault != nil {
		if err := s.Fault.Validate(); err != nil {
			return err
		}
	}
	if s.HTTPRedirect != nil {
		if err := s.HTTPRedirect.Validate(); err != nil {
			return err
		}
	}
	if s.Mirror != nil {
		if err := s.Mirror.Validate(); err != nil {
			return err
		}
	}
	if s.MirrorPercentage != nil {
		if err := s.MirrorPercentage.Validate(); err != nil {
			return err
		}
	}
	if s.Retries != nil {
		if err := s.Retries.Validate(); err != nil {
			return err
		}
	}
	if s.Rewrite != nil {
		if err := s.Rewrite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate struct {
	// The name of the virtual service.
	//
	// example:
	//
	// reviews
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the virtual service belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GetName() *string {
	return s.Name
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetName(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Name = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetNamespace(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Namespace = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault struct {
	// The configurations for aborting requests with specified error codes.
	Abort *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort `json:"Abort,omitempty" xml:"Abort,omitempty" type:"Struct"`
	// The duration to delay a request.
	Delay *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay `json:"Delay,omitempty" xml:"Delay,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GetAbort() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	return s.Abort
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GetDelay() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	return s.Delay
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetAbort(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Abort = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetDelay(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Delay = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) Validate() error {
	if s.Abort != nil {
		if err := s.Abort.Validate(); err != nil {
			return err
		}
	}
	if s.Delay != nil {
		if err := s.Delay.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort struct {
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The percentage of requests that are aborted with the specified error code.
	Percentage *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GetPercentage() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage {
	return s.Percentage
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetHttpStatus(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.HttpStatus = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetPercentage(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.Percentage = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) Validate() error {
	if s.Percentage != nil {
		if err := s.Percentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	//
	// example:
	//
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) GetValue() *float32 {
	return s.Value
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) SetValue(v float32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage {
	s.Value = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay struct {
	// The fixed duration for request delay.
	//
	// example:
	//
	// 5s
	FixedDelay *string `json:"FixedDelay,omitempty" xml:"FixedDelay,omitempty"`
	// The percentage of requests to which the delay fault is injected.
	Percentage *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GetFixedDelay() *string {
	return s.FixedDelay
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GetPercentage() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage {
	return s.Percentage
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetFixedDelay(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.FixedDelay = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetPercentage(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.Percentage = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) Validate() error {
	if s.Percentage != nil {
		if err := s.Percentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage struct {
	// The percentage of requests to which the delay fault is injected, which is expressed as a decimal.
	//
	// example:
	//
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) GetValue() *float32 {
	return s.Value
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) SetValue(v float32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage {
	s.Value = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect struct {
	// The value to be used to overwrite the value of the Authority or Host header during redirection.
	//
	// example:
	//
	// newratings.default.svc.cluster.local
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The HTTP status code to be used to indicate URL redirection. Default value: 301.
	//
	// example:
	//
	// 301
	RedirectCode *int32 `json:"RedirectCode,omitempty" xml:"RedirectCode,omitempty"`
	// The value to be used to overwrite the URL path during redirection.
	//
	// example:
	//
	// /v1/getProductRatings
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GetAuthority() *string {
	return s.Authority
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GetRedirectCode() *int32 {
	return s.RedirectCode
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GetUri() *string {
	return s.Uri
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetAuthority(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Authority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetRedirectCode(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.RedirectCode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetUri(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Uri = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror struct {
	// The name of the service defined in the service registry.
	//
	// example:
	//
	// reviews.default.svc.cluster.local
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the service subset.
	//
	// example:
	//
	// v1
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GetHost() *string {
	return s.Host
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GetSubset() *string {
	return s.Subset
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetHost(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Host = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetSubset(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Subset = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage struct {
	// The percentage of requests that are mirrored to another destination except for the original destination, which is expressed as a decimal.
	//
	// example:
	//
	// 0.2
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) GetValue() *float32 {
	return s.Value
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) SetValue(v float32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage {
	s.Value = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries struct {
	// The number of retries that are allowed for a request.
	//
	// example:
	//
	// 3
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The timeout period for each retry.
	//
	// example:
	//
	// 2s
	PerTryTimeout *string `json:"PerTryTimeout,omitempty" xml:"PerTryTimeout,omitempty"`
	// The condition for retries. Example: `connect-failure,refused-stream,503`.
	//
	// example:
	//
	// connect-failure,refused-stream,503
	RetryOn *string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty"`
	// Specifies whether to allow retries to other localities.
	RetryRemoteLocalities *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities `json:"RetryRemoteLocalities,omitempty" xml:"RetryRemoteLocalities,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetAttempts() *int32 {
	return s.Attempts
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetPerTryTimeout() *string {
	return s.PerTryTimeout
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetRetryOn() *string {
	return s.RetryOn
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetRetryRemoteLocalities() *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	return s.RetryRemoteLocalities
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetAttempts(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.Attempts = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetPerTryTimeout(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.PerTryTimeout = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryOn(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryOn = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryRemoteLocalities(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryRemoteLocalities = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) Validate() error {
	if s.RetryRemoteLocalities != nil {
		if err := s.RetryRemoteLocalities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities struct {
	// Specifies whether to allow retries to other localities. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Value *bool `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GetValue() *bool {
	return s.Value
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) SetValue(v bool) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	s.Value = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite struct {
	// The value to be used to overwrite the value of the Authority or Host header.
	//
	// example:
	//
	// newratings.default.svc.cluster.local
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The value to be used to overwrite the path or prefix of the URI.
	//
	// example:
	//
	// /v1/getProductRatings
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GetAuthority() *string {
	return s.Authority
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GetUri() *string {
	return s.Uri
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetAuthority(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Authority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetUri(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Uri = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest struct {
	// The request headers to be matched.
	Headers []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ports.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The matching rule for Transport Layer Security (TLS) traffic.
	TLSMatchAttributes []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes `json:"TLSMatchAttributes,omitempty" xml:"TLSMatchAttributes,omitempty" type:"Repeated"`
	// The matching rule for URIs.
	URI *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI `json:"URI,omitempty" xml:"URI,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetHeaders() []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	return s.Headers
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetPorts() []*int32 {
	return s.Ports
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetTLSMatchAttributes() []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	return s.TLSMatchAttributes
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetURI() *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	return s.URI
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetHeaders(v []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Headers = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetPorts(v []*int32) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Ports = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetTLSMatchAttributes(v []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.TLSMatchAttributes = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetURI(v *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.URI = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TLSMatchAttributes != nil {
		for _, item := range s.TLSMatchAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.URI != nil {
		if err := s.URI.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders struct {
	// The header value to be matched.
	//
	// example:
	//
	// v1
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the header value. Valid values:
	//
	// 	- `exact`: exact match
	//
	// 	- `prefix`: match by prefix
	//
	// 	- `regex`: match by regular expression
	//
	// example:
	//
	// exact
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
	// The header key to be matched.
	//
	// example:
	//
	// x-request-id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GetMatchingContent() *string {
	return s.MatchingContent
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GetName() *string {
	return s.Name
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingContent(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingContent = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingMode(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingMode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetName(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.Name = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes struct {
	// The Server Name Indication (SNI) values to be matched.
	SNIHosts []*string `json:"SNIHosts,omitempty" xml:"SNIHosts,omitempty" type:"Repeated"`
	// The TLS port.
	//
	// example:
	//
	// 443
	TLSPort *int32 `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GetSNIHosts() []*string {
	return s.SNIHosts
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GetTLSPort() *int32 {
	return s.TLSPort
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetSNIHosts(v []*string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.SNIHosts = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetTLSPort(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.TLSPort = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI struct {
	// The content to be matched.
	//
	// example:
	//
	// /ratings/v2/
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the routing rule. Valid values:
	//
	// 	- `exact`: exact match
	//
	// 	- `prefix`: match by prefix
	//
	// 	- `regex`: match by regular expression
	//
	// example:
	//
	// prefix
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GetMatchingContent() *string {
	return s.MatchingContent
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingContent(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingContent = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingMode(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingMode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) Validate() error {
	return dara.Validate(s)
}

type UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations struct {
	// The unique endpoint of the destination service to which the specified requests are sent.
	Destination *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Struct"`
	// The weight of the service subset.
	//
	// example:
	//
	// 80
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GetDestination() *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	return s.Destination
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetDestination(v *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Destination = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetWeight(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Weight = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) Validate() error {
	if s.Destination != nil {
		if err := s.Destination.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination struct {
	// The name of the service defined in the service registry.
	//
	// example:
	//
	// reviews
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port of the destination service.
	//
	// >  If the destination service of the route has only one port, this field can be left empty. If the destination service has multiple ports, you must specify the port number.
	Port *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Struct"`
	// The name of the service subset.
	//
	// example:
	//
	// v1
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GetHost() *string {
	return s.Host
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GetPort() *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort {
	return s.Port
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GetSubset() *string {
	return s.Subset
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetHost(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Host = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetPort(v *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Port = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetSubset(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Subset = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) Validate() error {
	if s.Port != nil {
		if err := s.Port.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort struct {
	// The port number.
	//
	// example:
	//
	// 80
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) GetNumber() *int32 {
	return s.Number
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) SetNumber(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort {
	s.Number = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) Validate() error {
	return dara.Validate(s)
}
