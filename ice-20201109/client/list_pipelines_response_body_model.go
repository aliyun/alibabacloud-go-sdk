// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineList(v []*ListPipelinesResponseBodyPipelineList) *ListPipelinesResponseBody
	GetPipelineList() []*ListPipelinesResponseBodyPipelineList
	SetRequestId(v string) *ListPipelinesResponseBody
	GetRequestId() *string
}

type ListPipelinesResponseBody struct {
	// The queried MPS queues.
	PipelineList []*ListPipelinesResponseBodyPipelineList `json:"PipelineList,omitempty" xml:"PipelineList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) GetPipelineList() []*ListPipelinesResponseBodyPipelineList {
	return s.PipelineList
}

func (s *ListPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelinesResponseBody) SetPipelineList(v []*ListPipelinesResponseBodyPipelineList) *ListPipelinesResponseBody {
	s.PipelineList = v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) Validate() error {
	if s.PipelineList != nil {
		for _, item := range s.PipelineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelineList struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the template was last modified.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the MPS queue.
	//
	// example:
	//
	// test-pipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the MPS queue.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the MPS queue.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the MPS queue.
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The state of the MPS queue.
	//
	// Valid values:
	//
	// 	- Active
	//
	// 	- Paused
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPipelinesResponseBodyPipelineList) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelineList) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelineList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPipelinesResponseBodyPipelineList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListPipelinesResponseBodyPipelineList) GetName() *string {
	return s.Name
}

func (s *ListPipelinesResponseBodyPipelineList) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListPipelinesResponseBodyPipelineList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPipelinesResponseBodyPipelineList) GetSpeed() *string {
	return s.Speed
}

func (s *ListPipelinesResponseBodyPipelineList) GetStatus() *string {
	return s.Status
}

func (s *ListPipelinesResponseBodyPipelineList) SetCreateTime(v string) *ListPipelinesResponseBodyPipelineList {
	s.CreateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) SetModifiedTime(v string) *ListPipelinesResponseBodyPipelineList {
	s.ModifiedTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) SetName(v string) *ListPipelinesResponseBodyPipelineList {
	s.Name = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) SetPipelineId(v string) *ListPipelinesResponseBodyPipelineList {
	s.PipelineId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) SetPriority(v int32) *ListPipelinesResponseBodyPipelineList {
	s.Priority = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) SetSpeed(v string) *ListPipelinesResponseBodyPipelineList {
	s.Speed = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) SetStatus(v string) *ListPipelinesResponseBodyPipelineList {
	s.Status = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelineList) Validate() error {
	return dara.Validate(s)
}
