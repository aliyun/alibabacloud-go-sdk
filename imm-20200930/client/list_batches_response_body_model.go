// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatches(v []*DataIngestion) *ListBatchesResponseBody
	GetBatches() []*DataIngestion
	SetNextToken(v string) *ListBatchesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBatchesResponseBody
	GetRequestId() *string
}

type ListBatchesResponseBody struct {
	// The batch processing tasks.
	Batches []*DataIngestion `json:"Batches,omitempty" xml:"Batches,omitempty" type:"Repeated"`
	// The pagination token.
	//
	// The pagination token is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter. The next call to the operation returns results lexicographically after the NextToken parameter value.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEDC9B1F-30F2-4C1F-8ED2-B7860187****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBatchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchesResponseBody) GetBatches() []*DataIngestion {
	return s.Batches
}

func (s *ListBatchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchesResponseBody) SetBatches(v []*DataIngestion) *ListBatchesResponseBody {
	s.Batches = v
	return s
}

func (s *ListBatchesResponseBody) SetNextToken(v string) *ListBatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBatchesResponseBody) SetRequestId(v string) *ListBatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchesResponseBody) Validate() error {
	return dara.Validate(s)
}
