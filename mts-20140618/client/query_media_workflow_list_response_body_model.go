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
	// The media workflows.
	MediaWorkflowList *QueryMediaWorkflowListResponseBodyMediaWorkflowList `json:"MediaWorkflowList,omitempty" xml:"MediaWorkflowList,omitempty" type:"Struct"`
	// The workflows that do not exist.
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
	// The time when the media workflow was created.
	//
	// example:
	//
	// 2016-04-01T05:29:38Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e9181d4****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// example-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the media workflow. Valid values:
	//
	// 	- **Inactive**: The media workflow was deactivated.
	//
	// 	- **Active**: The media workflow was activated.
	//
	// 	- **Deleted**: The media workflow was deleted.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The topology of the media workflow.
	//
	// example:
	//
	// {mediaworkflow","State":"Active","Topology":"{\\"Activities\\":{\\"Act-Start\\":{\\"Parameters\\":{\\"PipelineId\\":\\"130266f58161436a80bf07cb12c8****\\",\\"InputFile\\":\\"{\\\\\\"Bucket\\\\\\": \\\\\\"example-bucket-****\\\\\\",\\\\\\"Location\\\\\\": \\\\\\"cn-shanghai\\\\\\"}\\"},\\"Type\\":\\"Start\\"},\\"Act-Report\\":{\\"Parameters\\":{},\\"Type\\":\\"Report\\"},\\"Act-Transcode-M3U8\\":{\\"Parameters\\":{\\"Outputs\\":\\"[{\\\\\\"Object\\\\\\":\\\\\\"transcode/{ObjectPrefix}{FileName}\\\\\\",\\\\\\"TemplateId\\\\\\": \\\\\\"957d1719ee85ed6527b90cf62726****\\\\\\"}]\\",\\"OutputBucket\\":\\"example-bucket-****\\",\\"OutputLocation\\":\\"cn-shanghai\\"},\\"Type\\":\\"Transcode\\"}},\\"Dependencies\\":{\\"Act-Start\\":[\\"Act-Transcode-M3U8\\"],\\"Act-Report\\":[],\\"Act-Transcode-M3U8\\":[\\"Act-Report\\"]}}","MediaWorkflowId":"93ab850b4f6f44eab54b6e91d24d****"}]},"RequestId":"16CD0CDD-457E-420D-9755-8385075A1234"}
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	// The trigger mode of the media workflow. Valid values:
	//
	// 	- **OssAutoTrigger**: The media workflow is automatically triggered.
	//
	// 	- **NotInAuto**: The media workflow is not automatically triggered.
	//
	// example:
	//
	// OssAutoTrigger
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
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
