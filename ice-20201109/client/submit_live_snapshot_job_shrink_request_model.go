// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveSnapshotJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *SubmitLiveSnapshotJobShrinkRequest
	GetCallbackUrl() *string
	SetJobName(v string) *SubmitLiveSnapshotJobShrinkRequest
	GetJobName() *string
	SetSnapshotOutputShrink(v string) *SubmitLiveSnapshotJobShrinkRequest
	GetSnapshotOutputShrink() *string
	SetStreamInputShrink(v string) *SubmitLiveSnapshotJobShrinkRequest
	GetStreamInputShrink() *string
	SetTemplateId(v string) *SubmitLiveSnapshotJobShrinkRequest
	GetTemplateId() *string
}

type SubmitLiveSnapshotJobShrinkRequest struct {
	// The snapshot callback URL.
	//
	// 	- It cannot exceed 255 characters in length.
	//
	// 	- Both HTTP and HTTPS URLs are supported.
	//
	// example:
	//
	// http://www.aliyun.com/snapshot/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The name of the job.
	//
	// 	- It cannot exceed 128 characters in length.
	//
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The information about the output snapshot.
	//
	// This parameter is required.
	SnapshotOutputShrink *string `json:"SnapshotOutput,omitempty" xml:"SnapshotOutput,omitempty"`
	// The information about the input stream.
	//
	// This parameter is required.
	StreamInputShrink *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitLiveSnapshotJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveSnapshotJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveSnapshotJobShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SubmitLiveSnapshotJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *SubmitLiveSnapshotJobShrinkRequest) GetSnapshotOutputShrink() *string {
	return s.SnapshotOutputShrink
}

func (s *SubmitLiveSnapshotJobShrinkRequest) GetStreamInputShrink() *string {
	return s.StreamInputShrink
}

func (s *SubmitLiveSnapshotJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitLiveSnapshotJobShrinkRequest) SetCallbackUrl(v string) *SubmitLiveSnapshotJobShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SubmitLiveSnapshotJobShrinkRequest) SetJobName(v string) *SubmitLiveSnapshotJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *SubmitLiveSnapshotJobShrinkRequest) SetSnapshotOutputShrink(v string) *SubmitLiveSnapshotJobShrinkRequest {
	s.SnapshotOutputShrink = &v
	return s
}

func (s *SubmitLiveSnapshotJobShrinkRequest) SetStreamInputShrink(v string) *SubmitLiveSnapshotJobShrinkRequest {
	s.StreamInputShrink = &v
	return s
}

func (s *SubmitLiveSnapshotJobShrinkRequest) SetTemplateId(v string) *SubmitLiveSnapshotJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitLiveSnapshotJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
