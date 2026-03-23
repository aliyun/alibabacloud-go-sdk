// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSyncCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*RunSyncCommandResponseBodyData) *RunSyncCommandResponseBody
	GetData() []*RunSyncCommandResponseBodyData
	SetRequestId(v string) *RunSyncCommandResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *RunSyncCommandResponseBody
	GetTotalCount() *string
}

type RunSyncCommandResponseBody struct {
	Data []*RunSyncCommandResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 31
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RunSyncCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSyncCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunSyncCommandResponseBody) GetData() []*RunSyncCommandResponseBodyData {
	return s.Data
}

func (s *RunSyncCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSyncCommandResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *RunSyncCommandResponseBody) SetData(v []*RunSyncCommandResponseBodyData) *RunSyncCommandResponseBody {
	s.Data = v
	return s
}

func (s *RunSyncCommandResponseBody) SetRequestId(v string) *RunSyncCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSyncCommandResponseBody) SetTotalCount(v string) *RunSyncCommandResponseBody {
	s.TotalCount = &v
	return s
}

func (s *RunSyncCommandResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunSyncCommandResponseBodyData struct {
	// example:
	//
	// 2022-08-11 17:45:03
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// t-15775dc8****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// example:
	//
	// RUNNING
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// example:
	//
	// success
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 2022-10-11T08:53:32Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RunSyncCommandResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunSyncCommandResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSyncCommandResponseBodyData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *RunSyncCommandResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunSyncCommandResponseBodyData) GetInvocationId() *string {
	return s.InvocationId
}

func (s *RunSyncCommandResponseBodyData) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *RunSyncCommandResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *RunSyncCommandResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *RunSyncCommandResponseBodyData) SetFinishTime(v string) *RunSyncCommandResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *RunSyncCommandResponseBodyData) SetInstanceId(v string) *RunSyncCommandResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *RunSyncCommandResponseBodyData) SetInvocationId(v string) *RunSyncCommandResponseBodyData {
	s.InvocationId = &v
	return s
}

func (s *RunSyncCommandResponseBodyData) SetInvocationStatus(v string) *RunSyncCommandResponseBodyData {
	s.InvocationStatus = &v
	return s
}

func (s *RunSyncCommandResponseBodyData) SetOutput(v string) *RunSyncCommandResponseBodyData {
	s.Output = &v
	return s
}

func (s *RunSyncCommandResponseBodyData) SetStartTime(v string) *RunSyncCommandResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *RunSyncCommandResponseBodyData) Validate() error {
	return dara.Validate(s)
}
