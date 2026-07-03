// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAiModelProviderResponseBody
	GetCode() *string
	SetData(v *CreateAiModelProviderResponseBodyData) *CreateAiModelProviderResponseBody
	GetData() *CreateAiModelProviderResponseBodyData
	SetMessage(v string) *CreateAiModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAiModelProviderResponseBody
	GetRequestId() *string
}

type CreateAiModelProviderResponseBody struct {
	Code      *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CreateAiModelProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAiModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAiModelProviderResponseBody) GetData() *CreateAiModelProviderResponseBodyData {
	return s.Data
}

func (s *CreateAiModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAiModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAiModelProviderResponseBody) SetCode(v string) *CreateAiModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAiModelProviderResponseBody) SetData(v *CreateAiModelProviderResponseBodyData) *CreateAiModelProviderResponseBody {
	s.Data = v
	return s
}

func (s *CreateAiModelProviderResponseBody) SetMessage(v string) *CreateAiModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAiModelProviderResponseBody) SetRequestId(v string) *CreateAiModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiModelProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAiModelProviderResponseBodyData struct {
	BoundServices   []*ServiceInfo                                     `json:"boundServices,omitempty" xml:"boundServices,omitempty" type:"Repeated"`
	DisplayName     *string                                            `json:"displayName,omitempty" xml:"displayName,omitempty"`
	GatewayId       *string                                            `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ModelCards      []*CreateAiModelProviderResponseBodyDataModelCards `json:"modelCards,omitempty" xml:"modelCards,omitempty" type:"Repeated"`
	ModelCount      *int32                                             `json:"modelCount,omitempty" xml:"modelCount,omitempty"`
	ModelProviderId *string                                            `json:"modelProviderId,omitempty" xml:"modelProviderId,omitempty"`
	Provider        *string                                            `json:"provider,omitempty" xml:"provider,omitempty"`
	Source          *string                                            `json:"source,omitempty" xml:"source,omitempty"`
	UpdateTime      *string                                            `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s CreateAiModelProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponseBodyData) GetBoundServices() []*ServiceInfo {
	return s.BoundServices
}

func (s *CreateAiModelProviderResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAiModelProviderResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAiModelProviderResponseBodyData) GetModelCards() []*CreateAiModelProviderResponseBodyDataModelCards {
	return s.ModelCards
}

func (s *CreateAiModelProviderResponseBodyData) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *CreateAiModelProviderResponseBodyData) GetModelProviderId() *string {
	return s.ModelProviderId
}

