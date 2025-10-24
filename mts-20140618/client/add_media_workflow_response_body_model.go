// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflow(v *AddMediaWorkflowResponseBodyMediaWorkflow) *AddMediaWorkflowResponseBody
	GetMediaWorkflow() *AddMediaWorkflowResponseBodyMediaWorkflow
	SetRequestId(v string) *AddMediaWorkflowResponseBody
	GetRequestId() *string
}

type AddMediaWorkflowResponseBody struct {
	// The information about the media workflow.
	MediaWorkflow *AddMediaWorkflowResponseBodyMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F1D21261-ADB9-406A-1234-491382139D59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMediaWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaWorkflowResponseBody) GetMediaWorkflow() *AddMediaWorkflowResponseBodyMediaWorkflow {
	return s.MediaWorkflow
}

func (s *AddMediaWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaWorkflowResponseBody) SetMediaWorkflow(v *AddMediaWorkflowResponseBodyMediaWorkflow) *AddMediaWorkflowResponseBody {
	s.MediaWorkflow = v
	return s
}

func (s *AddMediaWorkflowResponseBody) SetRequestId(v string) *AddMediaWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaWorkflowResponseBody) Validate() error {
	if s.MediaWorkflow != nil {
		if err := s.MediaWorkflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMediaWorkflowResponseBodyMediaWorkflow struct {
	// The time when the media workflow was created.
	//
	// example:
	//
	// 016-04-01T05:29:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow. We recommend that you keep this ID for later operations on this workflow.
	//
	// example:
	//
	// e00732b977da427d9177a4deb1aa****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// mediaworkflow-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the media workflow. By default, the created workflow is in the **Active*	- state.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The topology of the media workflow. The value is a JSON object that contains the activities and activity dependencies.
	//
	// example:
	//
	// {mediaworkflow","State":"Active","Topology":"{\\"Activities\\":{\\"Act-Start\\":{\\"Parameters\\":{\\"PipelineId\\":\\"130266f58161436a80bf07cb12c8****\\",\\"InputFile\\":\\"{\\\\\\"Bucket\\\\\\": \\\\\\"example-bucket-****\\\\\\",\\\\\\"Location\\\\\\": \\\\\\"cn-shanghai\\\\\\"}\\"},\\"Type\\":\\"Start\\"},\\"Act-Report\\":{\\"Parameters\\":{},\\"Type\\":\\"Report\\"},\\"Act-Transcode-M3U8\\":{\\"Parameters\\":{\\"Outputs\\":\\"[{\\\\\\"Object\\\\\\":\\\\\\"transcode/{ObjectPrefix}{FileName}\\\\\\",\\\\\\"TemplateId\\\\\\": \\\\\\"957d1719ee85ed6527b90cf62726****\\\\\\"}]\\",\\"OutputBucket\\":\\"example-bucket-****\\",\\"OutputLocation\\":\\"cn-shanghai\\"},\\"Type\\":\\"Transcode\\"}},\\"Dependencies\\":{\\"Act-Start\\":[\\"Act-Transcode-M3U8\\"],\\"Act-Report\\":[],\\"Act-Transcode-M3U8\\":[\\"Act-Report\\"]}}","MediaWorkflowId":"93ab850b4f6f44eab54b6e91d24d****"}]},"RequestId":"16CD0CDD-457E-420D-9755-8385075A1234"}
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	// The triggering mode of the media workflow. Valid values:
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

func (s AddMediaWorkflowResponseBodyMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s AddMediaWorkflowResponseBodyMediaWorkflow) GoString() string {
	return s.String()
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) GetState() *string {
	return s.State
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) SetCreationTime(v string) *AddMediaWorkflowResponseBodyMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) SetMediaWorkflowId(v string) *AddMediaWorkflowResponseBodyMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) SetName(v string) *AddMediaWorkflowResponseBodyMediaWorkflow {
	s.Name = &v
	return s
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) SetState(v string) *AddMediaWorkflowResponseBodyMediaWorkflow {
	s.State = &v
	return s
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) SetTopology(v string) *AddMediaWorkflowResponseBodyMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) SetTriggerMode(v string) *AddMediaWorkflowResponseBodyMediaWorkflow {
	s.TriggerMode = &v
	return s
}

func (s *AddMediaWorkflowResponseBodyMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
