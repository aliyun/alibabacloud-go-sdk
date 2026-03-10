// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMemoriesResponseBody
	GetRequestId() *string
	SetResults(v []*AddMemoriesResponseBodyResults) *AddMemoriesResponseBody
	GetResults() []*AddMemoriesResponseBodyResults
}

type AddMemoriesResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*AddMemoriesResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s AddMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMemoriesResponseBody) GetResults() []*AddMemoriesResponseBodyResults {
	return s.Results
}

func (s *AddMemoriesResponseBody) SetRequestId(v string) *AddMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMemoriesResponseBody) SetResults(v []*AddMemoriesResponseBodyResults) *AddMemoriesResponseBody {
	s.Results = v
	return s
}

func (s *AddMemoriesResponseBody) Validate() error {
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

type AddMemoriesResponseBodyResults struct {
	// example:
	//
	// 897294a7-67a4-4f60-976c-e136edc5f97e
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Memory processing has been queued for background execution
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AddMemoriesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddMemoriesResponseBodyResults) GetEventId() *string {
	return s.EventId
}

func (s *AddMemoriesResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AddMemoriesResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *AddMemoriesResponseBodyResults) SetEventId(v string) *AddMemoriesResponseBodyResults {
	s.EventId = &v
	return s
}

func (s *AddMemoriesResponseBodyResults) SetMessage(v string) *AddMemoriesResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddMemoriesResponseBodyResults) SetStatus(v string) *AddMemoriesResponseBodyResults {
	s.Status = &v
	return s
}

func (s *AddMemoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
