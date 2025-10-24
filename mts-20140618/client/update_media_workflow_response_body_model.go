// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflow(v *UpdateMediaWorkflowResponseBodyMediaWorkflow) *UpdateMediaWorkflowResponseBody
	GetMediaWorkflow() *UpdateMediaWorkflowResponseBodyMediaWorkflow
	SetRequestId(v string) *UpdateMediaWorkflowResponseBody
	GetRequestId() *string
}

type UpdateMediaWorkflowResponseBody struct {
	// The detailed information about the media workflow.
	MediaWorkflow *UpdateMediaWorkflowResponseBodyMediaWorkflow `json:"MediaWorkflow,omitempty" xml:"MediaWorkflow,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7D752035-97DA-54E5-88E2-E8405EEA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowResponseBody) GetMediaWorkflow() *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	return s.MediaWorkflow
}

func (s *UpdateMediaWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaWorkflowResponseBody) SetMediaWorkflow(v *UpdateMediaWorkflowResponseBodyMediaWorkflow) *UpdateMediaWorkflowResponseBody {
	s.MediaWorkflow = v
	return s
}

func (s *UpdateMediaWorkflowResponseBody) SetRequestId(v string) *UpdateMediaWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBody) Validate() error {
	if s.MediaWorkflow != nil {
		if err := s.MediaWorkflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaWorkflowResponseBodyMediaWorkflow struct {
	// The time when the media workflow was created.
	//
	// example:
	//
	// 2016-04-01T05:29:38Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media workflow that is updated.
	//
	// example:
	//
	// 6307eb0d3f85477882d205aa040f****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// examp-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the media workflow. Valid values:
	//
	// 	- **Inactive**: The media workflow is disabled.
	//
	// 	- **Active**: The media workflow is enabled.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The topology of the media workflow.
	//
	// example:
	//
	// {
	//
	//       "Activities": {
	//
	//             "Act-Start": {
	//
	//                   "Parameters": {
	//
	//                         "PipelineId": "130266f58161436a80bf07cb12c8****",
	//
	//                         "InputFile": "{\\"Bucket\\": \\"example-bucket-****\\",\\"Location\\": \\"cn-shanghai\\"}"
	//
	//                   },
	//
	//                   "Type": "Start"
	//
	//             },
	//
	//             "Act-Report": {
	//
	//                   "Parameters": {},
	//
	//                   "Type": "Report"
	//
	//             },
	//
	//             "Act-Transcode-M3U8": {
	//
	//                   "Parameters": {
	//
	//                         "Outputs": "[{\\"Object\\":\\"transcode/{ObjectPrefix}{FileName}\\",\\"TemplateId\\": \\"957d1719ee85ed6527b90cf62726****\\"}]",
	//
	//                         "OutputBucket": "example-bucket-****",
	//
	//                         "OutputLocation": "cn-shanghai"
	//
	//                   },
	//
	//                   "Type": "Transcode"
	//
	//             }
	//
	//       },
	//
	//       "Dependencies": {
	//
	//             "Act-Start": [
	//
	//                   "Act-Transcode-M3U8"
	//
	//             ],
	//
	//             "Act-Report": [],
	//
	//             "Act-Transcode-M3U8": [
	//
	//                   "Act-Report"
	//
	//             ]
	//
	//       }
	//
	// }
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

func (s UpdateMediaWorkflowResponseBodyMediaWorkflow) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowResponseBodyMediaWorkflow) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) GetCreationTime() *string {
	return s.CreationTime
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) GetName() *string {
	return s.Name
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) GetState() *string {
	return s.State
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) GetTopology() *string {
	return s.Topology
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) SetCreationTime(v string) *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	s.CreationTime = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) SetMediaWorkflowId(v string) *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	s.MediaWorkflowId = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) SetName(v string) *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	s.Name = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) SetState(v string) *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	s.State = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) SetTopology(v string) *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	s.Topology = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) SetTriggerMode(v string) *UpdateMediaWorkflowResponseBodyMediaWorkflow {
	s.TriggerMode = &v
	return s
}

func (s *UpdateMediaWorkflowResponseBodyMediaWorkflow) Validate() error {
	return dara.Validate(s)
}
