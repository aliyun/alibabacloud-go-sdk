// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListMultiAccountTagValuesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMultiAccountTagValuesResponseBody
	GetRequestId() *string
	SetTagValues(v []*string) *ListMultiAccountTagValuesResponseBody
	GetTagValues() []*string
}

type ListMultiAccountTagValuesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65DC00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountTagValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiAccountTagValuesResponseBody) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListMultiAccountTagValuesResponseBody) SetNextToken(v string) *ListMultiAccountTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetRequestId(v string) *ListMultiAccountTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetTagValues(v []*string) *ListMultiAccountTagValuesResponseBody {
	s.TagValues = v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
