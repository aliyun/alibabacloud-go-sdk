// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*File) *FuzzyQueryResponseBody
	GetFiles() []*File
	SetNextToken(v string) *FuzzyQueryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *FuzzyQueryResponseBody
	GetRequestId() *string
	SetTotalHits(v int64) *FuzzyQueryResponseBody
	GetTotalHits() *int64
}

type FuzzyQueryResponseBody struct {
	// The files.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// A pagination token.
	//
	// It can be used in the next request to retrieve a new page of results.
	//
	// If NextToken is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B3D5E0A-D8B8-4DA0-8127-ED32C851****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of hits.
	TotalHits *int64 `json:"TotalHits,omitempty" xml:"TotalHits,omitempty"`
}

func (s FuzzyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FuzzyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FuzzyQueryResponseBody) GetFiles() []*File {
	return s.Files
}

func (s *FuzzyQueryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *FuzzyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FuzzyQueryResponseBody) GetTotalHits() *int64 {
	return s.TotalHits
}

func (s *FuzzyQueryResponseBody) SetFiles(v []*File) *FuzzyQueryResponseBody {
	s.Files = v
	return s
}

func (s *FuzzyQueryResponseBody) SetNextToken(v string) *FuzzyQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryResponseBody) SetRequestId(v string) *FuzzyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FuzzyQueryResponseBody) SetTotalHits(v int64) *FuzzyQueryResponseBody {
	s.TotalHits = &v
	return s
}

func (s *FuzzyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
