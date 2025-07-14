// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGRPCAction interface {
	dara.Model
	String() string
	GoString() string
	SetPort(v int32) *GRPCAction
	GetPort() *int32
	SetService(v string) *GRPCAction
	GetService() *string
}

type GRPCAction struct {
	Port    *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s GRPCAction) String() string {
	return dara.Prettify(s)
}

func (s GRPCAction) GoString() string {
	return s.String()
}

func (s *GRPCAction) GetPort() *int32 {
	return s.Port
}

func (s *GRPCAction) GetService() *string {
	return s.Service
}

func (s *GRPCAction) SetPort(v int32) *GRPCAction {
	s.Port = &v
	return s
}

func (s *GRPCAction) SetService(v string) *GRPCAction {
	s.Service = &v
	return s
}

func (s *GRPCAction) Validate() error {
	return dara.Validate(s)
}
