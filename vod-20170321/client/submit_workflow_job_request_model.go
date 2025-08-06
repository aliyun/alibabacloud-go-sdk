// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitWorkflowJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *SubmitWorkflowJobRequest
	GetMediaId() *string
	SetWorkflowId(v string) *SubmitWorkflowJobRequest
	GetWorkflowId() *string
}

type SubmitWorkflowJobRequest struct {
	// The ID of the media file. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the audio or video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of the VideoId parameter when you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to upload media files.
	//
	// 	- Obtain the value of the VideoId parameter when you call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation after you upload media files.
	//
	// example:
	//
	// 058b39e75269da42b08f00459****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the workflow. To view the ID of the workflow, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Processing*	- > **Workflows**.
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

func (s *SubmitWorkflowJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitWorkflowJobRequest) GetWorkflowId() *string {
	return s.WorkflowId
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
