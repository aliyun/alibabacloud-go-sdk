// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConcurrencyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionArn(v string) *ConcurrencyConfig
	GetFunctionArn() *string
	SetReservedConcurrency(v int64) *ConcurrencyConfig
	GetReservedConcurrency() *int64
}

type ConcurrencyConfig struct {
	// example:
	//
	// acs:fc:cn-shanghai:123:functions/demo
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// 10
	ReservedConcurrency *int64 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s ConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s ConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *ConcurrencyConfig) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *ConcurrencyConfig) GetReservedConcurrency() *int64 {
	return s.ReservedConcurrency
}

func (s *ConcurrencyConfig) SetFunctionArn(v string) *ConcurrencyConfig {
	s.FunctionArn = &v
	return s
}

func (s *ConcurrencyConfig) SetReservedConcurrency(v int64) *ConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *ConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}
