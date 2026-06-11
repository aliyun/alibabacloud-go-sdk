// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContextsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddContextsResponseBody
	GetRequestId() *string
	SetResults(v []*AddContextsResponseBodyResults) *AddContextsResponseBody
	GetResults() []*AddContextsResponseBodyResults
}

type AddContextsResponseBody struct {
	// The unique ID for the request.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// An array of objects containing the results of the write operation.
	Results []*AddContextsResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s AddContextsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddContextsResponseBody) GoString() string {
	return s.String()
}

func (s *AddContextsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddContextsResponseBody) GetResults() []*AddContextsResponseBodyResults {
	return s.Results
}

func (s *AddContextsResponseBody) SetRequestId(v string) *AddContextsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContextsResponseBody) SetResults(v []*AddContextsResponseBodyResults) *AddContextsResponseBody {
	s.Results = v
	return s
}

func (s *AddContextsResponseBody) Validate() error {
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

type AddContextsResponseBodyResults struct {
	// The ID of the written record or event.
	//
	// example:
	//
	// 897294a7-67a4-4f60-976c-e136edc5f97e
	ContextId *string `json:"contextId,omitempty" xml:"contextId,omitempty"`
	// The write status. Can be "accepted", "queued", or "created".
	//
	// example:
	//
	// accepted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AddContextsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AddContextsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddContextsResponseBodyResults) GetContextId() *string {
	return s.ContextId
}

func (s *AddContextsResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *AddContextsResponseBodyResults) SetContextId(v string) *AddContextsResponseBodyResults {
	s.ContextId = &v
	return s
}

func (s *AddContextsResponseBodyResults) SetStatus(v string) *AddContextsResponseBodyResults {
	s.Status = &v
	return s
}

func (s *AddContextsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
