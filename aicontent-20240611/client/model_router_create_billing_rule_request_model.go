// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBillingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingType(v string) *ModelRouterCreateBillingRuleRequest
	GetBillingType() *string
	SetEffectiveTime(v string) *ModelRouterCreateBillingRuleRequest
	GetEffectiveTime() *string
	SetExpireTime(v string) *ModelRouterCreateBillingRuleRequest
	GetExpireTime() *string
	SetModelId(v int64) *ModelRouterCreateBillingRuleRequest
	GetModelId() *int64
	SetPricingConfig(v interface{}) *ModelRouterCreateBillingRuleRequest
	GetPricingConfig() interface{}
	SetVersion(v int32) *ModelRouterCreateBillingRuleRequest
	GetVersion() *int32
}

type ModelRouterCreateBillingRuleRequest struct {
	// The billing type. The value must be `configurable`.
	//
	// example:
	//
	// configurable
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// The effective time, in RFC3339 format.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// The expiration time, in RFC3339 format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The `pricingConfig` is a JSON object whose internal field structure varies depending on the billing type.
	//
	// 1. **Tiered token billing**<br>Applicable to chat models. This type uses tiered pricing based on the number of input tokens and supports different rates for standard mode, thinking mode, and cache hits.<br>JSON format:<br><br><br>
	//
	//    Field descriptions:Constraints:
	//
	// 2. **Per-image billing**<br>Applicable to `ImageGeneration` and `ImageEdit` models. Billing is based on the number of images generated or processed.<br>JSON format:<br><br><br>
	//
	//    Field descriptions:
	//
	// 3. **Video matrix billing**<br>Applicable to `VideoGeneration` and `VideoImageGeneration` models. Pricing is based on a combination of video resolution and the presence of an audio track.<br	Note: While the frontend UI may use a `matrix` field, API calls must use the `tiers` field to save the configuration. The `matrix` field is automatically converted to `tiers` on the server side. The format below is the standard API format.<br>JSON format:<br><br><br><br>
	//
	//    Field descriptions:Constraints:
	//
	// 4. **Billing by duration**<br>Applicable to automatic speech recognition (ASR) models. Billing is based on the audio duration.<br>JSON format:<br><br><br>
	//
	//    Field descriptions:
	//
	// 5. **Per-character billing**<br>Applicable to text-to-speech (TTS) models. Billing is based on the number of characters in the synthesized text.<br>JSON format:<br><br><br>
	//
	//    Field descriptions:
	//
	// 6. **Flat-rate token billing**<br>Applicable to models such as `Embedding`, `Rerank`, `MultimodalEmbedding`, and `MultimodalRerank`. This type uses a flat-rate pricing model without tiers.<br>JSON format:<br><br><br>
	//
	//    Field descriptions:
	//
	// 7. **Full-modal multi-dimensional billing**<br>Applicable to full-modal models such as `ChatFullmodal` (e.g., `qwen3.5-omni-plus`). It sets separate prices for the input and output of different modalities, such as text, audio, images, and video.<br>JSON format:<br><br><br>
	//
	//    Field descriptions:
	//
	// example:
	//
	// {
	//
	// 				"tiers":[
	//
	// 					{
	//
	// 						"input_price":2.5,
	//
	// 						"cached_input_price":2.5,
	//
	// 						"max_tokens":32000,
	//
	// 						"output_price":10,
	//
	// 						"thinking_output_price":10,
	//
	// 						"thinking_input_price":2.5,
	//
	// 						"min_tokens":0
	//
	// 					},
	//
	// 					{
	//
	// 						"input_price":4,
	//
	// 						"cached_input_price":4,
	//
	// 						"max_tokens":128000,
	//
	// 						"output_price":16,
	//
	// 						"thinking_output_price":16,
	//
	// 						"thinking_input_price":4,
	//
	// 						"min_tokens":32000
	//
	// 					},
	//
	// 					{
	//
	// 						"input_price":7,
	//
	// 						"cached_input_price":7,
	//
	// 						"max_tokens":252000,
	//
	// 						"output_price":28,
	//
	// 						"thinking_output_price":28,
	//
	// 						"thinking_input_price":7,
	//
	// 						"min_tokens":128000
	//
	// 					}
	//
	// 				]
	//
	// 			}
	PricingConfig interface{} `json:"pricingConfig,omitempty" xml:"pricingConfig,omitempty"`
	// The billing rule version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelRouterCreateBillingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBillingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBillingRuleRequest) GetBillingType() *string {
	return s.BillingType
}

func (s *ModelRouterCreateBillingRuleRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterCreateBillingRuleRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterCreateBillingRuleRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterCreateBillingRuleRequest) GetPricingConfig() interface{} {
	return s.PricingConfig
}

func (s *ModelRouterCreateBillingRuleRequest) GetVersion() *int32 {
	return s.Version
}

func (s *ModelRouterCreateBillingRuleRequest) SetBillingType(v string) *ModelRouterCreateBillingRuleRequest {
	s.BillingType = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetEffectiveTime(v string) *ModelRouterCreateBillingRuleRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetExpireTime(v string) *ModelRouterCreateBillingRuleRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetModelId(v int64) *ModelRouterCreateBillingRuleRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetPricingConfig(v interface{}) *ModelRouterCreateBillingRuleRequest {
	s.PricingConfig = v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetVersion(v int32) *ModelRouterCreateBillingRuleRequest {
	s.Version = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) Validate() error {
	return dara.Validate(s)
}
