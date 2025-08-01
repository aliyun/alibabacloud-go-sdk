// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayAuthConsumerRequest
	GetAcceptLanguage() *string
	SetConsumerStatus(v bool) *ListGatewayAuthConsumerRequest
	GetConsumerStatus() *bool
	SetGatewayUniqueId(v string) *ListGatewayAuthConsumerRequest
	GetGatewayUniqueId() *string
	SetName(v string) *ListGatewayAuthConsumerRequest
	GetName() *string
	SetPageNum(v string) *ListGatewayAuthConsumerRequest
	GetPageNum() *string
	SetPageSize(v string) *ListGatewayAuthConsumerRequest
	GetPageSize() *string
	SetType(v string) *ListGatewayAuthConsumerRequest
	GetType() *string
}

type ListGatewayAuthConsumerRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The status of the consumer. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	ConsumerStatus *bool `json:"ConsumerStatus,omitempty" xml:"ConsumerStatus,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-0fe488252dc44d55a9dd57875193****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The name of the consumer.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- JWT
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayAuthConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayAuthConsumerRequest) GetConsumerStatus() *bool {
	return s.ConsumerStatus
}

func (s *ListGatewayAuthConsumerRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayAuthConsumerRequest) GetName() *string {
	return s.Name
}

func (s *ListGatewayAuthConsumerRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *ListGatewayAuthConsumerRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGatewayAuthConsumerRequest) GetType() *string {
	return s.Type
}

func (s *ListGatewayAuthConsumerRequest) SetAcceptLanguage(v string) *ListGatewayAuthConsumerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) SetConsumerStatus(v bool) *ListGatewayAuthConsumerRequest {
	s.ConsumerStatus = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) SetGatewayUniqueId(v string) *ListGatewayAuthConsumerRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) SetName(v string) *ListGatewayAuthConsumerRequest {
	s.Name = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) SetPageNum(v string) *ListGatewayAuthConsumerRequest {
	s.PageNum = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) SetPageSize(v string) *ListGatewayAuthConsumerRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) SetType(v string) *ListGatewayAuthConsumerRequest {
	s.Type = &v
	return s
}

func (s *ListGatewayAuthConsumerRequest) Validate() error {
	return dara.Validate(s)
}
