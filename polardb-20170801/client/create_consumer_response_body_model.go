// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateConsumerResponseBody
	GetApiKey() *string
	SetConsumerId(v string) *CreateConsumerResponseBody
	GetConsumerId() *string
	SetKeyType(v string) *CreateConsumerResponseBody
	GetKeyType() *string
	SetRequestId(v string) *CreateConsumerResponseBody
	GetRequestId() *string
}

type CreateConsumerResponseBody struct {
	// example:
	//
	// 6c4b1f0317cd4fd7a5b446d3503d**
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// example:
	//
	// ApiKey
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateConsumerResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *CreateConsumerResponseBody) GetKeyType() *string {
	return s.KeyType
}

func (s *CreateConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerResponseBody) SetApiKey(v string) *CreateConsumerResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateConsumerResponseBody) SetConsumerId(v string) *CreateConsumerResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *CreateConsumerResponseBody) SetKeyType(v string) *CreateConsumerResponseBody {
	s.KeyType = &v
	return s
}

func (s *CreateConsumerResponseBody) SetRequestId(v string) *CreateConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
