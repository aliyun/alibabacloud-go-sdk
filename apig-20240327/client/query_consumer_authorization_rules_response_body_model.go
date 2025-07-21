// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConsumerAuthorizationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryConsumerAuthorizationRulesResponseBody
	GetCode() *string
	SetData(v *QueryConsumerAuthorizationRulesResponseBodyData) *QueryConsumerAuthorizationRulesResponseBody
	GetData() *QueryConsumerAuthorizationRulesResponseBodyData
	SetMessage(v string) *QueryConsumerAuthorizationRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryConsumerAuthorizationRulesResponseBody
	GetRequestId() *string
}

type QueryConsumerAuthorizationRulesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response parameters.
	Data *QueryConsumerAuthorizationRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A60EE5CA-1294-532A-9775-8D2FD1C6EFBF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryConsumerAuthorizationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConsumerAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConsumerAuthorizationRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryConsumerAuthorizationRulesResponseBody) GetData() *QueryConsumerAuthorizationRulesResponseBodyData {
	return s.Data
}

func (s *QueryConsumerAuthorizationRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryConsumerAuthorizationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConsumerAuthorizationRulesResponseBody) SetCode(v string) *QueryConsumerAuthorizationRulesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBody) SetData(v *QueryConsumerAuthorizationRulesResponseBodyData) *QueryConsumerAuthorizationRulesResponseBody {
	s.Data = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBody) SetMessage(v string) *QueryConsumerAuthorizationRulesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBody) SetRequestId(v string) *QueryConsumerAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryConsumerAuthorizationRulesResponseBodyData struct {
	// The rules.
	Items []*QueryConsumerAuthorizationRulesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalSize *string `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s QueryConsumerAuthorizationRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryConsumerAuthorizationRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) GetItems() []*QueryConsumerAuthorizationRulesResponseBodyDataItems {
	return s.Items
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) GetTotalSize() *string {
	return s.TotalSize
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) SetItems(v []*QueryConsumerAuthorizationRulesResponseBodyDataItems) *QueryConsumerAuthorizationRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) SetPageNumber(v int32) *QueryConsumerAuthorizationRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) SetPageSize(v int32) *QueryConsumerAuthorizationRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) SetTotalSize(v string) *QueryConsumerAuthorizationRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryConsumerAuthorizationRulesResponseBodyDataItems struct {
	// The API details.
	ApiInfo *HttpApiApiInfo `json:"apiInfo,omitempty" xml:"apiInfo,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// car-csgeka5lhtggrjcprok0
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// The consumer ID.
	//
	// example:
	//
	// cs-csheiftlhtgmp0j0hp4g
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The consumer information.
	ConsumerInfo *ConsumerInfo `json:"consumerInfo,omitempty" xml:"consumerInfo,omitempty"`
	// The creation timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The deployment status of the API in the current environment.
	//
	// example:
	//
	// {}
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// The environment information.
	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty"`
	// The expiry mode. Valid values: LongTerm and ShortTerm.
	//
	// example:
	//
	// ShortTerm
	ExpireMode *string `json:"expireMode,omitempty" xml:"expireMode,omitempty"`
	// The rule status.
	//
	// example:
	//
	// InEffect
	ExpireStatus *string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty"`
	// The time when the rule expires.
	//
	// example:
	//
	// 172086834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// The instance information.
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// The resource IDs.
	//
	// example:
	//
	// 2351944
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource information.
	ResourceInfo *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo `json:"resourceInfo,omitempty" xml:"resourceInfo,omitempty" type:"Struct"`
	// The resource type.
	//
	// example:
	//
	// HttpApiRoute
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The update timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s QueryConsumerAuthorizationRulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryConsumerAuthorizationRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetApiInfo() *HttpApiApiInfo {
	return s.ApiInfo
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetConsumerAuthorizationRuleId() *string {
	return s.ConsumerAuthorizationRuleId
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetConsumerInfo() *ConsumerInfo {
	return s.ConsumerInfo
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetEnvironmentInfo() *EnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetResourceInfo() *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo {
	return s.ResourceInfo
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetApiInfo(v *HttpApiApiInfo) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ApiInfo = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetConsumerAuthorizationRuleId(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetConsumerId(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ConsumerId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetConsumerInfo(v *ConsumerInfo) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ConsumerInfo = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetCreateTimestamp(v int64) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetDeployStatus(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.DeployStatus = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetEnvironmentInfo(v *EnvironmentInfo) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.EnvironmentInfo = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetExpireMode(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ExpireMode = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetExpireStatus(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ExpireStatus = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetExpireTimestamp(v int64) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ExpireTimestamp = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetGatewayInfo(v *GatewayInfo) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.GatewayInfo = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetResourceId(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ResourceId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetResourceInfo(v *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ResourceInfo = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetResourceType(v string) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) SetUpdateTimestamp(v int64) *QueryConsumerAuthorizationRulesResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

type QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo struct {
	// The operation information.
	OperationInfo *HttpApiOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty"`
	// The route.
	Route *HttpRoute `json:"route,omitempty" xml:"route,omitempty"`
}

func (s QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) GoString() string {
	return s.String()
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) GetOperationInfo() *HttpApiOperationInfo {
	return s.OperationInfo
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) GetRoute() *HttpRoute {
	return s.Route
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) SetOperationInfo(v *HttpApiOperationInfo) *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo {
	s.OperationInfo = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) SetRoute(v *HttpRoute) *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo {
	s.Route = v
	return s
}

func (s *QueryConsumerAuthorizationRulesResponseBodyDataItemsResourceInfo) Validate() error {
	return dara.Validate(s)
}
