// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateMediaWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflow(v *DeactivateMediaWorkflowResponseBodyMediaWorkflow) *DeactivateMediaWorkflowResponseBody
	GetMediaWorkflow() *DeactivateMediaWorkflowResponseBodyMediaWorkflow
	SetRequestId(v string) *DeactivateMediaWorkflowResponseBody
	GetRequestId() *string
}

type DeactivateMediaWorkflowResponseBody struct {
	// The topology of the media workflow.
	MediaWorkflow *DeactivateMediaWorkflowResponseBodyMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Struct"`
	// The name of the media workflow that is deactivated.
	//
	// example:
	//
	// 16CD0CDD-457E-420D-9755-8385075A1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactivateMediaWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateMediaWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateMediaWorkflowResponseBody) GetMediaWorkflow() *DeactivateMediaWorkflowResponseBodyMediaWorkflow {
	return s.MediaWorkflow
}

func (s *DeactivateMediaWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateMediaWorkflowResponseBody) SetMediaWorkflow(v *DeactivateMediaWorkflowResponseBodyMediaWorkflow) *DeactivateMediaWorkflowResponseBody {
	s.MediaWorkflow = v
	return s
}

func (s *DeactivateMediaWorkflowResponseBody) SetRequestId(v string) *DeactivateMediaWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateMediaWorkflowResponseBody) Validate() error {
	if s.MediaWorkflow != nil {
		if err := s.MediaWorkflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeactivateMediaWorkflowResponseBodyMediaWorkflow struct {
	// 	- After you deactivate a media workflow, you can modify the workflow information.
	//
	// 	- After you delete or deactivate a media workflow, the workflow cannot be used. In this case, the workflow is not automatically triggered when you upload a file to the bucket specified by the workflow.
	//
	// ## Limits on QPS
	//
	// You can call this operation up to 100 times per second. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limits](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
	//
	// example:
	//
	// 2016-04-01T05:29:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow that you want to deactivate. To obtain the ID of the media workflow, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Workflow Settings*	- in the left-side navigation pane.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e9181d4****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The details of the media workflow.
	//
	// example:
	//
	// example-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The topology of the media workflow.The status of the media workflow. The value is **Inactive**.
	//
	// example:
	//
	// Inactive
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the media workflow. The value is **Inactive**.
	//
	// example:
	//
	// {mediaworkflow","State":"Active","Topology":"{\\"Activities\\":{\\"Act-Start\\":{\\"Parameters\\":{\\"PipelineId\\":\\"130266f58161436a80bf07cb12c8****\\",\\"InputFile\\":\\"{\\\\\\"Bucket\\\\\\": \\\\\\"example-bucket-****\\\\\\",\\\\\\"Location\\\\\\": \\\\\\"cn-shanghai\\\\\\"}\\"},\\"Type\\":\\"Start\\"},\\"Act-Report\\":{\\"Parameters\\":{},\\"Type\\":\\"Report\\"},\\"Act-Transcode-M3U8\\":{\\"Parameters\\":{\\"Outputs\\":\\"[{\\\\\\"Object\\\\\\":\\\\\\"transcode/{ObjectPrefix}{FileName}\\\\\\",\\\\\\"TemplateId\\\\\\": \\\\\\"957d1719ee85ed6527b90cf62726****\\\\\\"}]\\",\\"OutputBucket\\":\\"example-bucket-****\\",\\"OutputLocation\\":\\"cn-shanghai\\"},\\"Type\\":\\"Transcode\\"}},\\"Dependencies\\":{\\"Act-Start\\":[\\"Act-Transcode-M3U8\\"],\\"Act-Report\\":[],\\"Act-Transcode-M3U8\\":[\\"Act-Report\\"]}}","MediaWorkflowId":"93ab850b4f6f44eab54b6e91d24d****"}]},"RequestId":"16CD0CDD-457E-420D-9755-8385075A1234"}
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
}

func (s DeactivateMediaWorkflowResponseBodyMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s DeactivateMediaWorkflowResponseBodyMediaWorkflow) GoString() string {
	return s.String()
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) GetState() *string {
	return s.State
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) SetCreationTime(v string) *DeactivateMediaWorkflowResponseBodyMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) SetMediaWorkflowId(v string) *DeactivateMediaWorkflowResponseBodyMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) SetName(v string) *DeactivateMediaWorkflowResponseBodyMediaWorkflow {
	s.Name = &v
	return s
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) SetState(v string) *DeactivateMediaWorkflowResponseBodyMediaWorkflow {
	s.State = &v
	return s
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) SetTopology(v string) *DeactivateMediaWorkflowResponseBodyMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *DeactivateMediaWorkflowResponseBodyMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
