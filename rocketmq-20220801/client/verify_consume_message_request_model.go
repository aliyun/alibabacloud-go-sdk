// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConsumeMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *VerifyConsumeMessageRequest
	GetClientId() *string
	SetConsumerGroupId(v string) *VerifyConsumeMessageRequest
	GetConsumerGroupId() *string
}

type VerifyConsumeMessageRequest struct {
	// The client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// zeekr-settlement-server-dc555456f-v2lcg@1@1@qfvorazqns
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The ID of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST_FINANCE_STOCK_OUT_GROUP
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
}

func (s VerifyConsumeMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyConsumeMessageRequest) GoString() string {
	return s.String()
}

func (s *VerifyConsumeMessageRequest) GetClientId() *string {
	return s.ClientId
}

func (s *VerifyConsumeMessageRequest) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *VerifyConsumeMessageRequest) SetClientId(v string) *VerifyConsumeMessageRequest {
	s.ClientId = &v
	return s
}

func (s *VerifyConsumeMessageRequest) SetConsumerGroupId(v string) *VerifyConsumeMessageRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *VerifyConsumeMessageRequest) Validate() error {
	return dara.Validate(s)
}
