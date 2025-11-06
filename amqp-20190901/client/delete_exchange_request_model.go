// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExchangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollina(v string) *DeleteExchangeRequest
	GetCollina() *string
	SetConsoleSessionId(v string) *DeleteExchangeRequest
	GetConsoleSessionId() *string
	SetExchangeName(v string) *DeleteExchangeRequest
	GetExchangeName() *string
	SetExchangeNames(v map[string]interface{}) *DeleteExchangeRequest
	GetExchangeNames() map[string]interface{}
	SetInstanceId(v string) *DeleteExchangeRequest
	GetInstanceId() *string
	SetUmidToken(v string) *DeleteExchangeRequest
	GetUmidToken() *string
	SetVhostName(v string) *DeleteExchangeRequest
	GetVhostName() *string
}

type DeleteExchangeRequest struct {
	Collina          *string                `json:"Collina,omitempty" xml:"Collina,omitempty"`
	ConsoleSessionId *string                `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	ExchangeName     *string                `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	ExchangeNames    map[string]interface{} `json:"ExchangeNames,omitempty" xml:"ExchangeNames,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UmidToken        *string                `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s DeleteExchangeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExchangeRequest) GoString() string {
	return s.String()
}

func (s *DeleteExchangeRequest) GetCollina() *string {
	return s.Collina
}

func (s *DeleteExchangeRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteExchangeRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *DeleteExchangeRequest) GetExchangeNames() map[string]interface{} {
	return s.ExchangeNames
}

func (s *DeleteExchangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteExchangeRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *DeleteExchangeRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *DeleteExchangeRequest) SetCollina(v string) *DeleteExchangeRequest {
	s.Collina = &v
	return s
}

func (s *DeleteExchangeRequest) SetConsoleSessionId(v string) *DeleteExchangeRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteExchangeRequest) SetExchangeName(v string) *DeleteExchangeRequest {
	s.ExchangeName = &v
	return s
}

func (s *DeleteExchangeRequest) SetExchangeNames(v map[string]interface{}) *DeleteExchangeRequest {
	s.ExchangeNames = v
	return s
}

func (s *DeleteExchangeRequest) SetInstanceId(v string) *DeleteExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExchangeRequest) SetUmidToken(v string) *DeleteExchangeRequest {
	s.UmidToken = &v
	return s
}

func (s *DeleteExchangeRequest) SetVhostName(v string) *DeleteExchangeRequest {
	s.VhostName = &v
	return s
}

func (s *DeleteExchangeRequest) Validate() error {
	return dara.Validate(s)
}
