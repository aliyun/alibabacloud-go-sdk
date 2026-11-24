// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelTokenRateLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelTokenRateLimitResponseBody
	GetCode() *string
	SetData(v *UpdateModelTokenRateLimitResponseBodyData) *UpdateModelTokenRateLimitResponseBody
	GetData() *UpdateModelTokenRateLimitResponseBodyData
	SetMessage(v string) *UpdateModelTokenRateLimitResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelTokenRateLimitResponseBody
	GetRequestId() *string
}

type UpdateModelTokenRateLimitResponseBody struct {
	// The result code of the request. The value is Ok if the request is successful.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The saved limits, version, and rule IDs for the current model.
	//
	// example:
	//
	// {"gatewayId":"gw-example01","modelId":"msc-example01","revision":1,"appliedRevision":1,"globalTpm":1800,"consumers":[{"consumerId":"cs-example01","tpm":300}],"globalRuleId":"model-tpm/gw-example01/msc-example01/global","consumerRuleIds":{"cs-example01":"model-tpm/gw-example01/msc-example01/consumer/cs-example01"}}
	Data *UpdateModelTokenRateLimitResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The cause of the error. This parameter is omitted if the request is successful.
	//
	// example:
	//
	// The specified gateway does not exist.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 7A35E9B2-4C16-4E90-8A71-24F3BB508C01
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateModelTokenRateLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTokenRateLimitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelTokenRateLimitResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelTokenRateLimitResponseBody) GetData() *UpdateModelTokenRateLimitResponseBodyData {
	return s.Data
}

func (s *UpdateModelTokenRateLimitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelTokenRateLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelTokenRateLimitResponseBody) SetCode(v string) *UpdateModelTokenRateLimitResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBody) SetData(v *UpdateModelTokenRateLimitResponseBodyData) *UpdateModelTokenRateLimitResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBody) SetMessage(v string) *UpdateModelTokenRateLimitResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBody) SetRequestId(v string) *UpdateModelTokenRateLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelTokenRateLimitResponseBodyData struct {
	// The configuration version that was last successfully delivered to the control plane.
	//
	// example:
	//
	// 1
	AppliedRevision *int64 `json:"appliedRevision,omitempty" xml:"appliedRevision,omitempty"`
	// The mapping of consumer IDs to model throttling rule IDs.
	//
	// example:
	//
	// {"cs-example01":"model-tpm/gw-example01/msc-example01/consumer/cs-example01"}
	ConsumerRuleIds map[string]*string `json:"consumerRuleIds,omitempty" xml:"consumerRuleIds,omitempty"`
	// The list of configured consumer limits for the model.
	//
	// example:
	//
	// [{"consumerId":"cs-example01","tpm":300}]
	Consumers []*UpdateModelTokenRateLimitResponseBodyDataConsumers `json:"consumers,omitempty" xml:"consumers,omitempty" type:"Repeated"`
	// The instance ID of the dedicated AI gateway.
	//
	// example:
	//
	// gw-example01
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The rule ID for the global model limit, which can be matched with the rule hits in the gateway throttling logs. This parameter is returned only when the global limit is configured.
	//
	// example:
	//
	// model-tpm/gw-example01/msc-example01/global
	GlobalRuleId *string `json:"globalRuleId,omitempty" xml:"globalRuleId,omitempty"`
	// The global tokens per minute (TPM) limit shared by all consumers for the model within the instance.
	//
	// example:
	//
	// 1800
	GlobalTpm *int64 `json:"globalTpm,omitempty" xml:"globalTpm,omitempty"`
	// The actual model card ID that starts with msc- or the logical model ID that starts with vm-.
	//
	// example:
	//
	// msc-example01
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The saved version of the model throttling configuration. Pass this value as expectedRevision for subsequent modifications.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
}

func (s UpdateModelTokenRateLimitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTokenRateLimitResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetAppliedRevision() *int64 {
	return s.AppliedRevision
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetConsumerRuleIds() map[string]*string {
	return s.ConsumerRuleIds
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetConsumers() []*UpdateModelTokenRateLimitResponseBodyDataConsumers {
	return s.Consumers
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetGlobalRuleId() *string {
	return s.GlobalRuleId
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetGlobalTpm() *int64 {
	return s.GlobalTpm
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *UpdateModelTokenRateLimitResponseBodyData) GetRevision() *int64 {
	return s.Revision
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetAppliedRevision(v int64) *UpdateModelTokenRateLimitResponseBodyData {
	s.AppliedRevision = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetConsumerRuleIds(v map[string]*string) *UpdateModelTokenRateLimitResponseBodyData {
	s.ConsumerRuleIds = v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetConsumers(v []*UpdateModelTokenRateLimitResponseBodyDataConsumers) *UpdateModelTokenRateLimitResponseBodyData {
	s.Consumers = v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetGatewayId(v string) *UpdateModelTokenRateLimitResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetGlobalRuleId(v string) *UpdateModelTokenRateLimitResponseBodyData {
	s.GlobalRuleId = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetGlobalTpm(v int64) *UpdateModelTokenRateLimitResponseBodyData {
	s.GlobalTpm = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetModelId(v string) *UpdateModelTokenRateLimitResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) SetRevision(v int64) *UpdateModelTokenRateLimitResponseBodyData {
	s.Revision = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyData) Validate() error {
	if s.Consumers != nil {
		for _, item := range s.Consumers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateModelTokenRateLimitResponseBodyDataConsumers struct {
	// The ID of the consumer.
	//
	// This parameter is required.
	//
	// example:
	//
	// cs-example01
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The tokens per minute (TPM) limit for the consumer using this model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Tpm *int64 `json:"tpm,omitempty" xml:"tpm,omitempty"`
}

func (s UpdateModelTokenRateLimitResponseBodyDataConsumers) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTokenRateLimitResponseBodyDataConsumers) GoString() string {
	return s.String()
}

func (s *UpdateModelTokenRateLimitResponseBodyDataConsumers) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *UpdateModelTokenRateLimitResponseBodyDataConsumers) GetTpm() *int64 {
	return s.Tpm
}

func (s *UpdateModelTokenRateLimitResponseBodyDataConsumers) SetConsumerId(v string) *UpdateModelTokenRateLimitResponseBodyDataConsumers {
	s.ConsumerId = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyDataConsumers) SetTpm(v int64) *UpdateModelTokenRateLimitResponseBodyDataConsumers {
	s.Tpm = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponseBodyDataConsumers) Validate() error {
	return dara.Validate(s)
}
