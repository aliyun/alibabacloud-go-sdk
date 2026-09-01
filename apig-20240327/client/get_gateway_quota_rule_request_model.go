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
	SetWithSubjects(v bool) *GetGatewayQuotaRuleRequest
	GetWithSubjects() *bool
}

type GetGatewayQuotaRuleRequest struct {
	// Deprecated
	//
	// The page number.
	//
	// example:
	//
	// 1
	ConsumerPageNumber *string `json:"consumerPageNumber,omitempty" xml:"consumerPageNumber,omitempty"`
	// Deprecated
	//
	// The page size.
	//
	// example:
	//
	// 10
	ConsumerPageSize *string `json:"consumerPageSize,omitempty" xml:"consumerPageSize,omitempty"`
	// Specifies whether to return the API consumer list.
	WithConsumers *bool `json:"withConsumers,omitempty" xml:"withConsumers,omitempty"`
	// Specifies whether to return the general subject list. This parameter applies to both API consumer and API consumer group rules.
	WithSubjects *bool `json:"withSubjects,omitempty" xml:"withSubjects,omitempty"`
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

func (s *GetGatewayQuotaRuleRequest) GetWithSubjects() *bool {
	return s.WithSubjects
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

func (s *GetGatewayQuotaRuleRequest) SetWithSubjects(v bool) *GetGatewayQuotaRuleRequest {
	s.WithSubjects = &v
	return s
}

func (s *GetGatewayQuotaRuleRequest) Validate() error {
	return dara.Validate(s)
}