func (s *CreateAiModelProviderResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *CreateAiModelProviderResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *CreateAiModelProviderResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAiModelProviderResponseBodyData) SetBoundServices(v []*ServiceInfo) *CreateAiModelProviderResponseBodyData {
	s.BoundServices = v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetDisplayName(v string) *CreateAiModelProviderResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetGatewayId(v string) *CreateAiModelProviderResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetModelCards(v []*CreateAiModelProviderResponseBodyDataModelCards) *CreateAiModelProviderResponseBodyData {
	s.ModelCards = v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetModelCount(v int32) *CreateAiModelProviderResponseBodyData {
	s.ModelCount = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetModelProviderId(v string) *CreateAiModelProviderResponseBodyData {
	s.ModelProviderId = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetProvider(v string) *CreateAiModelProviderResponseBodyData {
	s.Provider = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetSource(v string) *CreateAiModelProviderResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) SetUpdateTime(v string) *CreateAiModelProviderResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyData) Validate() error {
	if s.BoundServices != nil {
		for _, item := range s.BoundServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelCards != nil {
		for _, item := range s.ModelCards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAiModelProviderResponseBodyDataModelCards struct {
	AvailablePaths []*CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	Credit         *CreateAiModelProviderResponseBodyDataModelCardsCredit           `json:"credit,omitempty" xml:"credit,omitempty" type:"Struct"`
	Features       map[string]interface{}                                           `json:"features,omitempty" xml:"features,omitempty"`
	GatewayId      *string                                                          `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Meta           *CreateAiModelProviderResponseBodyDataModelCardsMeta             `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	ModelCardId    *string                                                          `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	ModelName      *string                                                          `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ModelProvider  *string                                                          `json:"modelProvider,omitempty" xml:"modelProvider,omitempty"`
	Source         *string                                                          `json:"source,omitempty" xml:"source,omitempty"`
	UpdateTime     *string                                                          `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s CreateAiModelProviderResponseBodyDataModelCards) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponseBodyDataModelCards) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetAvailablePaths() []*CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	return s.AvailablePaths
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetCredit() *CreateAiModelProviderResponseBodyDataModelCardsCredit {
	return s.Credit
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetMeta() *CreateAiModelProviderResponseBodyDataModelCardsMeta {
	return s.Meta
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetModelName() *string {
	return s.ModelName
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetSource() *string {
	return s.Source
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetAvailablePaths(v []*CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) *CreateAiModelProviderResponseBodyDataModelCards {
	s.AvailablePaths = v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetCredit(v *CreateAiModelProviderResponseBodyDataModelCardsCredit) *CreateAiModelProviderResponseBodyDataModelCards {
	s.Credit = v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetFeatures(v map[string]interface{}) *CreateAiModelProviderResponseBodyDataModelCards {
	s.Features = v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetGatewayId(v string) *CreateAiModelProviderResponseBodyDataModelCards {
	s.GatewayId = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetMeta(v *CreateAiModelProviderResponseBodyDataModelCardsMeta) *CreateAiModelProviderResponseBodyDataModelCards {
	s.Meta = v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetModelCardId(v string) *CreateAiModelProviderResponseBodyDataModelCards {
	s.ModelCardId = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetModelName(v string) *CreateAiModelProviderResponseBodyDataModelCards {
	s.ModelName = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetModelProvider(v string) *CreateAiModelProviderResponseBodyDataModelCards {
	s.ModelProvider = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetSource(v string) *CreateAiModelProviderResponseBodyDataModelCards {
	s.Source = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) SetUpdateTime(v string) *CreateAiModelProviderResponseBodyDataModelCards {
	s.UpdateTime = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCards) Validate() error {
	if s.AvailablePaths != nil {
		for _, item := range s.AvailablePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Credit != nil {
		if err := s.Credit.Validate(); err != nil {
			return err
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths struct {
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) GetType() *string {
	return s.Type
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) SetPath(v string) *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	s.Path = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) SetType(v string) *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths {
	s.Type = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type CreateAiModelProviderResponseBodyDataModelCardsCredit struct {
	CacheCost  *float32 `json:"cacheCost,omitempty" xml:"cacheCost,omitempty"`
	InputCost  *float32 `json:"inputCost,omitempty" xml:"inputCost,omitempty"`
	OutputCost *float32 `json:"outputCost,omitempty" xml:"outputCost,omitempty"`
	Type       *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAiModelProviderResponseBodyDataModelCardsCredit) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponseBodyDataModelCardsCredit) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) GetCacheCost() *float32 {
	return s.CacheCost
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) GetInputCost() *float32 {
	return s.InputCost
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) GetOutputCost() *float32 {
	return s.OutputCost
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) GetType() *string {
	return s.Type
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) SetCacheCost(v float32) *CreateAiModelProviderResponseBodyDataModelCardsCredit {
	s.CacheCost = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) SetInputCost(v float32) *CreateAiModelProviderResponseBodyDataModelCardsCredit {
	s.InputCost = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) SetOutputCost(v float32) *CreateAiModelProviderResponseBodyDataModelCardsCredit {
	s.OutputCost = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) SetType(v string) *CreateAiModelProviderResponseBodyDataModelCardsCredit {
	s.Type = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsCredit) Validate() error {
	return dara.Validate(s)
}

type CreateAiModelProviderResponseBodyDataModelCardsMeta struct {
	MaxInputTokens            *int64    `json:"maxInputTokens,omitempty" xml:"maxInputTokens,omitempty"`
	MaxOutputTokens           *int64    `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	MaxTokens                 *int64    `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	SupportedInputModalities  []*string `json:"supportedInputModalities,omitempty" xml:"supportedInputModalities,omitempty" type:"Repeated"`
	SupportedOutputModalities []*string `json:"supportedOutputModalities,omitempty" xml:"supportedOutputModalities,omitempty" type:"Repeated"`
}

func (s CreateAiModelProviderResponseBodyDataModelCardsMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponseBodyDataModelCardsMeta) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) GetMaxInputTokens() *int64 {
	return s.MaxInputTokens
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) GetMaxOutputTokens() *int64 {
	return s.MaxOutputTokens
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) GetSupportedInputModalities() []*string {
	return s.SupportedInputModalities
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) GetSupportedOutputModalities() []*string {
	return s.SupportedOutputModalities
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) SetMaxInputTokens(v int64) *CreateAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxInputTokens = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) SetMaxOutputTokens(v int64) *CreateAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxOutputTokens = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) SetMaxTokens(v int64) *CreateAiModelProviderResponseBodyDataModelCardsMeta {
	s.MaxTokens = &v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) SetSupportedInputModalities(v []*string) *CreateAiModelProviderResponseBodyDataModelCardsMeta {
	s.SupportedInputModalities = v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) SetSupportedOutputModalities(v []*string) *CreateAiModelProviderResponseBodyDataModelCardsMeta {
	s.SupportedOutputModalities = v
	return s
}

func (s *CreateAiModelProviderResponseBodyDataModelCardsMeta) Validate() error {
	return dara.Validate(s)
}
