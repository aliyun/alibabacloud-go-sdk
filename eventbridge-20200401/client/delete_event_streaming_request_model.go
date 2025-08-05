// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventStreamingName(v string) *DeleteEventStreamingRequest
	GetEventStreamingName() *string
}

type DeleteEventStreamingRequest struct {
	// The name of the event stream that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// rocketmq-sync
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s DeleteEventStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventStreamingRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *DeleteEventStreamingRequest) SetEventStreamingName(v string) *DeleteEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *DeleteEventStreamingRequest) Validate() error {
	return dara.Validate(s)
}
