// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowId(v string) *UpdateMediaWorkflowRequest
	GetMediaWorkflowId() *string
	SetName(v string) *UpdateMediaWorkflowRequest
	GetName() *string
	SetOwnerAccount(v string) *UpdateMediaWorkflowRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateMediaWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateMediaWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaWorkflowRequest
	GetResourceOwnerId() *int64
	SetTopology(v string) *UpdateMediaWorkflowRequest
	GetTopology() *string
	SetTriggerMode(v string) *UpdateMediaWorkflowRequest
	GetTriggerMode() *string
}

type UpdateMediaWorkflowRequest struct {
	// The ID of the media workflow that you want to update. To obtain the ID of the media workflow, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Workflow Settings*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6307eb0d3f85477882d205aa040f****
	MediaWorkflowId      *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new topology of the media workflow. The value is a JSON object that contains the activity list and activity dependencies.
	//
	// > The Object Storage Service (OSS) bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
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
	Topology    *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s UpdateMediaWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *UpdateMediaWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaWorkflowRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateMediaWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaWorkflowRequest) GetTopology() *string {
	return s.Topology
}

func (s *UpdateMediaWorkflowRequest) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *UpdateMediaWorkflowRequest) SetMediaWorkflowId(v string) *UpdateMediaWorkflowRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetName(v string) *UpdateMediaWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetOwnerAccount(v string) *UpdateMediaWorkflowRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetOwnerId(v int64) *UpdateMediaWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetResourceOwnerAccount(v string) *UpdateMediaWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetResourceOwnerId(v int64) *UpdateMediaWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetTopology(v string) *UpdateMediaWorkflowRequest {
	s.Topology = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) SetTriggerMode(v string) *UpdateMediaWorkflowRequest {
	s.TriggerMode = &v
	return s
}

func (s *UpdateMediaWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
