// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeRateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetExchangeRateShrinkRequest
	GetConsoleSessionId() *string
	SetExchangeNamesShrink(v string) *GetExchangeRateShrinkRequest
	GetExchangeNamesShrink() *string
	SetInstanceId(v string) *GetExchangeRateShrinkRequest
	GetInstanceId() *string
	SetVhostName(v string) *GetExchangeRateShrinkRequest
	GetVhostName() *string
}

type GetExchangeRateShrinkRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ExchangeNamesShrink *string `json:"ExchangeNames,omitempty" xml:"ExchangeNames,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetExchangeRateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeRateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetExchangeRateShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetExchangeRateShrinkRequest) GetExchangeNamesShrink() *string {
	return s.ExchangeNamesShrink
}

func (s *GetExchangeRateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetExchangeRateShrinkRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetExchangeRateShrinkRequest) SetConsoleSessionId(v string) *GetExchangeRateShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetExchangeRateShrinkRequest) SetExchangeNamesShrink(v string) *GetExchangeRateShrinkRequest {
	s.ExchangeNamesShrink = &v
	return s
}

func (s *GetExchangeRateShrinkRequest) SetInstanceId(v string) *GetExchangeRateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetExchangeRateShrinkRequest) SetVhostName(v string) *GetExchangeRateShrinkRequest {
	s.VhostName = &v
	return s
}

func (s *GetExchangeRateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
