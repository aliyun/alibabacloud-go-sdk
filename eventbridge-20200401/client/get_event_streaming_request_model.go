// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventStreamingName(v string) *GetEventStreamingRequest
	GetEventStreamingName() *string
}

type GetEventStreamingRequest struct {
	// The name of the event stream whose details you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// myeventstreaming
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s GetEventStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *GetEventStreamingRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *GetEventStreamingRequest) SetEventStreamingName(v string) *GetEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *GetEventStreamingRequest) Validate() error {
	return dara.Validate(s)
}
