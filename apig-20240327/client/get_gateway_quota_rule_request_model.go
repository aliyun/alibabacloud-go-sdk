// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerPageNumber(v string) *GetGatewayQuotaRuleRequest
	GetConsumerPageNumber() *string
	SetConsumerPageSize(v string) *GetGatewayQuotaRuleRequest
	GetConsumerPageSize() *string
	SetWithConsumers(v bool) *GetGatewayQuotaRuleRequest
	GetWithConsumers() *bool
}

type GetGatewayQuotaRuleRequest struct {
	// Deprecated
	//
	// The page number of the consumer list.
	//
	// example:
	//
	// 1
	ConsumerPageNumber *string `json:"consumerPageNumber,omitempty" xml:"consumerPageNumber,omitempty"`
	// Deprecated
	//
	// The page size of the consumer list.
	//
	// example:
	//
	// 10
	ConsumerPageSize *string `json:"consumerPageSize,omitempty" xml:"consumerPageSize,omitempty"`
	// Specifies whether to return the consumer list.
	WithConsumers *bool `json:"withConsumers,omitempty" xml:"withConsumers,omitempty"`
}

func (s GetGatewayQuotaRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleRequest) GetConsumerPageNumber() *string {
	return s.ConsumerPageNumber
}

func (s *GetGatewayQuotaRuleRequest) GetConsumerPageSize() *string {
	return s.ConsumerPageSize
}

func (s *GetGatewayQuotaRuleRequest) GetWithConsumers() *bool {
	return s.WithConsumers
}

func (s *GetGatewayQuotaRuleRequest) SetConsumerPageNumber(v string) *GetGatewayQuotaRuleRequest {
	s.ConsumerPageNumber = &v
	return s
}

func (s *GetGatewayQuotaRuleRequest) SetConsumerPageSize(v string) *GetGatewayQuotaRuleRequest {
	s.ConsumerPageSize = &v
	return s
}

func (s *GetGatewayQuotaRuleRequest) SetWithConsumers(v bool) *GetGatewayQuotaRuleRequest {
	s.WithConsumers = &v
	return s
}

func (s *GetGatewayQuotaRuleRequest) Validate() error {
	return dara.Validate(s)
}
