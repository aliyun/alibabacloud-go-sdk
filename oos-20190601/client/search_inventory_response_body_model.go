// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []map[string]interface{}) *SearchInventoryResponseBody
	GetEntities() []map[string]interface{}
	SetMaxResults(v int32) *SearchInventoryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchInventoryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchInventoryResponseBody
	GetRequestId() *string
}

type SearchInventoryResponseBody struct {
	Entities []map[string]interface{} `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A81E4B2E-6B33-4BAE-9856-55DB7C893E01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInventoryResponseBody) GetEntities() []map[string]interface{} {
	return s.Entities
}

func (s *SearchInventoryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchInventoryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchInventoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchInventoryResponseBody) SetEntities(v []map[string]interface{}) *SearchInventoryResponseBody {
	s.Entities = v
	return s
}

func (s *SearchInventoryResponseBody) SetMaxResults(v int32) *SearchInventoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchInventoryResponseBody) SetNextToken(v string) *SearchInventoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchInventoryResponseBody) SetRequestId(v string) *SearchInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchInventoryResponseBody) Validate() error {
	return dara.Validate(s)
}
