// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayIsolationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateGatewayIsolationRuleResponseBodyData) *UpdateGatewayIsolationRuleResponseBody
	GetData() *UpdateGatewayIsolationRuleResponseBodyData
	SetRequestId(v string) *UpdateGatewayIsolationRuleResponseBody
	GetRequestId() *string
}

type UpdateGatewayIsolationRuleResponseBody struct {
	Data *UpdateGatewayIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 94B12406-E44D-57C9-BF93-A8B35BFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayIsolationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayIsolationRuleResponseBody) GetData() *UpdateGatewayIsolationRuleResponseBodyData {
	return s.Data
}

func (s *UpdateGatewayIsolationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayIsolationRuleResponseBody) SetData(v *UpdateGatewayIsolationRuleResponseBodyData) *UpdateGatewayIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBody) SetRequestId(v string) *UpdateGatewayIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayIsolationRuleResponseBodyData struct {
	// example:
	//
	// 0
	BehaviorType *int32 `json:"BehaviorType,omitempty" xml:"BehaviorType,omitempty"`
	// example:
	//
	// 0
	BodyEncoding *int32 `json:"BodyEncoding,omitempty" xml:"BodyEncoding,omitempty"`
	// example:
	//
	// 0
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 14407
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 358
	Id     *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LimitMode *int32 `json:"LimitMode,omitempty" xml:"LimitMode,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// key=value
	ResponseAdditionalHeaders *string `json:"ResponseAdditionalHeaders,omitempty" xml:"ResponseAdditionalHeaders,omitempty"`
	// example:
	//
	// text
	ResponseContentBody *string `json:"ResponseContentBody,omitempty" xml:"ResponseContentBody,omitempty"`
	// example:
	//
	// www.******.com
	ResponseRedirectUrl *string `json:"ResponseRedirectUrl,omitempty" xml:"ResponseRedirectUrl,omitempty"`
	// example:
	//
	// 429
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty" xml:"ResponseStatusCode,omitempty"`
	// example:
	//
	// 52853
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// routeName
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s UpdateGatewayIsolationRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetBehaviorType() *int32 {
	return s.BehaviorType
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetBodyEncoding() *int32 {
	return s.BodyEncoding
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetEnable() *int32 {
	return s.Enable
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetIdList() []*int64 {
	return s.IdList
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetLimitMode() *int32 {
	return s.LimitMode
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetResponseAdditionalHeaders() *string {
	return s.ResponseAdditionalHeaders
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetResponseContentBody() *string {
	return s.ResponseContentBody
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetResponseRedirectUrl() *string {
	return s.ResponseRedirectUrl
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetResponseStatusCode() *int32 {
	return s.ResponseStatusCode
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetBehaviorType(v int32) *UpdateGatewayIsolationRuleResponseBodyData {
	s.BehaviorType = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetBodyEncoding(v int32) *UpdateGatewayIsolationRuleResponseBodyData {
	s.BodyEncoding = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetEnable(v int32) *UpdateGatewayIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetGatewayId(v int64) *UpdateGatewayIsolationRuleResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetGatewayUniqueId(v string) *UpdateGatewayIsolationRuleResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetId(v int64) *UpdateGatewayIsolationRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetIdList(v []*int64) *UpdateGatewayIsolationRuleResponseBodyData {
	s.IdList = v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetLimitMode(v int32) *UpdateGatewayIsolationRuleResponseBodyData {
	s.LimitMode = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetMaxConcurrency(v int32) *UpdateGatewayIsolationRuleResponseBodyData {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetResponseAdditionalHeaders(v string) *UpdateGatewayIsolationRuleResponseBodyData {
	s.ResponseAdditionalHeaders = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetResponseContentBody(v string) *UpdateGatewayIsolationRuleResponseBodyData {
	s.ResponseContentBody = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetResponseRedirectUrl(v string) *UpdateGatewayIsolationRuleResponseBodyData {
	s.ResponseRedirectUrl = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetResponseStatusCode(v int32) *UpdateGatewayIsolationRuleResponseBodyData {
	s.ResponseStatusCode = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetRouteId(v int64) *UpdateGatewayIsolationRuleResponseBodyData {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) SetRouteName(v string) *UpdateGatewayIsolationRuleResponseBodyData {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
