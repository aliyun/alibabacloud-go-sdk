// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *SendMessageRequest
	GetBody() *string
	SetConsoleSessionId(v string) *SendMessageRequest
	GetConsoleSessionId() *string
	SetExchangeName(v string) *SendMessageRequest
	GetExchangeName() *string
	SetInstanceId(v string) *SendMessageRequest
	GetInstanceId() *string
	SetMessageId(v string) *SendMessageRequest
	GetMessageId() *string
	SetProps(v string) *SendMessageRequest
	GetProps() *string
	SetRoutingKey(v string) *SendMessageRequest
	GetRoutingKey() *string
	SetVhostName(v string) *SendMessageRequest
	GetVhostName() *string
}

type SendMessageRequest struct {
	// This parameter is required.
	Body             *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Props     *string `json:"Props,omitempty" xml:"Props,omitempty"`
	// This parameter is required.
	RoutingKey *string `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s SendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) GetBody() *string {
	return s.Body
}

func (s *SendMessageRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *SendMessageRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *SendMessageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *SendMessageRequest) GetProps() *string {
	return s.Props
}

func (s *SendMessageRequest) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *SendMessageRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *SendMessageRequest) SetBody(v string) *SendMessageRequest {
	s.Body = &v
	return s
}

func (s *SendMessageRequest) SetConsoleSessionId(v string) *SendMessageRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *SendMessageRequest) SetExchangeName(v string) *SendMessageRequest {
	s.ExchangeName = &v
	return s
}

func (s *SendMessageRequest) SetInstanceId(v string) *SendMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *SendMessageRequest) SetMessageId(v string) *SendMessageRequest {
	s.MessageId = &v
	return s
}

func (s *SendMessageRequest) SetProps(v string) *SendMessageRequest {
	s.Props = &v
	return s
}

func (s *SendMessageRequest) SetRoutingKey(v string) *SendMessageRequest {
	s.RoutingKey = &v
	return s
}

func (s *SendMessageRequest) SetVhostName(v string) *SendMessageRequest {
	s.VhostName = &v
	return s
}

func (s *SendMessageRequest) Validate() error {
	return dara.Validate(s)
}
