// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopStreamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStopStreamsResponseBody
	GetRequestId() *string
	SetResults(v []*BatchStopStreamsResponseBodyResults) *BatchStopStreamsResponseBody
	GetResults() []*BatchStopStreamsResponseBodyResults
}

type BatchStopStreamsResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchStopStreamsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchStopStreamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopStreamsResponseBody) GetResults() []*BatchStopStreamsResponseBodyResults {
	return s.Results
}

func (s *BatchStopStreamsResponseBody) SetRequestId(v string) *BatchStopStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopStreamsResponseBody) SetResults(v []*BatchStopStreamsResponseBodyResults) *BatchStopStreamsResponseBody {
	s.Results = v
	return s
}

func (s *BatchStopStreamsResponseBody) Validate() error {
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

type BatchStopStreamsResponseBodyResults struct {
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
	// 3100000*****00000002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchStopStreamsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchStopStreamsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchStopStreamsResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *BatchStopStreamsResponseBodyResults) GetName() *string {
	return s.Name
}

func (s *BatchStopStreamsResponseBodyResults) SetError(v string) *BatchStopStreamsResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchStopStreamsResponseBodyResults) SetId(v string) *BatchStopStreamsResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchStopStreamsResponseBodyResults) SetName(v string) *BatchStopStreamsResponseBodyResults {
	s.Name = &v
	return s
}

func (s *BatchStopStreamsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
