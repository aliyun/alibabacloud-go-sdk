// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicalModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLogicalModelsResponseBody
	GetCode() *string
	SetData(v *ListLogicalModelsResponseBodyData) *ListLogicalModelsResponseBody
	GetData() *ListLogicalModelsResponseBodyData
	SetMessage(v string) *ListLogicalModelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLogicalModelsResponseBody
	GetRequestId() *string
}

type ListLogicalModelsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListLogicalModelsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLogicalModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLogicalModelsResponseBody) GetData() *ListLogicalModelsResponseBodyData {
	return s.Data
}

func (s *ListLogicalModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLogicalModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogicalModelsResponseBody) SetCode(v string) *ListLogicalModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLogicalModelsResponseBody) SetData(v *ListLogicalModelsResponseBodyData) *ListLogicalModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListLogicalModelsResponseBody) SetMessage(v string) *ListLogicalModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLogicalModelsResponseBody) SetRequestId(v string) *ListLogicalModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogicalModelsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLogicalModelsResponseBodyData struct {
	// The list of logical models.
	Items []*ListLogicalModelsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries returned for the current request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The cursor token for querying the next page. This parameter is not returned if no more data is available.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListLogicalModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyData) GetItems() []*ListLogicalModelsResponseBodyDataItems {
	return s.Items
}

func (s *ListLogicalModelsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLogicalModelsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLogicalModelsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLogicalModelsResponseBodyData) SetItems(v []*ListLogicalModelsResponseBodyDataItems) *ListLogicalModelsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListLogicalModelsResponseBodyData) SetMaxResults(v int32) *ListLogicalModelsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListLogicalModelsResponseBodyData) SetNextToken(v string) *ListLogicalModelsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLogicalModelsResponseBodyData) SetTotalCount(v int64) *ListLogicalModelsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLogicalModelsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogicalModelsResponseBodyDataItems struct {
	// The list of call paths commonly supported by all primary targets.
	AvailablePaths []*ListLogicalModelsResponseBodyDataItemsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The call strategy. This parameter is returned only for the Auto policy. Valid values: balanced, cheap, fast, and quality. This parameter can be omitted. If omitted, the default strategy is balanced. This parameter is not returned for the Mapping and Weighted policies.
	//
	// example:
	//
	// balanced
	CallStrategy *string `json:"callStrategy,omitempty" xml:"callStrategy,omitempty"`
	// The configuration delivery record. A successful API response only indicates that the request is accepted, not that the configuration has taken effect.
	Deployment *ListLogicalModelsResponseBodyDataItemsDeployment `json:"deployment,omitempty" xml:"deployment,omitempty" type:"Struct"`
	// The ordered fallback targets. A maximum of 5 fallback targets are supported. The targets are deduplicated by the actual card ID. This parameter can be an empty array.
	FallbackTargets []*ListLogicalModelsResponseBodyDataItemsFallbackTargets `json:"fallbackTargets,omitempty" xml:"fallbackTargets,omitempty" type:"Repeated"`
	// The intersection of the Boolean capabilities of all primary targets.
	//
	// example:
	//
	// {"nativeStreaming":true}
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The ID of the gateway instance to which the logical model belongs.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The logical model ID. The ID starts with "vm-".
	//
	// example:
	//
	// vm-xxxxxxxx
	LogicalModelId *string `json:"logicalModelId,omitempty" xml:"logicalModelId,omitempty"`
	// The aggregated information about the capabilities of the primary targets.
	Meta *ListLogicalModelsResponseBodyDataItemsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The client logical call name. It supports a maximum of 256 bytes in UTF-8, must not contain spaces or slashes, and must not start with "auto".
	//
	// example:
	//
	// chat-entry
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether activation is required. This parameter is returned only for the Auto policy. The value is true if the AI Credits Router commodity is not activated, and false if it is activated. This parameter is omitted if the read operation fails or the status is unknown. This parameter is omitted for the Mapping and Weighted policies.
	RequiresActivation *bool `json:"requiresActivation,omitempty" xml:"requiresActivation,omitempty"`
	// The response model type. If the value is Actual, the public name of the actual model is returned. If the value is Logical, the entry name of the logical model is returned.
	//
	// example:
	//
	// Logical
	ResponseModel *string `json:"responseModel,omitempty" xml:"responseModel,omitempty"`
	// The object version number. The initial value is 1 upon creation and increments after each successful update.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The routing policy. Valid values: Mapping, Weighted, and Auto. If the value is Auto, the targets parameter must be an empty array and the slots parameter is required.
	//
	// example:
	//
	// Mapping
	RoutingPolicy *string `json:"routingPolicy,omitempty" xml:"routingPolicy,omitempty"`
	// The slots. This parameter is returned only for the Auto policy. Each of the simple, standard, and complex slots contains 1 to 3 source cards. The three slots are globally deduplicated by modelCardId. This parameter is omitted for the Mapping and Weighted policies.
	Slots *ListLogicalModelsResponseBodyDataItemsSlots `json:"slots,omitempty" xml:"slots,omitempty" type:"Struct"`
	// The primary targets. For the Mapping policy, exactly one target is allowed. For the Weighted policy, the sum of original relative weights must be greater than 0. For the Auto policy, this must be an empty array.
	Targets []*ListLogicalModelsResponseBodyDataItemsTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// The last modification time. The time is in the ISO 8601 (RFC 3339) format and UTC time zone, such as 2026-05-15T10:00:00Z.
	//
	// example:
	//
	// 2026-05-15T10:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItems) GetAvailablePaths() []*ListLogicalModelsResponseBodyDataItemsAvailablePaths {
	return s.AvailablePaths
}

