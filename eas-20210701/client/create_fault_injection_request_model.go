// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaultInjectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaultArgs(v interface{}) *CreateFaultInjectionRequest
	GetFaultArgs() interface{}
	SetFaultType(v string) *CreateFaultInjectionRequest
	GetFaultType() *string
}

type CreateFaultInjectionRequest struct {
	// example:
	//
	// {
	//
	//   "FaultType": "DiskFillTask",
	//
	//   "FaultArgs": {
	//
	//     "FaultAction": "fill",
	//
	//     "Percent": 80
	//
	//   }
	//
	// }
	FaultArgs interface{} `json:"FaultArgs,omitempty" xml:"FaultArgs,omitempty"`
	// example:
	//
	// CpuFullloadTask
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
}

func (s CreateFaultInjectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFaultInjectionRequest) GoString() string {
	return s.String()
}

func (s *CreateFaultInjectionRequest) GetFaultArgs() interface{} {
	return s.FaultArgs
}

func (s *CreateFaultInjectionRequest) GetFaultType() *string {
	return s.FaultType
}

func (s *CreateFaultInjectionRequest) SetFaultArgs(v interface{}) *CreateFaultInjectionRequest {
	s.FaultArgs = v
	return s
}

func (s *CreateFaultInjectionRequest) SetFaultType(v string) *CreateFaultInjectionRequest {
	s.FaultType = &v
	return s
}

func (s *CreateFaultInjectionRequest) Validate() error {
	return dara.Validate(s)
}
