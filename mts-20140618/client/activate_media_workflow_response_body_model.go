// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateMediaWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflow(v *ActivateMediaWorkflowResponseBodyMediaWorkflow) *ActivateMediaWorkflowResponseBody
	GetMediaWorkflow() *ActivateMediaWorkflowResponseBodyMediaWorkflow
	SetRequestId(v string) *ActivateMediaWorkflowResponseBody
	GetRequestId() *string
}

type ActivateMediaWorkflowResponseBody struct {
	// The details of the media workflow.
	MediaWorkflow *ActivateMediaWorkflowResponseBodyMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A1326BD4-30B1-4CB6-Q123-3330B877B0D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateMediaWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateMediaWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateMediaWorkflowResponseBody) GetMediaWorkflow() *ActivateMediaWorkflowResponseBodyMediaWorkflow {
	return s.MediaWorkflow
}

func (s *ActivateMediaWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateMediaWorkflowResponseBody) SetMediaWorkflow(v *ActivateMediaWorkflowResponseBodyMediaWorkflow) *ActivateMediaWorkflowResponseBody {
	s.MediaWorkflow = v
	return s
}

func (s *ActivateMediaWorkflowResponseBody) SetRequestId(v string) *ActivateMediaWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateMediaWorkflowResponseBody) Validate() error {
	if s.MediaWorkflow != nil {
		if err := s.MediaWorkflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActivateMediaWorkflowResponseBodyMediaWorkflow struct {
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
	// 93ab850b4f6f44eab54b6e9181d4****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// mediaworkflow-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the media workflow. The value is **Active**.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The topology of the media workflow.
	//
	// example:
	//
	// {\\"Activities\\":{\\"Act-Start\\":{\\"Parameters\\":{\\"PipelineId\\":\\"130266f58161436a80bf07cb12c8****\\",\\"InputFile\\":\\"{\\\\\\"Bucket\\\\\\": \\\\\\"example\\\\\\",\\\\\\"Location\\\\\\": \\\\\\"oss-cn-hangzhou\\\\\\"}\\"},\\"Type\\":\\"Start\\"},\\"Act-Report\\":{\\"Parameters\\":{},\\"Type\\":\\"Report\\"},\\"Act-Transcode-M3U8\\":{\\"Parameters\\":{\\"Outputs\\":\\"[{\\\\\\"OutputObject\\\\\\":\\\\\\"transcode%2F%7BObjectPrefix%7D%7BFileName%7D\\\\\\",\\\\\\"TemplateId\\\\\\": \\\\\\"957d1719ee85ed6527b90cf62726****\\\\\\"}]\\",\\"OutputBucket\\":\\"panda-vod-hls\\",\\"OutputLocation\\":\\"oss-cn-hangzhou\\"},\\"Type\\":\\"Transcode\\"}},\\"Dependencies\\":{\\"Act-Start\\":[\\"Act-Transcode-M3U8\\"],\\"Act-Report\\":[],\\"Act-Transcode-M3U8\\":[\\"Act-Report\\"]}}
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
}

func (s ActivateMediaWorkflowResponseBodyMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s ActivateMediaWorkflowResponseBodyMediaWorkflow) GoString() string {
	return s.String()
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) GetState() *string {
	return s.State
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) SetCreationTime(v string) *ActivateMediaWorkflowResponseBodyMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) SetMediaWorkflowId(v string) *ActivateMediaWorkflowResponseBodyMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) SetName(v string) *ActivateMediaWorkflowResponseBodyMediaWorkflow {
	s.Name = &v
	return s
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) SetState(v string) *ActivateMediaWorkflowResponseBodyMediaWorkflow {
	s.State = &v
	return s
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) SetTopology(v string) *ActivateMediaWorkflowResponseBodyMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *ActivateMediaWorkflowResponseBodyMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