func (s *ListLogicalModelsResponseBodyDataItems) GetCallStrategy() *string {
	return s.CallStrategy
}

func (s *ListLogicalModelsResponseBodyDataItems) GetDeployment() *ListLogicalModelsResponseBodyDataItemsDeployment {
	return s.Deployment
}

func (s *ListLogicalModelsResponseBodyDataItems) GetFallbackTargets() []*ListLogicalModelsResponseBodyDataItemsFallbackTargets {
	return s.FallbackTargets
}

func (s *ListLogicalModelsResponseBodyDataItems) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *ListLogicalModelsResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListLogicalModelsResponseBodyDataItems) GetLogicalModelId() *string {
	return s.LogicalModelId
}

func (s *ListLogicalModelsResponseBodyDataItems) GetMeta() *ListLogicalModelsResponseBodyDataItemsMeta {
	return s.Meta
}

func (s *ListLogicalModelsResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListLogicalModelsResponseBodyDataItems) GetRequiresActivation() *bool {
	return s.RequiresActivation
}

func (s *ListLogicalModelsResponseBodyDataItems) GetResponseModel() *string {
	return s.ResponseModel
}

func (s *ListLogicalModelsResponseBodyDataItems) GetRevision() *int64 {
	return s.Revision
}

func (s *ListLogicalModelsResponseBodyDataItems) GetRoutingPolicy() *string {
	return s.RoutingPolicy
}

func (s *ListLogicalModelsResponseBodyDataItems) GetSlots() *ListLogicalModelsResponseBodyDataItemsSlots {
	return s.Slots
}

func (s *ListLogicalModelsResponseBodyDataItems) GetTargets() []*ListLogicalModelsResponseBodyDataItemsTargets {
	return s.Targets
}

func (s *ListLogicalModelsResponseBodyDataItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListLogicalModelsResponseBodyDataItems) SetAvailablePaths(v []*ListLogicalModelsResponseBodyDataItemsAvailablePaths) *ListLogicalModelsResponseBodyDataItems {
	s.AvailablePaths = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetCallStrategy(v string) *ListLogicalModelsResponseBodyDataItems {
	s.CallStrategy = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetDeployment(v *ListLogicalModelsResponseBodyDataItemsDeployment) *ListLogicalModelsResponseBodyDataItems {
	s.Deployment = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetFallbackTargets(v []*ListLogicalModelsResponseBodyDataItemsFallbackTargets) *ListLogicalModelsResponseBodyDataItems {
	s.FallbackTargets = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetFeatures(v map[string]interface{}) *ListLogicalModelsResponseBodyDataItems {
	s.Features = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetGatewayId(v string) *ListLogicalModelsResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetLogicalModelId(v string) *ListLogicalModelsResponseBodyDataItems {
	s.LogicalModelId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetMeta(v *ListLogicalModelsResponseBodyDataItemsMeta) *ListLogicalModelsResponseBodyDataItems {
	s.Meta = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetName(v string) *ListLogicalModelsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetRequiresActivation(v bool) *ListLogicalModelsResponseBodyDataItems {
	s.RequiresActivation = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetResponseModel(v string) *ListLogicalModelsResponseBodyDataItems {
	s.ResponseModel = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetRevision(v int64) *ListLogicalModelsResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetRoutingPolicy(v string) *ListLogicalModelsResponseBodyDataItems {
	s.RoutingPolicy = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetSlots(v *ListLogicalModelsResponseBodyDataItemsSlots) *ListLogicalModelsResponseBodyDataItems {
	s.Slots = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetTargets(v []*ListLogicalModelsResponseBodyDataItemsTargets) *ListLogicalModelsResponseBodyDataItems {
	s.Targets = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) SetUpdateTime(v string) *ListLogicalModelsResponseBodyDataItems {
	s.UpdateTime = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItems) Validate() error {
	if s.AvailablePaths != nil {
		for _, item := range s.AvailablePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Deployment != nil {
		if err := s.Deployment.Validate(); err != nil {
			return err
		}
	}
	if s.FallbackTargets != nil {
		for _, item := range s.FallbackTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	if s.Slots != nil {
		if err := s.Slots.Validate(); err != nil {
			return err
		}
	}
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogicalModelsResponseBodyDataItemsAvailablePaths struct {
	// The call path.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The path type.
	//
	// example:
	//
	// type-xxxxxxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsAvailablePaths) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *ListLogicalModelsResponseBodyDataItemsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *ListLogicalModelsResponseBodyDataItemsAvailablePaths) SetPath(v string) *ListLogicalModelsResponseBodyDataItemsAvailablePaths {
	s.Path = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsAvailablePaths) SetType(v string) *ListLogicalModelsResponseBodyDataItemsAvailablePaths {
	s.Type = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsDeployment struct {
	// The ID of the current configuration delivery operation.
	//
	// example:
	//
	// op-xxxxxxxx
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The object configuration version.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The configuration delivery status. Valid values: Pending, Running, Succeeded, PartiallyFailed, and Failed. This status is independent of the API release status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsDeployment) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsDeployment) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) GetOperationId() *string {
	return s.OperationId
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) GetRevision() *int64 {
	return s.Revision
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) GetStatus() *string {
	return s.Status
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) SetOperationId(v string) *ListLogicalModelsResponseBodyDataItemsDeployment {
	s.OperationId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) SetRevision(v int64) *ListLogicalModelsResponseBodyDataItemsDeployment {
	s.Revision = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) SetStatus(v string) *ListLogicalModelsResponseBodyDataItemsDeployment {
	s.Status = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsDeployment) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsFallbackTargets struct {
	// The ID of the actual source card within the same instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsFallbackTargets) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsFallbackTargets) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsFallbackTargets) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListLogicalModelsResponseBodyDataItemsFallbackTargets) SetModelCardId(v string) *ListLogicalModelsResponseBodyDataItemsFallbackTargets {
	s.ModelCardId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsFallbackTargets) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsMeta struct {
	// The minimum input token limit when all primary targets declare the limit. This parameter is not returned if the limit of any target is unknown.
	//
	// example:
	//
	// 128000
	MaxInputTokens *int64 `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	// The minimum output token limit when all primary targets declare the limit. This parameter is not returned if the limit of any target is unknown.
	//
	// example:
	//
	// 16384
	MaxOutputTokens *int64 `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	// The minimum total token limit when all primary targets declare the limit. This parameter is not returned if the limit of any target is unknown.
	//
	// example:
	//
	// 16384
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The input modalities commonly supported by all primary targets.
	SupportedInputModalities []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	// The output modalities commonly supported by all primary targets.
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s ListLogicalModelsResponseBodyDataItemsMeta) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsMeta) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) SetMaxInputTokens(v int64) *ListLogicalModelsResponseBodyDataItemsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) SetMaxOutputTokens(v int64) *ListLogicalModelsResponseBodyDataItemsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) SetMaxTokens(v int64) *ListLogicalModelsResponseBodyDataItemsMeta {
	s.MaxTokens = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) SetSupportedInputModalities(v []*string) *ListLogicalModelsResponseBodyDataItemsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) SetSupportedOutputModalities(v []*string) *ListLogicalModelsResponseBodyDataItemsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsMeta) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsSlots struct {
	// The complex slot, which contains 1 to 3 source cards. The three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Complex []*ListLogicalModelsResponseBodyDataItemsSlotsComplex `json:"complex,omitempty" xml:"complex,omitempty" type:"Repeated"`
	// The simple slot, which contains 1 to 3 source cards. The three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Simple []*ListLogicalModelsResponseBodyDataItemsSlotsSimple `json:"simple,omitempty" xml:"simple,omitempty" type:"Repeated"`
	// The standard slot, which contains 1 to 3 source cards. The three slots are globally deduplicated by modelCardId.
	//
	// This parameter is required.
	Standard []*ListLogicalModelsResponseBodyDataItemsSlotsStandard `json:"standard,omitempty" xml:"standard,omitempty" type:"Repeated"`
}

func (s ListLogicalModelsResponseBodyDataItemsSlots) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsSlots) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) GetComplex() []*ListLogicalModelsResponseBodyDataItemsSlotsComplex {
	return s.Complex
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) GetSimple() []*ListLogicalModelsResponseBodyDataItemsSlotsSimple {
	return s.Simple
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) GetStandard() []*ListLogicalModelsResponseBodyDataItemsSlotsStandard {
	return s.Standard
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) SetComplex(v []*ListLogicalModelsResponseBodyDataItemsSlotsComplex) *ListLogicalModelsResponseBodyDataItemsSlots {
	s.Complex = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) SetSimple(v []*ListLogicalModelsResponseBodyDataItemsSlotsSimple) *ListLogicalModelsResponseBodyDataItemsSlots {
	s.Simple = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) SetStandard(v []*ListLogicalModelsResponseBodyDataItemsSlotsStandard) *ListLogicalModelsResponseBodyDataItemsSlots {
	s.Standard = v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsSlots) Validate() error {
	if s.Complex != nil {
		for _, item := range s.Complex {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Simple != nil {
		for _, item := range s.Simple {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Standard != nil {
		for _, item := range s.Standard {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogicalModelsResponseBodyDataItemsSlotsComplex struct {
	// The ID of the actual source card within the same instance. The ID starts with "msc-". VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsSlotsComplex) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsSlotsComplex) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsComplex) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsComplex) SetModelCardId(v string) *ListLogicalModelsResponseBodyDataItemsSlotsComplex {
	s.ModelCardId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsComplex) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsSlotsSimple struct {
	// The ID of the actual source card within the same instance. The ID starts with "msc-". VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsSlotsSimple) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsSlotsSimple) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsSimple) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsSimple) SetModelCardId(v string) *ListLogicalModelsResponseBodyDataItemsSlotsSimple {
	s.ModelCardId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsSimple) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsSlotsStandard struct {
	// The ID of the actual source card within the same instance. The ID starts with "msc-". VMs or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsSlotsStandard) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsSlotsStandard) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsStandard) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsStandard) SetModelCardId(v string) *ListLogicalModelsResponseBodyDataItemsSlotsStandard {
	s.ModelCardId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsSlotsStandard) Validate() error {
	return dara.Validate(s)
}

type ListLogicalModelsResponseBodyDataItemsTargets struct {
	// The ID of the actual source model card within the same instance. The ID starts with "msc-". Virtual machines (VMs) or service overrides are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// msc-xxxxxxxx
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The weight. The value is an integer from 0 to 10000 and is saved as the original value. A value of 0 is allowed for the Mapping policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListLogicalModelsResponseBodyDataItemsTargets) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponseBodyDataItemsTargets) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponseBodyDataItemsTargets) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListLogicalModelsResponseBodyDataItemsTargets) GetWeight() *int32 {
	return s.Weight
}

func (s *ListLogicalModelsResponseBodyDataItemsTargets) SetModelCardId(v string) *ListLogicalModelsResponseBodyDataItemsTargets {
	s.ModelCardId = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsTargets) SetWeight(v int32) *ListLogicalModelsResponseBodyDataItemsTargets {
	s.Weight = &v
	return s
}

func (s *ListLogicalModelsResponseBodyDataItemsTargets) Validate() error {
	return dara.Validate(s)
}
