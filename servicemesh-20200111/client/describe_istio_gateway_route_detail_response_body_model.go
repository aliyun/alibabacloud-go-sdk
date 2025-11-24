// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayRouteDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeIstioGatewayRouteDetailResponseBody
	GetDescription() *string
	SetNamespace(v string) *DescribeIstioGatewayRouteDetailResponseBody
	GetNamespace() *string
	SetPriority(v int32) *DescribeIstioGatewayRouteDetailResponseBody
	GetPriority() *int32
	SetRequestId(v string) *DescribeIstioGatewayRouteDetailResponseBody
	GetRequestId() *string
	SetRouteDetail(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) *DescribeIstioGatewayRouteDetailResponseBody
	GetRouteDetail() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail
	SetStatus(v int32) *DescribeIstioGatewayRouteDetailResponseBody
	GetStatus() *int32
}

type DescribeIstioGatewayRouteDetailResponseBody struct {
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the routing rule.
	RouteDetail *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail `json:"RouteDetail,omitempty" xml:"RouteDetail,omitempty" type:"Struct"`
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

func (s DescribeIstioGatewayRouteDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) GetRouteDetail() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	return s.RouteDetail
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetDescription(v string) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetNamespace(v string) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetPriority(v int32) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetRequestId(v string) *DescribeIstioGatewayRouteDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetRouteDetail(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) *DescribeIstioGatewayRouteDetailResponseBody {
	s.RouteDetail = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetStatus(v int32) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) Validate() error {
	if s.RouteDetail != nil {
		if err := s.RouteDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetail struct {
	// Domain list.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The advanced settings for routing HTTP traffic.
	HTTPAdvancedOptions *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions `json:"HTTPAdvancedOptions,omitempty" xml:"HTTPAdvancedOptions,omitempty" type:"Struct"`
	// If the value is true, the original YAML file contains features that are not supported on the current interface.
	//
	// example:
	//
	// true
	HasUnsafeFeatures *bool `json:"HasUnsafeFeatures,omitempty" xml:"HasUnsafeFeatures,omitempty"`
	// The matching rules for traffic routing.
	MatchRequest *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest `json:"MatchRequest,omitempty" xml:"MatchRequest,omitempty" type:"Struct"`
	// The original YAML file of the virtual service that is serialized into a JSON string.
	//
	// example:
	//
	// {}
	RawVSRoute *string `json:"RawVSRoute,omitempty" xml:"RawVSRoute,omitempty"`
	// The endpoints of destination services for Layer 4 weighted routing.
	RouteDestinations []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations `json:"RouteDestinations,omitempty" xml:"RouteDestinations,omitempty" type:"Repeated"`
	// The name of the routing rule.
	//
	// example:
	//
	// demo-route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The type of the traffic to be routed. Valid values: `HTTP`, `TLS`, and `TCP`.
	//
	// example:
	//
	// HTTP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetHTTPAdvancedOptions() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	return s.HTTPAdvancedOptions
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetHasUnsafeFeatures() *bool {
	return s.HasUnsafeFeatures
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetMatchRequest() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	return s.MatchRequest
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetRawVSRoute() *string {
	return s.RawVSRoute
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetRouteDestinations() []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	return s.RouteDestinations
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetRouteName() *string {
	return s.RouteName
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GetRouteType() *string {
	return s.RouteType
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetDomains(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.Domains = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetHTTPAdvancedOptions(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.HTTPAdvancedOptions = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetHasUnsafeFeatures(v bool) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.HasUnsafeFeatures = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetMatchRequest(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.MatchRequest = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRawVSRoute(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RawVSRoute = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRouteDestinations(v []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RouteDestinations = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRouteName(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RouteName = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRouteType(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RouteType = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) Validate() error {
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

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions struct {
	// The virtual service that defines traffic routing.
	Delegate *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate `json:"Delegate,omitempty" xml:"Delegate,omitempty" type:"Struct"`
	// The configurations of fault injection.
	Fault *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault `json:"Fault,omitempty" xml:"Fault,omitempty" type:"Struct"`
	// The HTTP redirection rule.
	HTTPRedirect *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect `json:"HTTPRedirect,omitempty" xml:"HTTPRedirect,omitempty" type:"Struct"`
	// The configurations for mirroring HTTP traffic to another destination in addition to forwarding requests to the specified destination.
	Mirror *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror `json:"Mirror,omitempty" xml:"Mirror,omitempty" type:"Struct"`
	// The percentage of requests that are aborted with the specified error code.
	MirrorPercentage *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage `json:"MirrorPercentage,omitempty" xml:"MirrorPercentage,omitempty" type:"Struct"`
	// The configurations of retries for failed requests.
	Retries *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries `json:"Retries,omitempty" xml:"Retries,omitempty" type:"Struct"`
	// The configurations for rewriting the virtual service.
	Rewrite *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Struct"`
	// The timeout period for requests.
	//
	// example:
	//
	// 5s
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetDelegate() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate {
	return s.Delegate
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetFault() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault {
	return s.Fault
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetHTTPRedirect() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	return s.HTTPRedirect
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetMirror() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror {
	return s.Mirror
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetMirrorPercentage() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage {
	return s.MirrorPercentage
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetRetries() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	return s.Retries
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetRewrite() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite {
	return s.Rewrite
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GetTimeout() *string {
	return s.Timeout
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetDelegate(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Delegate = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetFault(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Fault = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetHTTPRedirect(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.HTTPRedirect = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetMirror(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Mirror = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetMirrorPercentage(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.MirrorPercentage = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetRetries(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Retries = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetRewrite(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Rewrite = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetTimeout(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Timeout = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) Validate() error {
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

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate struct {
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) GetName() *string {
	return s.Name
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) SetName(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate {
	s.Name = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) SetNamespace(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault struct {
	// The configurations for aborting requests with specified error codes.
	Abort *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort `json:"Abort,omitempty" xml:"Abort,omitempty" type:"Struct"`
	// The duration to delay a request.
	Delay *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay `json:"Delay,omitempty" xml:"Delay,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) GetAbort() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort {
	return s.Abort
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) GetDelay() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	return s.Delay
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) SetAbort(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault {
	s.Abort = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) SetDelay(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault {
	s.Delay = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) Validate() error {
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

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort struct {
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The percentage of requests that are aborted with the specified error code.
	Percentage *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) GetPercentage() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage {
	return s.Percentage
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) SetHttpStatus(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort {
	s.HttpStatus = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) SetPercentage(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort {
	s.Percentage = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) Validate() error {
	if s.Percentage != nil {
		if err := s.Percentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage struct {
	// The percentage of requests that are mirrored to another destination except for the original destination, which is expressed as a decimal.
	//
	// example:
	//
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) GetValue() *float32 {
	return s.Value
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) SetValue(v float32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage {
	s.Value = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay struct {
	// The duration for request delay is expressed as 2 raised to the power of x. You must specify the value of x.
	//
	// example:
	//
	// 3
	ExponentialDelay *string `json:"ExponentialDelay,omitempty" xml:"ExponentialDelay,omitempty"`
	// The fixed duration for request delay.
	//
	// example:
	//
	// 5s
	FixedDelay *string `json:"FixedDelay,omitempty" xml:"FixedDelay,omitempty"`
	// The percentage of requests to which the delay fault is injected.
	Percentage *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) GetExponentialDelay() *string {
	return s.ExponentialDelay
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) GetFixedDelay() *string {
	return s.FixedDelay
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) GetPercentage() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage {
	return s.Percentage
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) SetExponentialDelay(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	s.ExponentialDelay = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) SetFixedDelay(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	s.FixedDelay = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) SetPercentage(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	s.Percentage = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) Validate() error {
	if s.Percentage != nil {
		if err := s.Percentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	//
	// example:
	//
	// 0.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) GetValue() *float32 {
	return s.Value
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) SetValue(v float32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage {
	s.Value = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect struct {
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) GetAuthority() *string {
	return s.Authority
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) GetRedirectCode() *int32 {
	return s.RedirectCode
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) GetUri() *string {
	return s.Uri
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) SetAuthority(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	s.Authority = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) SetRedirectCode(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	s.RedirectCode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) SetUri(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	s.Uri = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror struct {
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) GetHost() *string {
	return s.Host
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) GetSubset() *string {
	return s.Subset
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) SetHost(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror {
	s.Host = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) SetSubset(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror {
	s.Subset = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	//
	// example:
	//
	// 0.2
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) GetValue() *float32 {
	return s.Value
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) SetValue(v float32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage {
	s.Value = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries struct {
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
	RetryRemoteLocalities *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities `json:"RetryRemoteLocalities,omitempty" xml:"RetryRemoteLocalities,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) GetAttempts() *int32 {
	return s.Attempts
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) GetPerTryTimeout() *string {
	return s.PerTryTimeout
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) GetRetryOn() *string {
	return s.RetryOn
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) GetRetryRemoteLocalities() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	return s.RetryRemoteLocalities
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetAttempts(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.Attempts = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetPerTryTimeout(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.PerTryTimeout = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetRetryOn(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.RetryOn = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetRetryRemoteLocalities(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.RetryRemoteLocalities = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) Validate() error {
	if s.RetryRemoteLocalities != nil {
		if err := s.RetryRemoteLocalities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities struct {
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GetValue() *bool {
	return s.Value
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) SetValue(v bool) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	s.Value = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite struct {
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) GetAuthority() *string {
	return s.Authority
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) GetUri() *string {
	return s.Uri
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) SetAuthority(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite {
	s.Authority = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) SetUri(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite {
	s.Uri = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest struct {
	// The request headers to be matched.
	Headers []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ports.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The matching rules for Transport Layer Security (TLS) traffic.
	TLSMatchAttributes []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes `json:"TLSMatchAttributes,omitempty" xml:"TLSMatchAttributes,omitempty" type:"Repeated"`
	// The matching rule for URIs.
	URI *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI `json:"URI,omitempty" xml:"URI,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) GetHeaders() []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	return s.Headers
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) GetPorts() []*int32 {
	return s.Ports
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) GetTLSMatchAttributes() []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes {
	return s.TLSMatchAttributes
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) GetURI() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI {
	return s.URI
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetHeaders(v []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetPorts(v []*int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.Ports = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetTLSMatchAttributes(v []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.TLSMatchAttributes = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetURI(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.URI = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) Validate() error {
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

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders struct {
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) GetMatchingContent() *string {
	return s.MatchingContent
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) GetName() *string {
	return s.Name
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) SetMatchingContent(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	s.MatchingContent = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) SetMatchingMode(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	s.MatchingMode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) SetName(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	s.Name = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes struct {
	// The Server Name Indication (SNI) values to be matched.
	SNIHosts []*string `json:"SNIHosts,omitempty" xml:"SNIHosts,omitempty" type:"Repeated"`
	// The TLS port.
	//
	// example:
	//
	// 443
	TLSPort *int32 `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) GetSNIHosts() []*string {
	return s.SNIHosts
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) GetTLSPort() *int32 {
	return s.TLSPort
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) SetSNIHosts(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes {
	s.SNIHosts = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) SetTLSPort(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes {
	s.TLSPort = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI struct {
	// The content to be matched.
	//
	// example:
	//
	// /ratings/v2/
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode. Valid values:
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

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) GetMatchingContent() *string {
	return s.MatchingContent
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) SetMatchingContent(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI {
	s.MatchingContent = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) SetMatchingMode(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI {
	s.MatchingMode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations struct {
	// The unique endpoint of the destination service to which the specified requests are sent.
	Destination *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Struct"`
	// The list of the request headers to be matched.
	Headers *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The traffic weight. Valid values: 1 to 100.
	//
	// example:
	//
	// 80
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) GetDestination() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	return s.Destination
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) GetHeaders() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders {
	return s.Headers
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) SetDestination(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	s.Destination = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) SetHeaders(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) SetWeight(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	s.Weight = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) Validate() error {
	if s.Destination != nil {
		if err := s.Destination.Validate(); err != nil {
			return err
		}
	}
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination struct {
	// The name of the service defined in the service registry.
	//
	// example:
	//
	// reviews
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ports of the specified hosts from which the traffic is routed.
	Port *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Struct"`
	// The name of the service subset.
	//
	// example:
	//
	// v1
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) GetHost() *string {
	return s.Host
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) GetPort() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort {
	return s.Port
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) GetSubset() *string {
	return s.Subset
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) SetHost(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	s.Host = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) SetPort(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	s.Port = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) SetSubset(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	s.Subset = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) Validate() error {
	if s.Port != nil {
		if err := s.Port.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort struct {
	// The ports of the specified hosts to which the traffic is routed.
	//
	// example:
	//
	// 443
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) GetNumber() *int32 {
	return s.Number
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) SetNumber(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort {
	s.Number = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders struct {
	// The request header to be matched.
	Request *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// The processing of the headers of the response that is to be returned.
	Response *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) GetRequest() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	return s.Request
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) GetResponse() *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	return s.Response
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) SetRequest(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders {
	s.Request = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) SetResponse(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders {
	s.Response = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) Validate() error {
	if s.Request != nil {
		if err := s.Request.Validate(); err != nil {
			return err
		}
	}
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest struct {
	// The values to be added to the header key.
	//
	// example:
	//
	// key
	Add map[string]interface{} `json:"Add,omitempty" xml:"Add,omitempty"`
	// The header value to be deleted.
	Remove []*string `json:"Remove,omitempty" xml:"Remove,omitempty" type:"Repeated"`
	// The header key to be used to overwrite the original header key.
	Set map[string]*string `json:"Set,omitempty" xml:"Set,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) GetAdd() map[string]interface{} {
	return s.Add
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) GetRemove() []*string {
	return s.Remove
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) GetSet() map[string]*string {
	return s.Set
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) SetAdd(v map[string]interface{}) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	s.Add = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) SetRemove(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	s.Remove = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) SetSet(v map[string]*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	s.Set = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse struct {
	// The values to be added to the header key.
	//
	// example:
	//
	// key
	Add map[string]interface{} `json:"Add,omitempty" xml:"Add,omitempty"`
	// The header value to be deleted.
	Remove []*string `json:"Remove,omitempty" xml:"Remove,omitempty" type:"Repeated"`
	// The header key to be used to overwrite the original header key.
	//
	// example:
	//
	// key
	Set map[string]interface{} `json:"Set,omitempty" xml:"Set,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) GetAdd() map[string]interface{} {
	return s.Add
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) GetRemove() []*string {
	return s.Remove
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) GetSet() map[string]interface{} {
	return s.Set
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) SetAdd(v map[string]interface{}) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	s.Add = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) SetRemove(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	s.Remove = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) SetSet(v map[string]interface{}) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	s.Set = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) Validate() error {
	return dara.Validate(s)
}
