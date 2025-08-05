// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventStreamingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateEventStreamingShrinkRequest
	GetDescription() *string
	SetEventStreamingName(v string) *CreateEventStreamingShrinkRequest
	GetEventStreamingName() *string
	SetFilterPattern(v string) *CreateEventStreamingShrinkRequest
	GetFilterPattern() *string
	SetRunOptionsShrink(v string) *CreateEventStreamingShrinkRequest
	GetRunOptionsShrink() *string
	SetSinkShrink(v string) *CreateEventStreamingShrinkRequest
	GetSinkShrink() *string
	SetSourceShrink(v string) *CreateEventStreamingShrinkRequest
	GetSourceShrink() *string
	SetTags(v []*CreateEventStreamingShrinkRequestTags) *CreateEventStreamingShrinkRequest
	GetTags() []*CreateEventStreamingShrinkRequestTags
	SetTransformsShrink(v string) *CreateEventStreamingShrinkRequest
	GetTransformsShrink() *string
}

type CreateEventStreamingShrinkRequest struct {
	// The description of the event stream.
	//
	// example:
	//
	// rocketmq2mns
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// myeventstreaming
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	//
	// This parameter is required.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptionsShrink *string `json:"RunOptions,omitempty" xml:"RunOptions,omitempty"`
	// The event target. You must and can specify only one event target.
	//
	// This parameter is required.
	SinkShrink *string `json:"Sink,omitempty" xml:"Sink,omitempty"`
	// The event provider, which is also known as the event source. You must and can specify only one event source.
	//
	// This parameter is required.
	SourceShrink     *string                                  `json:"Source,omitempty" xml:"Source,omitempty"`
	Tags             []*CreateEventStreamingShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TransformsShrink *string                                  `json:"Transforms,omitempty" xml:"Transforms,omitempty"`
}

func (s CreateEventStreamingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventStreamingShrinkRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *CreateEventStreamingShrinkRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *CreateEventStreamingShrinkRequest) GetRunOptionsShrink() *string {
	return s.RunOptionsShrink
}

func (s *CreateEventStreamingShrinkRequest) GetSinkShrink() *string {
	return s.SinkShrink
}

func (s *CreateEventStreamingShrinkRequest) GetSourceShrink() *string {
	return s.SourceShrink
}

func (s *CreateEventStreamingShrinkRequest) GetTags() []*CreateEventStreamingShrinkRequestTags {
	return s.Tags
}

func (s *CreateEventStreamingShrinkRequest) GetTransformsShrink() *string {
	return s.TransformsShrink
}

func (s *CreateEventStreamingShrinkRequest) SetDescription(v string) *CreateEventStreamingShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetEventStreamingName(v string) *CreateEventStreamingShrinkRequest {
	s.EventStreamingName = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetFilterPattern(v string) *CreateEventStreamingShrinkRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetRunOptionsShrink(v string) *CreateEventStreamingShrinkRequest {
	s.RunOptionsShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetSinkShrink(v string) *CreateEventStreamingShrinkRequest {
	s.SinkShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetSourceShrink(v string) *CreateEventStreamingShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetTags(v []*CreateEventStreamingShrinkRequestTags) *CreateEventStreamingShrinkRequest {
	s.Tags = v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetTransformsShrink(v string) *CreateEventStreamingShrinkRequest {
	s.TransformsShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingShrinkRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateEventStreamingShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateEventStreamingShrinkRequestTags) SetKey(v string) *CreateEventStreamingShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateEventStreamingShrinkRequestTags) SetValue(v string) *CreateEventStreamingShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *CreateEventStreamingShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
