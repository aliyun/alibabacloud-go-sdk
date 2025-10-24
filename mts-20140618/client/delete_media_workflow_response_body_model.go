// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflow(v *DeleteMediaWorkflowResponseBodyMediaWorkflow) *DeleteMediaWorkflowResponseBody
	GetMediaWorkflow() *DeleteMediaWorkflowResponseBodyMediaWorkflow
	SetRequestId(v string) *DeleteMediaWorkflowResponseBody
	GetRequestId() *string
}

type DeleteMediaWorkflowResponseBody struct {
	// The information about the media workflow.
	MediaWorkflow *DeleteMediaWorkflowResponseBodyMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7D752035-97DA-54E5-88E2-E8405EEA4394
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaWorkflowResponseBody) GetMediaWorkflow() *DeleteMediaWorkflowResponseBodyMediaWorkflow {
	return s.MediaWorkflow
}

func (s *DeleteMediaWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaWorkflowResponseBody) SetMediaWorkflow(v *DeleteMediaWorkflowResponseBodyMediaWorkflow) *DeleteMediaWorkflowResponseBody {
	s.MediaWorkflow = v
	return s
}

func (s *DeleteMediaWorkflowResponseBody) SetRequestId(v string) *DeleteMediaWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaWorkflowResponseBody) Validate() error {
	if s.MediaWorkflow != nil {
		if err := s.MediaWorkflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMediaWorkflowResponseBodyMediaWorkflow struct {
	// The time when the media workflow was created.
	//
	// example:
	//
	// 2016-04-01T05:29:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow that is deleted.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e9181d4****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow that is deleted.
	//
	// example:
	//
	// example-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the media workflow. The value is **Deleted**.
	//
	// example:
	//
	// Deleted
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The topology of the media workflow.
	//
	// example:
	//
	// {mediaworkflow","State":"Active","Topology":"{\\"Activities\\":{\\"Act-Start\\":{\\"Parameters\\":{\\"PipelineId\\":\\"130266f58161436a80bf07cb12c8****\\",\\"InputFile\\":\\"{\\\\\\"Bucket\\\\\\": \\\\\\"example-bucket-****\\\\\\",\\\\\\"Location\\\\\\": \\\\\\"cn-shanghai\\\\\\"}\\"},\\"Type\\":\\"Start\\"},\\"Act-Report\\":{\\"Parameters\\":{},\\"Type\\":\\"Report\\"},\\"Act-Transcode-M3U8\\":{\\"Parameters\\":{\\"Outputs\\":\\"[{\\\\\\"Object\\\\\\":\\\\\\"transcode/{ObjectPrefix}{FileName}\\\\\\",\\\\\\"TemplateId\\\\\\": \\\\\\"957d1719ee85ed6527b90cf62726****\\\\\\"}]\\",\\"OutputBucket\\":\\"example-bucket-****\\",\\"OutputLocation\\":\\"cn-shanghai\\"},\\"Type\\":\\"Transcode\\"}},\\"Dependencies\\":{\\"Act-Start\\":[\\"Act-Transcode-M3U8\\"],\\"Act-Report\\":[],\\"Act-Transcode-M3U8\\":[\\"Act-Report\\"]}}","MediaWorkflowId":"93ab850b4f6f44eab54b6e91d24d****"}]},"RequestId":"16CD0CDD-457E-420D-9755-8385075A1234"}
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
}

func (s DeleteMediaWorkflowResponseBodyMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaWorkflowResponseBodyMediaWorkflow) GoString() string {
	return s.String()
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) GetState() *string {
	return s.State
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) SetCreationTime(v string) *DeleteMediaWorkflowResponseBodyMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) SetMediaWorkflowId(v string) *DeleteMediaWorkflowResponseBodyMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) SetName(v string) *DeleteMediaWorkflowResponseBodyMediaWorkflow {
	s.Name = &v
	return s
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) SetState(v string) *DeleteMediaWorkflowResponseBodyMediaWorkflow {
	s.State = &v
	return s
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) SetTopology(v string) *DeleteMediaWorkflowResponseBodyMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *DeleteMediaWorkflowResponseBodyMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
