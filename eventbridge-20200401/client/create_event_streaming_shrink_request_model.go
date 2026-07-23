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
	SetMetadata(v string) *CreateEventStreamingShrinkRequest
	GetMetadata() *string
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
	// The event filtering rule. If not specified, all events are matched.
	//
	// example:
	//
	// {
	//
	// "source": [
	//
	// {
	//
	// "prefix": "acs:mns"
	//
	// }
	//
	// ],
	//
	// "type": [
	//
	// {
	//
	// "prefix": "mns:Queue"
	//
	// }
	//
	// ],
	//
	// "subject": [
	//
	// {
	//
	// "prefix": "acs:mns:cn-hangzhou:123456789098****:queues/zeus"
	//
	// }
	//
	// ]
	//
	// }
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	Metadata      *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The runtime environment parameters.
	RunOptionsShrink *string `json:"RunOptions,omitempty" xml:"RunOptions,omitempty"`
	// The event target. You must select exactly one Sink type.
	SinkShrink *string `json:"Sink,omitempty" xml:"Sink,omitempty"`
	// The event provider. You must select exactly one Source type.
	SourceShrink *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The tag list. A maximum of 20 items are supported.
	Tags []*CreateEventStreamingShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Transform-related configurations.
	TransformsShrink *string `json:"Transforms,omitempty" xml:"Transforms,omitempty"`
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

func (s *CreateEventStreamingShrinkRequest) GetMetadata() *string {
	return s.Metadata
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

func (s *CreateEventStreamingShrinkRequest) SetMetadata(v string) *CreateEventStreamingShrinkRequest {
	s.Metadata = &v
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEventStreamingShrinkRequestTags struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
