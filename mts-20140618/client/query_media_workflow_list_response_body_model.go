// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaWorkflowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowList(v *QueryMediaWorkflowListResponseBodyMediaWorkflowList) *QueryMediaWorkflowListResponseBody
	GetMediaWorkflowList() *QueryMediaWorkflowListResponseBodyMediaWorkflowList
	SetNonExistMediaWorkflowIds(v *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) *QueryMediaWorkflowListResponseBody
	GetNonExistMediaWorkflowIds() *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds
	SetRequestId(v string) *QueryMediaWorkflowListResponseBody
	GetRequestId() *string
}

type QueryMediaWorkflowListResponseBody struct {
	MediaWorkflowList        *QueryMediaWorkflowListResponseBodyMediaWorkflowList        `json:"MediaWorkflowList,omitempty" xml:"MediaWorkflowList,omitempty" type:"Struct"`
	NonExistMediaWorkflowIds *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds `json:"NonExistMediaWorkflowIds,omitempty" xml:"NonExistMediaWorkflowIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16CD0CDD-457E-420D-1234-8385075A618B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaWorkflowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowListResponseBody) GetMediaWorkflowList() *QueryMediaWorkflowListResponseBodyMediaWorkflowList {
	return s.MediaWorkflowList
}

func (s *QueryMediaWorkflowListResponseBody) GetNonExistMediaWorkflowIds() *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds {
	return s.NonExistMediaWorkflowIds
}

func (s *QueryMediaWorkflowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaWorkflowListResponseBody) SetMediaWorkflowList(v *QueryMediaWorkflowListResponseBodyMediaWorkflowList) *QueryMediaWorkflowListResponseBody {
	s.MediaWorkflowList = v
	return s
}

func (s *QueryMediaWorkflowListResponseBody) SetNonExistMediaWorkflowIds(v *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) *QueryMediaWorkflowListResponseBody {
	s.NonExistMediaWorkflowIds = v
	return s
}

func (s *QueryMediaWorkflowListResponseBody) SetRequestId(v string) *QueryMediaWorkflowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBody) Validate() error {
	if s.MediaWorkflowList != nil {
		if err := s.MediaWorkflowList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistMediaWorkflowIds != nil {
		if err := s.NonExistMediaWorkflowIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaWorkflowListResponseBodyMediaWorkflowList struct {
	MediaWorkflow []*QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Repeated"`
}

func (s QueryMediaWorkflowListResponseBodyMediaWorkflowList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowListResponseBodyMediaWorkflowList) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowList) GetMediaWorkflow() []*QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	return s.MediaWorkflow
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowList) SetMediaWorkflow(v []*QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) *QueryMediaWorkflowListResponseBodyMediaWorkflowList {
	s.MediaWorkflow = v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowList) Validate() error {
	if s.MediaWorkflow != nil {
		for _, item := range s.MediaWorkflow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow struct {
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Topology        *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	TriggerMode     *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GetState() *string {
	return s.State
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) SetCreationTime(v string) *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) SetMediaWorkflowId(v string) *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) SetName(v string) *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	s.Name = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) SetState(v string) *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	s.State = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) SetTopology(v string) *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) SetTriggerMode(v string) *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow {
	s.TriggerMode = &v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyMediaWorkflowListMediaWorkflow) Validate() error {
	return dara.Validate(s)
}

type QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds struct {
	MediaWorkflowId []*string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty" type:"Repeated"`
}

func (s QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) GetMediaWorkflowId() []*string {
	return s.MediaWorkflowId
}

func (s *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) SetMediaWorkflowId(v []*string) *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds {
	s.MediaWorkflowId = v
	return s
}

func (s *QueryMediaWorkflowListResponseBodyNonExistMediaWorkflowIds) Validate() error {
	return dara.Validate(s)
}
