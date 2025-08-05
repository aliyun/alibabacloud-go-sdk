// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEventStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventStreamingName(v string) *StartEventStreamingRequest
	GetEventStreamingName() *string
}

type StartEventStreamingRequest struct {
	// The name of the event stream that you want to enable.
	//
	// This parameter is required.
	//
	// example:
	//
	// rocketmq-sync
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s StartEventStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s StartEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *StartEventStreamingRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *StartEventStreamingRequest) SetEventStreamingName(v string) *StartEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *StartEventStreamingRequest) Validate() error {
	return dara.Validate(s)
}
