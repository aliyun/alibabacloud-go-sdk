// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaWorkflowTriggerModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflow(v *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) *UpdateMediaWorkflowTriggerModeResponseBody
	GetMediaWorkflow() *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow
	SetRequestId(v string) *UpdateMediaWorkflowTriggerModeResponseBody
	GetRequestId() *string
}

type UpdateMediaWorkflowTriggerModeResponseBody struct {
	// The information about the media workflow.
	MediaWorkflow *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16CD0CDD-457E-420D-9755-8385075A1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaWorkflowTriggerModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowTriggerModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowTriggerModeResponseBody) GetMediaWorkflow() *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	return s.MediaWorkflow
}

func (s *UpdateMediaWorkflowTriggerModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaWorkflowTriggerModeResponseBody) SetMediaWorkflow(v *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) *UpdateMediaWorkflowTriggerModeResponseBody {
	s.MediaWorkflow = v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBody) SetRequestId(v string) *UpdateMediaWorkflowTriggerModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBody) Validate() error {
	if s.MediaWorkflow != nil {
		if err := s.MediaWorkflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow struct {
	// The time when the media workflow was created.
	//
	// example:
	//
	// 2016-04-01T05:29:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow.
	//
	// example:
	//
	// e00732b977da427d9177a4dee646****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// example-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the media workflow. Valid values:
	//
	// 	- **Inactive**: The media workflow is disabled.
	//
	// 	- **Active**: The media workflow is enabled.
	//
	// example:
	//
	// Inactive
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
	// NotInAuto
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GetState() *string {
	return s.State
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) SetCreationTime(v string) *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) SetMediaWorkflowId(v string) *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) SetName(v string) *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	s.Name = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) SetState(v string) *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	s.State = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) SetTopology(v string) *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) SetTriggerMode(v string) *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow {
	s.TriggerMode = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeResponseBodyMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
