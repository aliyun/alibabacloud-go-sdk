// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetConsumerApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ResetConsumerApiKeyResponseBody
	GetApiKey() *string
	SetConsumerId(v string) *ResetConsumerApiKeyResponseBody
	GetConsumerId() *string
	SetRequestId(v string) *ResetConsumerApiKeyResponseBody
	GetRequestId() *string
}

type ResetConsumerApiKeyResponseBody struct {
	// example:
	//
	// xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AA69096-757C-4647-B36C-29EBC2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetConsumerApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetConsumerApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetConsumerApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *ResetConsumerApiKeyResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ResetConsumerApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetConsumerApiKeyResponseBody) SetApiKey(v string) *ResetConsumerApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *ResetConsumerApiKeyResponseBody) SetConsumerId(v string) *ResetConsumerApiKeyResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *ResetConsumerApiKeyResponseBody) SetRequestId(v string) *ResetConsumerApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetConsumerApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
