// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateEventStreamingShrinkRequest
	GetDescription() *string
	SetEventStreamingName(v string) *UpdateEventStreamingShrinkRequest
	GetEventStreamingName() *string
	SetFilterPattern(v string) *UpdateEventStreamingShrinkRequest
	GetFilterPattern() *string
	SetRunOptionsShrink(v string) *UpdateEventStreamingShrinkRequest
	GetRunOptionsShrink() *string
	SetSinkShrink(v string) *UpdateEventStreamingShrinkRequest
	GetSinkShrink() *string
	SetSourceShrink(v string) *UpdateEventStreamingShrinkRequest
	GetSourceShrink() *string
	SetTransformsShrink(v string) *UpdateEventStreamingShrinkRequest
	GetTransformsShrink() *string
}

type UpdateEventStreamingShrinkRequest struct {
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
	//
	// example:
	//
	// {
	//
	//     "source": [
	//
	//         {
	//
	//             "prefix": "acs:mns"
	//
	//         }
	//
	//     ],
	//
	//     "type": [
	//
	//         {
	//
	//             "prefix": "mns:Queue"
	//
	//         }
	//
	//     ],
	//
	//     "subject": [
	//
	//         {
	//
	//             "prefix": "acs:mns:cn-hangzhou:123456789098****:queues/zeus"
	//
	//         }
	//
	//     ]
	//
	// }
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
	SourceShrink     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TransformsShrink *string `json:"Transforms,omitempty" xml:"Transforms,omitempty"`
}

func (s UpdateEventStreamingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEventStreamingShrinkRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *UpdateEventStreamingShrinkRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *UpdateEventStreamingShrinkRequest) GetRunOptionsShrink() *string {
	return s.RunOptionsShrink
}

func (s *UpdateEventStreamingShrinkRequest) GetSinkShrink() *string {
	return s.SinkShrink
}

func (s *UpdateEventStreamingShrinkRequest) GetSourceShrink() *string {
	return s.SourceShrink
}

func (s *UpdateEventStreamingShrinkRequest) GetTransformsShrink() *string {
	return s.TransformsShrink
}

func (s *UpdateEventStreamingShrinkRequest) SetDescription(v string) *UpdateEventStreamingShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetEventStreamingName(v string) *UpdateEventStreamingShrinkRequest {
	s.EventStreamingName = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetFilterPattern(v string) *UpdateEventStreamingShrinkRequest {
	s.FilterPattern = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetRunOptionsShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.RunOptionsShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetSinkShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.SinkShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetSourceShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetTransformsShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.TransformsShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
