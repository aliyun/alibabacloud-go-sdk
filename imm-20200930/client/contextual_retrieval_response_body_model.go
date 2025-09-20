// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualRetrievalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ContextualRetrievalResponseBody
	GetRequestId() *string
	SetResults(v []*File) *ContextualRetrievalResponseBody
	GetResults() []*File
}

type ContextualRetrievalResponseBody struct {
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*File `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ContextualRetrievalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContextualRetrievalResponseBody) GoString() string {
	return s.String()
}

func (s *ContextualRetrievalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContextualRetrievalResponseBody) GetResults() []*File {
	return s.Results
}

func (s *ContextualRetrievalResponseBody) SetRequestId(v string) *ContextualRetrievalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContextualRetrievalResponseBody) SetResults(v []*File) *ContextualRetrievalResponseBody {
	s.Results = v
	return s
}

func (s *ContextualRetrievalResponseBody) Validate() error {
	return dara.Validate(s)
}
