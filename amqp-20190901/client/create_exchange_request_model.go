// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExchangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlternateExchange(v string) *CreateExchangeRequest
	GetAlternateExchange() *string
	SetAutoDelete(v bool) *CreateExchangeRequest
	GetAutoDelete() *bool
	SetConsoleSessionId(v string) *CreateExchangeRequest
	GetConsoleSessionId() *string
	SetExchangeName(v string) *CreateExchangeRequest
	GetExchangeName() *string
	SetExchangeType(v int32) *CreateExchangeRequest
	GetExchangeType() *int32
	SetInstanceId(v string) *CreateExchangeRequest
	GetInstanceId() *string
	SetInternal(v bool) *CreateExchangeRequest
	GetInternal() *bool
	SetVhostName(v string) *CreateExchangeRequest
	GetVhostName() *string
	SetXDelayedType(v string) *CreateExchangeRequest
	GetXDelayedType() *string
	SetXHashHeader(v string) *CreateExchangeRequest
	GetXHashHeader() *string
}

type CreateExchangeRequest struct {
	AlternateExchange *string `json:"AlternateExchange,omitempty" xml:"AlternateExchange,omitempty"`
	AutoDelete        *bool   `json:"AutoDelete,omitempty" xml:"AutoDelete,omitempty"`
	ConsoleSessionId  *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	// This parameter is required.
	ExchangeType *int32  `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Internal     *bool   `json:"Internal,omitempty" xml:"Internal,omitempty"`
	// This parameter is required.
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
	XDelayedType *string `json:"XDelayedType,omitempty" xml:"XDelayedType,omitempty"`
	XHashHeader  *string `json:"XHashHeader,omitempty" xml:"XHashHeader,omitempty"`
}

func (s CreateExchangeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExchangeRequest) GoString() string {
	return s.String()
}

func (s *CreateExchangeRequest) GetAlternateExchange() *string {
	return s.AlternateExchange
}

func (s *CreateExchangeRequest) GetAutoDelete() *bool {
	return s.AutoDelete
}

func (s *CreateExchangeRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *CreateExchangeRequest) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *CreateExchangeRequest) GetExchangeType() *int32 {
	return s.ExchangeType
}

func (s *CreateExchangeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExchangeRequest) GetInternal() *bool {
	return s.Internal
}

func (s *CreateExchangeRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *CreateExchangeRequest) GetXDelayedType() *string {
	return s.XDelayedType
}

func (s *CreateExchangeRequest) GetXHashHeader() *string {
	return s.XHashHeader
}

func (s *CreateExchangeRequest) SetAlternateExchange(v string) *CreateExchangeRequest {
	s.AlternateExchange = &v
	return s
}

func (s *CreateExchangeRequest) SetAutoDelete(v bool) *CreateExchangeRequest {
	s.AutoDelete = &v
	return s
}

func (s *CreateExchangeRequest) SetConsoleSessionId(v string) *CreateExchangeRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *CreateExchangeRequest) SetExchangeName(v string) *CreateExchangeRequest {
	s.ExchangeName = &v
	return s
}

func (s *CreateExchangeRequest) SetExchangeType(v int32) *CreateExchangeRequest {
	s.ExchangeType = &v
	return s
}

func (s *CreateExchangeRequest) SetInstanceId(v string) *CreateExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExchangeRequest) SetInternal(v bool) *CreateExchangeRequest {
	s.Internal = &v
	return s
}

func (s *CreateExchangeRequest) SetVhostName(v string) *CreateExchangeRequest {
	s.VhostName = &v
	return s
}

func (s *CreateExchangeRequest) SetXDelayedType(v string) *CreateExchangeRequest {
	s.XDelayedType = &v
	return s
}

func (s *CreateExchangeRequest) SetXHashHeader(v string) *CreateExchangeRequest {
	s.XHashHeader = &v
	return s
}

func (s *CreateExchangeRequest) Validate() error {
	return dara.Validate(s)
}
