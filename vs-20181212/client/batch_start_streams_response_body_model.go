// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartStreamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStartStreamsResponseBody
	GetRequestId() *string
	SetResults(v []*BatchStartStreamsResponseBodyResults) *BatchStartStreamsResponseBody
	GetResults() []*BatchStartStreamsResponseBodyResults
}

type BatchStartStreamsResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStartStreamsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStartStreamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStartStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStartStreamsResponseBody) GetResults() []*BatchStartStreamsResponseBodyResults {
	return s.Results
}

func (s *BatchStartStreamsResponseBody) SetRequestId(v string) *BatchStartStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartStreamsResponseBody) SetResults(v []*BatchStartStreamsResponseBodyResults) *BatchStartStreamsResponseBody {
	s.Results = v
	return s
}

func (s *BatchStartStreamsResponseBody) Validate() error {
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

type BatchStartStreamsResponseBodyResults struct {
	// example:
	//
	// stream not found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 31000000*****0000002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStartStreamsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchStartStreamsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchStartStreamsResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *BatchStartStreamsResponseBodyResults) GetName() *string {
	return s.Name
}

func (s *BatchStartStreamsResponseBodyResults) SetError(v string) *BatchStartStreamsResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchStartStreamsResponseBodyResults) SetId(v string) *BatchStartStreamsResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStartStreamsResponseBodyResults) SetName(v string) *BatchStartStreamsResponseBodyResults {
	s.Name = &v
	return s
}

func (s *BatchStartStreamsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
