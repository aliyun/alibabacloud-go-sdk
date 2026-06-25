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
	// The parameters for each fault type are described as follows:
	//
	// 1. CpuFullloadTask (CPU full load fault)
	//
	//    `{ "FaultType": "CpuFullloadTask", "FaultArgs": { "FaultAction": "fullload", "CpuPercent": 50 } }`
	//
	// 2. MemLoadTask (Memory load fault)
	//
	//    `{ "FaultType": "MemLoadTask", "FaultArgs": { "FaultAction": "load", "MemPercent": 80 } }`
	//
	// 3. NetworkTask (Network fault)
	//
	//    3.a. NetworkDelayAction (Network delay)
	//
	//    `{ "FaultType": "NetworkTask", "FaultArgs": { "FaultAction": "delay", "Time": 3000, "Offset": 100 } }`
	//
	//    3.b. NetworkCorruptAction (Network packet corruption)
	//
	//    `{ "FaultType": "NetworkTask", "FaultArgs": { "FaultAction": "corrupt", "Percent": 50 } }`
	//
	//    3.c. NetworkLossAction (Network packet loss)
	//
	//    `{ "FaultType": "NetworkTask", "FaultArgs": { "FaultAction": "loss", "Percent": 30 } }`
	//
	// 4. DiskBurnTask (Disk read/write fault)
	//
	//    `{ "FaultType": "DiskBurnTask", "FaultArgs": { "FaultAction": "burn", "Read": true, "Write": true, "Size": 100 } }`
	//
	// 5. DiskFillTask (Disk fill fault)
	//
	//    `{ "FaultType": "DiskFillTask", "FaultArgs": { "FaultAction": "fill", "Percent": 80 } }`
	//
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
	// The fault type.
	//
	// Device faults: 1. CPU full load fault. 2. Memory load fault. 3. Network fault. 4. Disk read/write fault. 5. Disk fill fault.
	//
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
