// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseEventStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventStreamingName(v string) *PauseEventStreamingRequest
	GetEventStreamingName() *string
}

type PauseEventStreamingRequest struct {
	// The name of the event stream that you want to stop.
	//
	// This parameter is required.
	//
	// example:
	//
	// rocketmq-sync
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s PauseEventStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *PauseEventStreamingRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *PauseEventStreamingRequest) SetEventStreamingName(v string) *PauseEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *PauseEventStreamingRequest) Validate() error {
	return dara.Validate(s)
}
