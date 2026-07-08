// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStartDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchStartDevicesResponseBodyResults) *BatchStartDevicesResponseBody
	GetResults() []*BatchStartDevicesResponseBodyResults
}

type BatchStartDevicesResponseBody struct {
	// The request ID of this task.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of results.
	Results []*BatchStartDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStartDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStartDevicesResponseBody) GetResults() []*BatchStartDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchStartDevicesResponseBody) SetRequestId(v string) *BatchStartDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartDevicesResponseBody) SetResults(v []*BatchStartDevicesResponseBodyResults) *BatchStartDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchStartDevicesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchStartDevicesResponseBodyResults struct {
	// Device ID.
	//
	// example:
	//
	// 32388487****92996-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// List of device streams.
	Streams []*BatchStartDevicesResponseBodyResultsStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
}

func (s BatchStartDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *BatchStartDevicesResponseBodyResults) GetStreams() []*BatchStartDevicesResponseBodyResultsStreams {
	return s.Streams
}

func (s *BatchStartDevicesResponseBodyResults) SetId(v string) *BatchStartDevicesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResults) SetStreams(v []*BatchStartDevicesResponseBodyResultsStreams) *BatchStartDevicesResponseBodyResults {
	s.Streams = v
	return s
}

func (s *BatchStartDevicesResponseBodyResults) Validate() error {
	if s.Streams != nil {
		for _, item := range s.Streams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchStartDevicesResponseBodyResultsStreams struct {
	// Error message for the stream. This field appears only when an error occurs.
	//
	// example:
	//
	// Stream not found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// Stream ID.
	//
	// example:
	//
	// 3238848****092997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Stream name.
	//
	// example:
	//
	// 310101*****187542126
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStartDevicesResponseBodyResultsStreams) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDevicesResponseBodyResultsStreams) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponseBodyResultsStreams) GetError() *string {
	return s.Error
}

func (s *BatchStartDevicesResponseBodyResultsStreams) GetId() *string {
	return s.Id
}

func (s *BatchStartDevicesResponseBodyResultsStreams) GetName() *string {
	return s.Name
}

func (s *BatchStartDevicesResponseBodyResultsStreams) SetError(v string) *BatchStartDevicesResponseBodyResultsStreams {
	s.Error = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResultsStreams) SetId(v string) *BatchStartDevicesResponseBodyResultsStreams {
	s.Id = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResultsStreams) SetName(v string) *BatchStartDevicesResponseBodyResultsStreams {
	s.Name = &v
	return s
}

func (s *BatchStartDevicesResponseBodyResultsStreams) Validate() error {
	return dara.Validate(s)
}
