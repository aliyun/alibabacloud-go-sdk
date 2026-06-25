// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceFaultInjectionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFaultInfoList(v []*ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) *ListServiceInstanceFaultInjectionInfoResponseBody
	GetFaultInfoList() []*ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList
	SetRequestId(v string) *ListServiceInstanceFaultInjectionInfoResponseBody
	GetRequestId() *string
}

type ListServiceInstanceFaultInjectionInfoResponseBody struct {
	// A list of injected faults.
	FaultInfoList []*ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList `json:"FaultInfoList,omitempty" xml:"FaultInfoList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServiceInstanceFaultInjectionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceFaultInjectionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBody) GetFaultInfoList() []*ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList {
	return s.FaultInfoList
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBody) SetFaultInfoList(v []*ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) *ListServiceInstanceFaultInjectionInfoResponseBody {
	s.FaultInfoList = v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBody) SetRequestId(v string) *ListServiceInstanceFaultInjectionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBody) Validate() error {
	if s.FaultInfoList != nil {
		for _, item := range s.FaultInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList struct {
	// The parameters for each fault type.
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
	//   "FaultType": "DiskBurnTask",
	//
	//   "FaultArgs": {
	//
	//     "FaultAction": "burn",
	//
	//     "Read": true,
	//
	//     "Write": true,
	//
	//     "Size": 100
	//
	//   }
	//
	// }
	FaultArgs interface{} `json:"FaultArgs,omitempty" xml:"FaultArgs,omitempty"`
	// The fault status.
	FaultStatus *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus `json:"FaultStatus,omitempty" xml:"FaultStatus,omitempty" type:"Struct"`
	// The fault type. Valid values:CpuFullloadTask: a CPU full load fault.MemLoadTask: a memory load fault.NetworkTask: a network fault.DiskBurnTask: a disk read/write fault.DiskFillTask: a disk fill fault.
	//
	// example:
	//
	// DiskFillTask
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
}

func (s ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) GetFaultArgs() interface{} {
	return s.FaultArgs
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) GetFaultStatus() *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus {
	return s.FaultStatus
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) GetFaultType() *string {
	return s.FaultType
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) SetFaultArgs(v interface{}) *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList {
	s.FaultArgs = v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) SetFaultStatus(v *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList {
	s.FaultStatus = v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) SetFaultType(v string) *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList {
	s.FaultType = &v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList) Validate() error {
	if s.FaultStatus != nil {
		if err := s.FaultStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus struct {
	// The status of the fault. Valid values:
	//
	// 1. FaultNotInjected: The task was created, but the fault was not successfully injected.
	//
	// 2. FaultInjectedSuccess: The fault was successfully injected.
	//
	// 3. FaultInjectedFailure: The fault injection failed. The failure may be caused by parameter errors or system issues.
	//
	// example:
	//
	// FaultInjectedSuccess
	FaultStatus *string `json:"FaultStatus,omitempty" xml:"FaultStatus,omitempty"`
	// The description of the fault injection.
	//
	// example:
	//
	// Network interface not found
	FaultStatusMessage *string `json:"FaultStatusMessage,omitempty" xml:"FaultStatusMessage,omitempty"`
}

func (s ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) GetFaultStatus() *string {
	return s.FaultStatus
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) GetFaultStatusMessage() *string {
	return s.FaultStatusMessage
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) SetFaultStatus(v string) *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus {
	s.FaultStatus = &v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) SetFaultStatusMessage(v string) *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus {
	s.FaultStatusMessage = &v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus) Validate() error {
	return dara.Validate(s)
}
