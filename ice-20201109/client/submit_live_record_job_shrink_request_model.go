// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveRecordJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SubmitLiveRecordJobShrinkRequest
	GetName() *string
	SetNotifyUrl(v string) *SubmitLiveRecordJobShrinkRequest
	GetNotifyUrl() *string
	SetRecordOutputShrink(v string) *SubmitLiveRecordJobShrinkRequest
	GetRecordOutputShrink() *string
	SetStreamInputShrink(v string) *SubmitLiveRecordJobShrinkRequest
	GetStreamInputShrink() *string
	SetTemplateId(v string) *SubmitLiveRecordJobShrinkRequest
	GetTemplateId() *string
}

type SubmitLiveRecordJobShrinkRequest struct {
	// The name of the recording job.
	//
	// This parameter is required.
	//
	// example:
	//
	// live stream record 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// https://example.com/imsnotify
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The storage address of the recording.
	//
	// This parameter is required.
	RecordOutputShrink *string `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty"`
	// The URL of the live stream.
	//
	// This parameter is required.
	StreamInputShrink *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The ID of the recording template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitLiveRecordJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveRecordJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveRecordJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitLiveRecordJobShrinkRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitLiveRecordJobShrinkRequest) GetRecordOutputShrink() *string {
	return s.RecordOutputShrink
}

func (s *SubmitLiveRecordJobShrinkRequest) GetStreamInputShrink() *string {
	return s.StreamInputShrink
}

func (s *SubmitLiveRecordJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitLiveRecordJobShrinkRequest) SetName(v string) *SubmitLiveRecordJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitLiveRecordJobShrinkRequest) SetNotifyUrl(v string) *SubmitLiveRecordJobShrinkRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitLiveRecordJobShrinkRequest) SetRecordOutputShrink(v string) *SubmitLiveRecordJobShrinkRequest {
	s.RecordOutputShrink = &v
	return s
}

func (s *SubmitLiveRecordJobShrinkRequest) SetStreamInputShrink(v string) *SubmitLiveRecordJobShrinkRequest {
	s.StreamInputShrink = &v
	return s
}

func (s *SubmitLiveRecordJobShrinkRequest) SetTemplateId(v string) *SubmitLiveRecordJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitLiveRecordJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
