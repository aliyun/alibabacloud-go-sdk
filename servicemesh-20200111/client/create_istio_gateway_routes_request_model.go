// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateIstioGatewayRoutesRequest
	GetDescription() *string
	SetGatewayRoute(v *CreateIstioGatewayRoutesRequestGatewayRoute) *CreateIstioGatewayRoutesRequest
	GetGatewayRoute() *CreateIstioGatewayRoutesRequestGatewayRoute
	SetIstioGatewayName(v string) *CreateIstioGatewayRoutesRequest
	GetIstioGatewayName() *string
	SetPriority(v int32) *CreateIstioGatewayRoutesRequest
	GetPriority() *int32
	SetServiceMeshId(v string) *CreateIstioGatewayRoutesRequest
	GetServiceMeshId() *string
	SetStatus(v int32) *CreateIstioGatewayRoutesRequest
	GetStatus() *int32
}

type CreateIstioGatewayRoutesRequest struct {
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be created for the ASM gateway.
	GatewayRoute *CreateIstioGatewayRoutesRequestGatewayRoute `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty" type:"Struct"`
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

func (s CreateIstioGatewayRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIstioGatewayRoutesRequest) GetGatewayRoute() *CreateIstioGatewayRoutesRequestGatewayRoute {
	return s.GatewayRoute
}

func (s *CreateIstioGatewayRoutesRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *CreateIstioGatewayRoutesRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateIstioGatewayRoutesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateIstioGatewayRoutesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateIstioGatewayRoutesRequest) SetDescription(v string) *CreateIstioGatewayRoutesRequest {
	s.Description = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetGatewayRoute(v *CreateIstioGatewayRoutesRequestGatewayRoute) *CreateIstioGatewayRoutesRequest {
	s.GatewayRoute = v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetIstioGatewayName(v string) *CreateIstioGatewayRoutesRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetPriority(v int32) *CreateIstioGatewayRoutesRequest {
	s.Priority = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetServiceMeshId(v string) *CreateIstioGatewayRoutesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetStatus(v int32) *CreateIstioGatewayRoutesRequest {
	s.Status = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) Validate() error {
	if s.GatewayRoute != nil {
		if err := s.GatewayRoute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIstioGatewayRoutesRequestGatewayRoute struct {
	// The requested domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The advanced settings for routing HTTP traffic.
	HTTPAdvancedOptions *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions `json:"HTTPAdvancedOptions,omitempty" xml:"HTTPAdvancedOptions,omitempty" type:"Struct"`
	// The matching rules for traffic routing.
	MatchRequest *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest `json:"MatchRequest,omitempty" xml:"MatchRequest,omitempty" type:"Struct"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// A JSON string. This parameter corresponds to the three routing types in virtual services and provides configuration entries for advanced features. The value of this parameter overwrites the configurations in RouteName, RouteType, MatchRequest, and HTTPAdvancedOptions.
	//
	// example:
	//
	// {
	//
	//   "http": {
	//
	//     "route": [
	//
	//       {
	//
	//         "destination": {
	//
	//           "host": "httpbin"
	//
	//         }
	//
	//       }
	//
	//     ],
	//
	//     "name": "httpbin",
	//
	//     "match": [
	//
	//       {
	//
	//         "uri": {
	//
	//           "prefix": "/"
	//
	//         }
	//
	//       }
	//
	//     ],
	//
	//     "fault": {
	//
	//       "delay": {
	//
	//         "fixedDelay": "2s",
	//
	//         "percentage": {
	//
	//           "value": 70
	//
	//         }
	//
	//       }
	//
	//     }
	//
	//   }
	//
	// }
	RawVSRoute interface{} `json:"RawVSRoute,omitempty" xml:"RawVSRoute,omitempty"`
	// The endpoints of destination services for Layer 4 weighted routing.
	RouteDestinations []*CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations `json:"RouteDestinations,omitempty" xml:"RouteDestinations,omitempty" type:"Repeated"`
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

func (s CreateIstioGatewayRoutesRequestGatewayRoute) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRoute) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetDomains() []*string {
	return s.Domains
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetHTTPAdvancedOptions() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	return s.HTTPAdvancedOptions
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetMatchRequest() *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	return s.MatchRequest
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetRawVSRoute() interface{} {
	return s.RawVSRoute
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetRouteDestinations() []*CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	return s.RouteDestinations
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetRouteName() *string {
	return s.RouteName
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) GetRouteType() *string {
	return s.RouteType
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetDomains(v []*string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.Domains = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetHTTPAdvancedOptions(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.HTTPAdvancedOptions = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetMatchRequest(v *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.MatchRequest = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetNamespace(v string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.Namespace = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRawVSRoute(v interface{}) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RawVSRoute = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRouteDestinations(v []*CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteDestinations = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRouteName(v string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteName = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRouteType(v string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteType = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) Validate() error {
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

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions struct {
	// The virtual service that defines traffic routing.
	Delegate *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate `json:"Delegate,omitempty" xml:"Delegate,omitempty" type:"Struct"`
	// The configurations of fault injection.
	Fault *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault `json:"Fault,omitempty" xml:"Fault,omitempty" type:"Struct"`
	// The HTTP redirection rule.
	HTTPRedirect *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect `json:"HTTPRedirect,omitempty" xml:"HTTPRedirect,omitempty" type:"Struct"`
	// The configurations for mirroring HTTP traffic to another destination in addition to forwarding requests to the specified destination.
	Mirror *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror `json:"Mirror,omitempty" xml:"Mirror,omitempty" type:"Struct"`
	// The percentage of requests that are mirrored to another destination except for the original destination.
	MirrorPercentage *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage `json:"MirrorPercentage,omitempty" xml:"MirrorPercentage,omitempty" type:"Struct"`
	// The configurations of retries for failed requests.
	Retries *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries `json:"Retries,omitempty" xml:"Retries,omitempty" type:"Struct"`
	// The configurations for rewriting the virtual service.
	Rewrite *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Struct"`
	// The timeout period for requests.
	//
	// example:
	//
	// 5s
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetDelegate() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	return s.Delegate
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetFault() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	return s.Fault
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetHTTPRedirect() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	return s.HTTPRedirect
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetMirror() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	return s.Mirror
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetMirrorPercentage() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage {
	return s.MirrorPercentage
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetRetries() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	return s.Retries
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetRewrite() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	return s.Rewrite
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GetTimeout() *string {
	return s.Timeout
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetDelegate(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Delegate = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetFault(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Fault = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetHTTPRedirect(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.HTTPRedirect = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirror(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Mirror = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirrorPercentage(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.MirrorPercentage = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRetries(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Retries = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRewrite(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Rewrite = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetTimeout(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Timeout = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) Validate() error {
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

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate struct {
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GetName() *string {
	return s.Name
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetName(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Name = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetNamespace(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Namespace = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault struct {
	// The configurations for aborting requests with specified error codes.
	Abort *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort `json:"Abort,omitempty" xml:"Abort,omitempty" type:"Struct"`
	// The duration to delay a request.
	Delay *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay `json:"Delay,omitempty" xml:"Delay,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GetAbort() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	return s.Abort
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GetDelay() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	return s.Delay
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetAbort(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Abort = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetDelay(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Delay = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) Validate() error {
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

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort struct {
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The percentage of requests that are aborted with the specified error code.
	Percentage *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GetPercentage() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage {
	return s.Percentage
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetHttpStatus(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.HttpStatus = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetPercentage(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.Percentage = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) Validate() error {
	if s.Percentage != nil {
		if err := s.Percentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	//
	// example:
	//
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) GetValue() *float32 {
	return s.Value
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) SetValue(v float32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage {
	s.Value = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay struct {
	// The fixed duration for request delay.
	//
	// example:
	//
	// 5s
	FixedDelay *string `json:"FixedDelay,omitempty" xml:"FixedDelay,omitempty"`
	// The percentage of requests to which the delay fault is injected.
	Percentage *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GetFixedDelay() *string {
	return s.FixedDelay
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GetPercentage() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage {
	return s.Percentage
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetFixedDelay(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.FixedDelay = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetPercentage(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.Percentage = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) Validate() error {
	if s.Percentage != nil {
		if err := s.Percentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage struct {
	// The percentage of requests to which the delay fault is injected, which is expressed as a decimal.
	//
	// example:
	//
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) GetValue() *float32 {
	return s.Value
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) SetValue(v float32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage {
	s.Value = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect struct {
	// The value to be used to overwrite the value of the Authority or Host header during redirection.``
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GetAuthority() *string {
	return s.Authority
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GetRedirectCode() *int32 {
	return s.RedirectCode
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GetUri() *string {
	return s.Uri
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetAuthority(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Authority = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetRedirectCode(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.RedirectCode = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetUri(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Uri = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror struct {
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GetHost() *string {
	return s.Host
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GetSubset() *string {
	return s.Subset
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetHost(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Host = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetSubset(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Subset = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage struct {
	// The percentage of requests that are mirrored to another destination except for the original destination, which is expressed as a decimal.
	//
	// example:
	//
	// 0.2
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) GetValue() *float32 {
	return s.Value
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) SetValue(v float32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage {
	s.Value = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries struct {
	// The number of retries that are allowed for a request.
	//
	// example:
	//
	// 3
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The timeout period for each retry. Example: `5s`.
	//
	// example:
	//
	// 5s
	PerTryTimeout *string `json:"PerTryTimeout,omitempty" xml:"PerTryTimeout,omitempty"`
	// The condition for retries. Example: `connect-failure,refused-stream,503`.
	//
	// example:
	//
	// connect-failure,refused-stream,503
	RetryOn *string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty"`
	// Specifies whether to allow retries to other localities.
	RetryRemoteLocalities *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities `json:"RetryRemoteLocalities,omitempty" xml:"RetryRemoteLocalities,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetAttempts() *int32 {
	return s.Attempts
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetPerTryTimeout() *string {
	return s.PerTryTimeout
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetRetryOn() *string {
	return s.RetryOn
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GetRetryRemoteLocalities() *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	return s.RetryRemoteLocalities
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetAttempts(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.Attempts = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetPerTryTimeout(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.PerTryTimeout = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryOn(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryOn = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryRemoteLocalities(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryRemoteLocalities = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) Validate() error {
	if s.RetryRemoteLocalities != nil {
		if err := s.RetryRemoteLocalities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities struct {
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GetValue() *bool {
	return s.Value
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) SetValue(v bool) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	s.Value = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite struct {
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GetAuthority() *string {
	return s.Authority
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GetUri() *string {
	return s.Uri
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetAuthority(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Authority = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetUri(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Uri = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest struct {
	// The request headers to be matched.
	Headers []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ports of destination services for Layer 4 weighted routing.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The matching rule for Transport Layer Security (TLS) traffic.
	TLSMatchAttributes []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes `json:"TLSMatchAttributes,omitempty" xml:"TLSMatchAttributes,omitempty" type:"Repeated"`
	// The matching rule for URIs.
	URI *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI `json:"URI,omitempty" xml:"URI,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetHeaders() []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	return s.Headers
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetPorts() []*int32 {
	return s.Ports
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetTLSMatchAttributes() []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	return s.TLSMatchAttributes
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GetURI() *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	return s.URI
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetHeaders(v []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Headers = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetPorts(v []*int32) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Ports = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetTLSMatchAttributes(v []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.TLSMatchAttributes = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetURI(v *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.URI = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) Validate() error {
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

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders struct {
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GetMatchingContent() *string {
	return s.MatchingContent
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GetName() *string {
	return s.Name
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingContent(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingContent = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingMode(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingMode = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetName(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.Name = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes struct {
	// The Server Name Indication (SNI) values to be matched.
	SNIHosts []*string `json:"SNIHosts,omitempty" xml:"SNIHosts,omitempty" type:"Repeated"`
	// The TLS port.
	//
	// example:
	//
	// 443
	TLSPort *int32 `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GetSNIHosts() []*string {
	return s.SNIHosts
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GetTLSPort() *int32 {
	return s.TLSPort
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetSNIHosts(v []*string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.SNIHosts = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetTLSPort(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.TLSPort = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI struct {
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

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GetMatchingContent() *string {
	return s.MatchingContent
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingContent(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingContent = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingMode(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingMode = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) Validate() error {
	return dara.Validate(s)
}

type CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations struct {
	// The unique endpoint of the destination service to which the specified requests are sent.
	Destination *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Struct"`
	// The weight of the service subset.
	//
	// example:
	//
	// 80
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GetDestination() *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	return s.Destination
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetDestination(v *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Destination = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetWeight(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Weight = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) Validate() error {
	if s.Destination != nil {
		if err := s.Destination.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination struct {
	// The name of the service defined in the service registry.
	//
	// example:
	//
	// reviews
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port.
	Port *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Struct"`
	// The name of the service subset.
	//
	// example:
	//
	// v1
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GetHost() *string {
	return s.Host
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GetPort() *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort {
	return s.Port
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GetSubset() *string {
	return s.Subset
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetHost(v string) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Host = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetPort(v *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Port = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetSubset(v string) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Subset = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) Validate() error {
	if s.Port != nil {
		if err := s.Port.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort struct {
	// The port number.
	//
	// example:
	//
	// 80
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) GetNumber() *int32 {
	return s.Number
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) SetNumber(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort {
	s.Number = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) Validate() error {
	return dara.Validate(s)
}
