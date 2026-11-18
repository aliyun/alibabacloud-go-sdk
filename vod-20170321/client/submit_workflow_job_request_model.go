// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitWorkflowJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraInputs(v string) *SubmitWorkflowJobRequest
	GetExtraInputs() *string
	SetMediaId(v string) *SubmitWorkflowJobRequest
	GetMediaId() *string
	SetWorkflowId(v string) *SubmitWorkflowJobRequest
	GetWorkflowId() *string
}

type SubmitWorkflowJobRequest struct {
	ExtraInputs *string `json:"ExtraInputs,omitempty" xml:"ExtraInputs,omitempty"`
	// The media asset ID, which is the video ID. You can obtain the ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Assets*	- > **Audio/Video*	- to view the video ID.
	//
	// - When you call the [Get the upload URL and credential for a video](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the video ID is the value of the VideoId response parameter.
	//
	// - After a video is uploaded, you can call the [Search for media information](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID. The video ID is the value of the VideoId response parameter.
	//
	// example:
	//
	// 058b39e75269da42b08f00459****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The workflow ID. You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing Configuration*	- > **Workflow Management*	- to view the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34d577eade633860bdf1237****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s SubmitWorkflowJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitWorkflowJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitWorkflowJobRequest) GetExtraInputs() *string {
	return s.ExtraInputs
}

func (s *SubmitWorkflowJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitWorkflowJobRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *SubmitWorkflowJobRequest) SetExtraInputs(v string) *SubmitWorkflowJobRequest {
	s.ExtraInputs = &v
	return s
}

func (s *SubmitWorkflowJobRequest) SetMediaId(v string) *SubmitWorkflowJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitWorkflowJobRequest) SetWorkflowId(v string) *SubmitWorkflowJobRequest {
	s.WorkflowId = &v
	return s
}

func (s *SubmitWorkflowJobRequest) Validate() error {
	return dara.Validate(s)
}
