// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelTokenRateLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumers(v []*UpdateModelTokenRateLimitRequestConsumers) *UpdateModelTokenRateLimitRequest
	GetConsumers() []*UpdateModelTokenRateLimitRequestConsumers
	SetExpectedRevision(v int64) *UpdateModelTokenRateLimitRequest
	GetExpectedRevision() *int64
	SetGlobalTpm(v int64) *UpdateModelTokenRateLimitRequest
	GetGlobalTpm() *int64
}

type UpdateModelTokenRateLimitRequest struct {
	// The complete list of consumer limits. Pass an empty array to clear the consumer limits for the model.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"consumerId":"cs-example01","tpm":300}]
	Consumers []*UpdateModelTokenRateLimitRequestConsumers `json:"consumers,omitempty" xml:"consumers,omitempty" type:"Repeated"`
	// The revision returned by the last query. Set this parameter to 0 for the first configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ExpectedRevision *int64 `json:"expectedRevision,omitempty" xml:"expectedRevision,omitempty"`
	// The global tokens per minute (TPM) limit for the model within the instance. If this parameter is omitted, the global limit is disabled.
	//
	// example:
	//
	// 1800
	GlobalTpm *int64 `json:"globalTpm,omitempty" xml:"globalTpm,omitempty"`
}

func (s UpdateModelTokenRateLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTokenRateLimitRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelTokenRateLimitRequest) GetConsumers() []*UpdateModelTokenRateLimitRequestConsumers {
	return s.Consumers
}

func (s *UpdateModelTokenRateLimitRequest) GetExpectedRevision() *int64 {
	return s.ExpectedRevision
}

func (s *UpdateModelTokenRateLimitRequest) GetGlobalTpm() *int64 {
	return s.GlobalTpm
}

func (s *UpdateModelTokenRateLimitRequest) SetConsumers(v []*UpdateModelTokenRateLimitRequestConsumers) *UpdateModelTokenRateLimitRequest {
	s.Consumers = v
	return s
}

func (s *UpdateModelTokenRateLimitRequest) SetExpectedRevision(v int64) *UpdateModelTokenRateLimitRequest {
	s.ExpectedRevision = &v
	return s
}

func (s *UpdateModelTokenRateLimitRequest) SetGlobalTpm(v int64) *UpdateModelTokenRateLimitRequest {
	s.GlobalTpm = &v
	return s
}

func (s *UpdateModelTokenRateLimitRequest) Validate() error {
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

type UpdateModelTokenRateLimitRequestConsumers struct {
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

func (s UpdateModelTokenRateLimitRequestConsumers) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTokenRateLimitRequestConsumers) GoString() string {
	return s.String()
}

func (s *UpdateModelTokenRateLimitRequestConsumers) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *UpdateModelTokenRateLimitRequestConsumers) GetTpm() *int64 {
	return s.Tpm
}

func (s *UpdateModelTokenRateLimitRequestConsumers) SetConsumerId(v string) *UpdateModelTokenRateLimitRequestConsumers {
	s.ConsumerId = &v
	return s
}

func (s *UpdateModelTokenRateLimitRequestConsumers) SetTpm(v int64) *UpdateModelTokenRateLimitRequestConsumers {
	s.Tpm = &v
	return s
}

func (s *UpdateModelTokenRateLimitRequestConsumers) Validate() error {
	return dara.Validate(s)
}
