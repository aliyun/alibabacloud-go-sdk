// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmMachineCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAlarmMachineCountResponseBodyData) *GetAlarmMachineCountResponseBody
	GetData() *GetAlarmMachineCountResponseBodyData
	SetRequestId(v string) *GetAlarmMachineCountResponseBody
	GetRequestId() *string
}

type GetAlarmMachineCountResponseBody struct {
	// The data returned.
	Data *GetAlarmMachineCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6D3A2E7D-1238-5DD4-B3C3-BF06FCAD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlarmMachineCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmMachineCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlarmMachineCountResponseBody) GetData() *GetAlarmMachineCountResponseBodyData {
	return s.Data
}

func (s *GetAlarmMachineCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlarmMachineCountResponseBody) SetData(v *GetAlarmMachineCountResponseBodyData) *GetAlarmMachineCountResponseBody {
	s.Data = v
	return s
}

func (s *GetAlarmMachineCountResponseBody) SetRequestId(v string) *GetAlarmMachineCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlarmMachineCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlarmMachineCountResponseBodyData struct {
	// The number of servers on which alerts are generated.
	//
	// example:
	//
	// 1
	MachineCount *int32 `json:"MachineCount,omitempty" xml:"MachineCount,omitempty"`
}

func (s GetAlarmMachineCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmMachineCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlarmMachineCountResponseBodyData) GetMachineCount() *int32 {
	return s.MachineCount
}

func (s *GetAlarmMachineCountResponseBodyData) SetMachineCount(v int32) *GetAlarmMachineCountResponseBodyData {
	s.MachineCount = &v
	return s
}

func (s *GetAlarmMachineCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
