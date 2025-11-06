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
	FaultInfoList []*ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoList `json:"FaultInfoList,omitempty" xml:"FaultInfoList,omitempty" type:"Repeated"`
	// Id of the request
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
	FaultArgs   interface{}                                                                `json:"FaultArgs,omitempty" xml:"FaultArgs,omitempty"`
	FaultStatus *ListServiceInstanceFaultInjectionInfoResponseBodyFaultInfoListFaultStatus `json:"FaultStatus,omitempty" xml:"FaultStatus,omitempty" type:"Struct"`
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
	// example:
	//
	// FaultInjectedSuccess
	FaultStatus *string `json:"FaultStatus,omitempty" xml:"FaultStatus,omitempty"`
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
