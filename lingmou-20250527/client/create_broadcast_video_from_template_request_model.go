// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastVideoFromTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateBroadcastVideoFromTemplateRequest
	GetName() *string
	SetTemplateId(v string) *CreateBroadcastVideoFromTemplateRequest
	GetTemplateId() *string
	SetVariables(v []*TemplateVariable) *CreateBroadcastVideoFromTemplateRequest
	GetVariables() []*TemplateVariable
	SetVideoOptions(v *CreateBroadcastVideoFromTemplateRequestVideoOptions) *CreateBroadcastVideoFromTemplateRequest
	GetVideoOptions() *CreateBroadcastVideoFromTemplateRequestVideoOptions
}

type CreateBroadcastVideoFromTemplateRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// BS1b2WNnRMu4ouRzT4clY9Jhg
	TemplateId   *string                                              `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Variables    []*TemplateVariable                                  `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
	VideoOptions *CreateBroadcastVideoFromTemplateRequestVideoOptions `json:"videoOptions,omitempty" xml:"videoOptions,omitempty" type:"Struct"`
}

func (s CreateBroadcastVideoFromTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastVideoFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateBroadcastVideoFromTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateBroadcastVideoFromTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateBroadcastVideoFromTemplateRequest) GetVariables() []*TemplateVariable {
	return s.Variables
}

func (s *CreateBroadcastVideoFromTemplateRequest) GetVideoOptions() *CreateBroadcastVideoFromTemplateRequestVideoOptions {
	return s.VideoOptions
}

func (s *CreateBroadcastVideoFromTemplateRequest) SetName(v string) *CreateBroadcastVideoFromTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequest) SetTemplateId(v string) *CreateBroadcastVideoFromTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequest) SetVariables(v []*TemplateVariable) *CreateBroadcastVideoFromTemplateRequest {
	s.Variables = v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequest) SetVideoOptions(v *CreateBroadcastVideoFromTemplateRequestVideoOptions) *CreateBroadcastVideoFromTemplateRequest {
	s.VideoOptions = v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequest) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoOptions != nil {
		if err := s.VideoOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBroadcastVideoFromTemplateRequestVideoOptions struct {
	// example:
	//
	// 30
	Fps *int32 `json:"fps,omitempty" xml:"fps,omitempty"`
	// example:
	//
	// 720p
	Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// example:
	//
	// True
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s CreateBroadcastVideoFromTemplateRequestVideoOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastVideoFromTemplateRequestVideoOptions) GoString() string {
	return s.String()
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) GetFps() *int32 {
	return s.Fps
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) GetResolution() *string {
	return s.Resolution
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) GetWatermark() *bool {
	return s.Watermark
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) SetFps(v int32) *CreateBroadcastVideoFromTemplateRequestVideoOptions {
	s.Fps = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) SetResolution(v string) *CreateBroadcastVideoFromTemplateRequestVideoOptions {
	s.Resolution = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) SetWatermark(v bool) *CreateBroadcastVideoFromTemplateRequestVideoOptions {
	s.Watermark = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateRequestVideoOptions) Validate() error {
	return dara.Validate(s)
}
