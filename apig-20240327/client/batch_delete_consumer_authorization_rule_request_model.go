// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteConsumerAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerAuthorizationRuleIds(v string) *BatchDeleteConsumerAuthorizationRuleRequest
	GetConsumerAuthorizationRuleIds() *string
}

type BatchDeleteConsumerAuthorizationRuleRequest struct {
	// The rule IDs.
	//
	// example:
	//
	// car-cus2d1em1hkg7732kuk0
	ConsumerAuthorizationRuleIds *string `json:"consumerAuthorizationRuleIds,omitempty" xml:"consumerAuthorizationRuleIds,omitempty"`
}

func (s BatchDeleteConsumerAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteConsumerAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteConsumerAuthorizationRuleRequest) GetConsumerAuthorizationRuleIds() *string {
	return s.ConsumerAuthorizationRuleIds
}

func (s *BatchDeleteConsumerAuthorizationRuleRequest) SetConsumerAuthorizationRuleIds(v string) *BatchDeleteConsumerAuthorizationRuleRequest {
	s.ConsumerAuthorizationRuleIds = &v
	return s
}

func (s *BatchDeleteConsumerAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
