// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExchangeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollina(v string) *DeleteExchangeShrinkRequest
	GetCollina() *string
	SetConsoleSessionId(v string) *DeleteExchangeShrinkRequest
	GetConsoleSessionId() *string
	SetExchangeName(v string) *DeleteExchangeShrinkRequest
	GetExchangeName() *string
	SetExchangeNamesShrink(v string) *DeleteExchangeShrinkRequest
	GetExchangeNamesShrink() *string
	SetInstanceId(v string) *DeleteExchangeShrinkRequest
	GetInstanceId() *string
	SetUmidToken(v string) *DeleteExchangeShrinkRequest
	GetUmidToken() *string
	SetVhostName(v string) *DeleteExchangeShrinkRequest
	GetVhostName() *string
}

type DeleteExchangeShrinkRequest struct {
	Collina             *string `json:"Collina,omitempty" xml:"Collina,omitempty"`
	ConsoleSessionId    *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	ExchangeName        *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	ExchangeNamesShrink *string `json:"ExchangeNames,omitempty" xml:"ExchangeNames,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UmidToken           *string `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s DeleteExchangeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExchangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteExchangeShrinkRequest) GetCollina() *string {
	return s.Collina
}

func (s *DeleteExchangeShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteExchangeShrinkRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *DeleteExchangeShrinkRequest) GetExchangeNamesShrink() *string {
	return s.ExchangeNamesShrink
}

func (s *DeleteExchangeShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteExchangeShrinkRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *DeleteExchangeShrinkRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *DeleteExchangeShrinkRequest) SetCollina(v string) *DeleteExchangeShrinkRequest {
	s.Collina = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) SetConsoleSessionId(v string) *DeleteExchangeShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) SetExchangeName(v string) *DeleteExchangeShrinkRequest {
	s.ExchangeName = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) SetExchangeNamesShrink(v string) *DeleteExchangeShrinkRequest {
	s.ExchangeNamesShrink = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) SetInstanceId(v string) *DeleteExchangeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) SetUmidToken(v string) *DeleteExchangeShrinkRequest {
	s.UmidToken = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) SetVhostName(v string) *DeleteExchangeShrinkRequest {
	s.VhostName = &v
	return s
}

func (s *DeleteExchangeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
