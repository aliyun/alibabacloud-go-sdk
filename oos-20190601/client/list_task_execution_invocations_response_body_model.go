// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskExecutionInvocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTaskExecutionInvocationsResponseBody
	GetRequestId() *string
	SetTaskExecutionInvocations(v []*ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) *ListTaskExecutionInvocationsResponseBody
	GetTaskExecutionInvocations() []*ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations
}

type ListTaskExecutionInvocationsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9E011F98-4EE5-5E3A-8FA3-D38F2C531C1F
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskExecutionInvocations []*ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations `json:"TaskExecutionInvocations,omitempty" xml:"TaskExecutionInvocations,omitempty" type:"Repeated"`
}

func (s ListTaskExecutionInvocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskExecutionInvocationsResponseBody) GetTaskExecutionInvocations() []*ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations {
	return s.TaskExecutionInvocations
}

func (s *ListTaskExecutionInvocationsResponseBody) SetRequestId(v string) *ListTaskExecutionInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBody) SetTaskExecutionInvocations(v []*ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) *ListTaskExecutionInvocationsResponseBody {
	s.TaskExecutionInvocations = v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBody) Validate() error {
	if s.TaskExecutionInvocations != nil {
		for _, item := range s.TaskExecutionInvocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations struct {
	// example:
	//
	// t-hz0jdfwd9f****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// example:
	//
	// exec-xxxx.t0001
	InvocationTaskExecutionId *string `json:"InvocationTaskExecutionId,omitempty" xml:"InvocationTaskExecutionId,omitempty"`
	// example:
	//
	// TaskName
	InvocationTaskName *string `json:"InvocationTaskName,omitempty" xml:"InvocationTaskName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) GetInvocationId() *string {
	return s.InvocationId
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) GetInvocationTaskExecutionId() *string {
	return s.InvocationTaskExecutionId
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) GetInvocationTaskName() *string {
	return s.InvocationTaskName
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) GetStatus() *string {
	return s.Status
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) SetInvocationId(v string) *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations {
	s.InvocationId = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) SetInvocationTaskExecutionId(v string) *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations {
	s.InvocationTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) SetInvocationTaskName(v string) *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations {
	s.InvocationTaskName = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) SetRegionId(v string) *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations {
	s.RegionId = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) SetStatus(v string) *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionInvocationsResponseBodyTaskExecutionInvocations) Validate() error {
	return dara.Validate(s)
}
