// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeRateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetExchangeRateRequest
	GetConsoleSessionId() *string
	SetExchangeNames(v map[string]interface{}) *GetExchangeRateRequest
	GetExchangeNames() map[string]interface{}
	SetInstanceId(v string) *GetExchangeRateRequest
	GetInstanceId() *string
	SetVhostName(v string) *GetExchangeRateRequest
	GetVhostName() *string
}

type GetExchangeRateRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ExchangeNames map[string]interface{} `json:"ExchangeNames,omitempty" xml:"ExchangeNames,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetExchangeRateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeRateRequest) GoString() string {
	return s.String()
}

func (s *GetExchangeRateRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetExchangeRateRequest) GetExchangeNames() map[string]interface{} {
	return s.ExchangeNames
}

func (s *GetExchangeRateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetExchangeRateRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetExchangeRateRequest) SetConsoleSessionId(v string) *GetExchangeRateRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetExchangeRateRequest) SetExchangeNames(v map[string]interface{}) *GetExchangeRateRequest {
	s.ExchangeNames = v
	return s
}

func (s *GetExchangeRateRequest) SetInstanceId(v string) *GetExchangeRateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetExchangeRateRequest) SetVhostName(v string) *GetExchangeRateRequest {
	s.VhostName = &v
	return s
}

func (s *GetExchangeRateRequest) Validate() error {
	return dara.Validate(s)
}
