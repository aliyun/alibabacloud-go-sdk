// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelTokenRateLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelTokenRateLimitResponseBody
	GetCode() *string
	SetData(v *GetModelTokenRateLimitResponseBodyData) *GetModelTokenRateLimitResponseBody
	GetData() *GetModelTokenRateLimitResponseBodyData
	SetMessage(v string) *GetModelTokenRateLimitResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelTokenRateLimitResponseBody
	GetRequestId() *string
}

type GetModelTokenRateLimitResponseBody struct {
	// The result code of the request. The value is Ok if the request is successful.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The saved quotas, versions, and rule IDs for the current model.
	//
	// example:
	//
	// {"gatewayId":"gw-example01","modelId":"msc-example01","revision":1,"appliedRevision":1,"globalTpm":1800,"consumers":[{"consumerId":"cs-example01","tpm":300}],"globalRuleId":"model-tpm/gw-example01/msc-example01/global","consumerRuleIds":{"cs-example01":"model-tpm/gw-example01/msc-example01/consumer/cs-example01"}}
	Data *GetModelTokenRateLimitResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The cause of the error. This field can be omitted if the request is successful.
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

func (s GetModelTokenRateLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelTokenRateLimitResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelTokenRateLimitResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelTokenRateLimitResponseBody) GetData() *GetModelTokenRateLimitResponseBodyData {
	return s.Data
}

func (s *GetModelTokenRateLimitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelTokenRateLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelTokenRateLimitResponseBody) SetCode(v string) *GetModelTokenRateLimitResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBody) SetData(v *GetModelTokenRateLimitResponseBodyData) *GetModelTokenRateLimitResponseBody {
	s.Data = v
	return s
}

func (s *GetModelTokenRateLimitResponseBody) SetMessage(v string) *GetModelTokenRateLimitResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBody) SetRequestId(v string) *GetModelTokenRateLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelTokenRateLimitResponseBodyData struct {
	// The version of the configuration that was last successfully delivered to the control plane.
	//
	// example:
	//
	// 1
	AppliedRevision *int64 `json:"appliedRevision,omitempty" xml:"appliedRevision,omitempty"`
	// The mapping from consumer IDs to model throttling rule IDs.
	//
	// example:
	//
	// {"cs-example01":"model-tpm/gw-example01/msc-example01/consumer/cs-example01"}
	ConsumerRuleIds map[string]*string `json:"consumerRuleIds,omitempty" xml:"consumerRuleIds,omitempty"`
	// The list of configured consumer quotas for this model.
	//
	// example:
	//
	// [{"consumerId":"cs-example01","tpm":300}]
	Consumers []*GetModelTokenRateLimitResponseBodyDataConsumers `json:"consumers,omitempty" xml:"consumers,omitempty" type:"Repeated"`
	// The instance ID of the dedicated AI gateway.
	//
	// example:
	//
	// gw-example01
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The rule ID for the global model quota, which can be matched with the rule hits in the gateway throttling logs. This field is returned only when a global quota is configured.
	//
	// example:
	//
	// model-tpm/gw-example01/msc-example01/global
	GlobalRuleId *string `json:"globalRuleId,omitempty" xml:"globalRuleId,omitempty"`
	// The maximum number of tokens per minute (TPM) shared by all consumers for this model within the instance.
	//
	// example:
	//
	// 1800
	GlobalTpm *int64 `json:"globalTpm,omitempty" xml:"globalTpm,omitempty"`
	// The physical model card ID that starts with msc- or the logical model ID that starts with vm-.
	//
	// example:
	//
	// msc-example01
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The saved version of the model throttling configuration. Pass this value as expectedRevision in subsequent modification requests.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
}

func (s GetModelTokenRateLimitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelTokenRateLimitResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelTokenRateLimitResponseBodyData) GetAppliedRevision() *int64 {
	return s.AppliedRevision
}

func (s *GetModelTokenRateLimitResponseBodyData) GetConsumerRuleIds() map[string]*string {
	return s.ConsumerRuleIds
}

func (s *GetModelTokenRateLimitResponseBodyData) GetConsumers() []*GetModelTokenRateLimitResponseBodyDataConsumers {
	return s.Consumers
}

func (s *GetModelTokenRateLimitResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetModelTokenRateLimitResponseBodyData) GetGlobalRuleId() *string {
	return s.GlobalRuleId
}

func (s *GetModelTokenRateLimitResponseBodyData) GetGlobalTpm() *int64 {
	return s.GlobalTpm
}

func (s *GetModelTokenRateLimitResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *GetModelTokenRateLimitResponseBodyData) GetRevision() *int64 {
	return s.Revision
}

func (s *GetModelTokenRateLimitResponseBodyData) SetAppliedRevision(v int64) *GetModelTokenRateLimitResponseBodyData {
	s.AppliedRevision = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetConsumerRuleIds(v map[string]*string) *GetModelTokenRateLimitResponseBodyData {
	s.ConsumerRuleIds = v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetConsumers(v []*GetModelTokenRateLimitResponseBodyDataConsumers) *GetModelTokenRateLimitResponseBodyData {
	s.Consumers = v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetGatewayId(v string) *GetModelTokenRateLimitResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetGlobalRuleId(v string) *GetModelTokenRateLimitResponseBodyData {
	s.GlobalRuleId = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetGlobalTpm(v int64) *GetModelTokenRateLimitResponseBodyData {
	s.GlobalTpm = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetModelId(v string) *GetModelTokenRateLimitResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) SetRevision(v int64) *GetModelTokenRateLimitResponseBodyData {
	s.Revision = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyData) Validate() error {
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

type GetModelTokenRateLimitResponseBodyDataConsumers struct {
	// The ID of the consumer.
	//
	// example:
	//
	// cs-example01
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The maximum number of tokens per minute (TPM) allowed for this consumer to use this model.
	//
	// example:
	//
	// 300
	Tpm *int64 `json:"tpm,omitempty" xml:"tpm,omitempty"`
}

func (s GetModelTokenRateLimitResponseBodyDataConsumers) String() string {
	return dara.Prettify(s)
}

func (s GetModelTokenRateLimitResponseBodyDataConsumers) GoString() string {
	return s.String()
}

func (s *GetModelTokenRateLimitResponseBodyDataConsumers) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetModelTokenRateLimitResponseBodyDataConsumers) GetTpm() *int64 {
	return s.Tpm
}

func (s *GetModelTokenRateLimitResponseBodyDataConsumers) SetConsumerId(v string) *GetModelTokenRateLimitResponseBodyDataConsumers {
	s.ConsumerId = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyDataConsumers) SetTpm(v int64) *GetModelTokenRateLimitResponseBodyDataConsumers {
	s.Tpm = &v
	return s
}

func (s *GetModelTokenRateLimitResponseBodyDataConsumers) Validate() error {
	return dara.Validate(s)
}
