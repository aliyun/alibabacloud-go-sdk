// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGatewayRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSwimmingLaneGatewayRoutesResponseBody
	GetCode() *string
	SetData(v []*ListSwimmingLaneGatewayRoutesResponseBodyData) *ListSwimmingLaneGatewayRoutesResponseBody
	GetData() []*ListSwimmingLaneGatewayRoutesResponseBodyData
	SetErrorCode(v string) *ListSwimmingLaneGatewayRoutesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListSwimmingLaneGatewayRoutesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSwimmingLaneGatewayRoutesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSwimmingLaneGatewayRoutesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListSwimmingLaneGatewayRoutesResponseBody
	GetTraceId() *string
}

type ListSwimmingLaneGatewayRoutesResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Responses.
	Data []*ListSwimmingLaneGatewayRoutesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Valid values:
	//
	// 	- The error message returned because the request is normal and **success*	- is returned.
	//
	// 	- If the request is abnormal, the specific exception error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: Valid values:
	//
	// 	- **true**: The configurations were obtained.
	//
	// 	- **false**: The configurations failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. This parameter is used to query the exact call information.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListSwimmingLaneGatewayRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetData() []*ListSwimmingLaneGatewayRoutesResponseBodyData {
	return s.Data
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetCode(v string) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetData(v []*ListSwimmingLaneGatewayRoutesResponseBodyData) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.Data = v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetErrorCode(v string) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetMessage(v string) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetRequestId(v string) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetSuccess(v bool) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) SetTraceId(v string) *ListSwimmingLaneGatewayRoutesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSwimmingLaneGatewayRoutesResponseBodyData struct {
	// The ID of the route.
	//
	// example:
	//
	// 16933
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test-route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The routing rule.
	RoutePredicate *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate `json:"RoutePredicate,omitempty" xml:"RoutePredicate,omitempty" type:"Struct"`
}

func (s ListSwimmingLaneGatewayRoutesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGatewayRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) GetRouteName() *string {
	return s.RouteName
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) GetRoutePredicate() *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate {
	return s.RoutePredicate
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) SetRouteId(v int64) *ListSwimmingLaneGatewayRoutesResponseBodyData {
	s.RouteId = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) SetRouteName(v string) *ListSwimmingLaneGatewayRoutesResponseBodyData {
	s.RouteName = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) SetRoutePredicate(v *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate) *ListSwimmingLaneGatewayRoutesResponseBodyData {
	s.RoutePredicate = v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyData) Validate() error {
	if s.RoutePredicate != nil {
		if err := s.RoutePredicate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate struct {
	// The path matching rule.
	PathPredicate *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate `json:"PathPredicate,omitempty" xml:"PathPredicate,omitempty" type:"Struct"`
}

func (s ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate) GetPathPredicate() *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate {
	return s.PathPredicate
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate) SetPathPredicate(v *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate {
	s.PathPredicate = v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicate) Validate() error {
	if s.PathPredicate != nil {
		if err := s.PathPredicate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate struct {
	// The route URL.
	//
	// example:
	//
	// /Path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of the protection rule.
	//
	// example:
	//
	// Header
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) GetPath() *string {
	return s.Path
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) GetType() *string {
	return s.Type
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) SetPath(v string) *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate {
	s.Path = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) SetType(v string) *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate {
	s.Type = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponseBodyDataRoutePredicatePathPredicate) Validate() error {
	return dara.Validate(s)
}
