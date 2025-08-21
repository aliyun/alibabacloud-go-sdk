// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStopDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchStopDevicesResponseBodyResults) *BatchStopDevicesResponseBody
	GetResults() []*BatchStopDevicesResponseBodyResults
}

type BatchStopDevicesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStopDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStopDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopDevicesResponseBody) GetResults() []*BatchStopDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchStopDevicesResponseBody) SetRequestId(v string) *BatchStopDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopDevicesResponseBody) SetResults(v []*BatchStopDevicesResponseBodyResults) *BatchStopDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchStopDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchStopDevicesResponseBodyResults struct {
	// example:
	//
	// 32388487****92996
	Id      *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Streams []*BatchStopDevicesResponseBodyResultsStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Repeated"`
}

func (s BatchStopDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *BatchStopDevicesResponseBodyResults) GetStreams() []*BatchStopDevicesResponseBodyResultsStreams {
	return s.Streams
}

func (s *BatchStopDevicesResponseBodyResults) SetId(v string) *BatchStopDevicesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResults) SetStreams(v []*BatchStopDevicesResponseBodyResultsStreams) *BatchStopDevicesResponseBodyResults {
	s.Streams = v
	return s
}

func (s *BatchStopDevicesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type BatchStopDevicesResponseBodyResultsStreams struct {
	// example:
	//
	// Stream not found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 323884****9092997
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 32388*****39092997
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStopDevicesResponseBodyResultsStreams) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDevicesResponseBodyResultsStreams) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponseBodyResultsStreams) GetError() *string {
	return s.Error
}

func (s *BatchStopDevicesResponseBodyResultsStreams) GetId() *string {
	return s.Id
}

func (s *BatchStopDevicesResponseBodyResultsStreams) GetName() *string {
	return s.Name
}

func (s *BatchStopDevicesResponseBodyResultsStreams) SetError(v string) *BatchStopDevicesResponseBodyResultsStreams {
	s.Error = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResultsStreams) SetId(v string) *BatchStopDevicesResponseBodyResultsStreams {
	s.Id = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResultsStreams) SetName(v string) *BatchStopDevicesResponseBodyResultsStreams {
	s.Name = &v
	return s
}

func (s *BatchStopDevicesResponseBodyResultsStreams) Validate() error {
	return dara.Validate(s)
}
