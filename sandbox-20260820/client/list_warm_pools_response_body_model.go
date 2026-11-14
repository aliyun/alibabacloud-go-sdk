// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarmPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWarmPoolsResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListWarmPoolsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWarmPoolsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWarmPoolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWarmPoolsResponseBody
	GetRequestId() *string
	SetWarmPools(v []*PublicWarmPoolConfig) *ListWarmPoolsResponseBody
	GetWarmPools() []*PublicWarmPoolConfig
}

type ListWarmPoolsResponseBody struct {
	// The business code in the response.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The actual maximum number of results returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token. An empty value indicates that no more results are available.
	//
	// example:
	//
	// example-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01234567-89AB-CDEF-0123-456789ABCDEF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of warm pools.
	WarmPools []*PublicWarmPoolConfig `json:"warmPools" xml:"warmPools" type:"Repeated"`
}

func (s ListWarmPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarmPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarmPoolsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWarmPoolsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWarmPoolsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWarmPoolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWarmPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarmPoolsResponseBody) GetWarmPools() []*PublicWarmPoolConfig {
	return s.WarmPools
}

func (s *ListWarmPoolsResponseBody) SetCode(v string) *ListWarmPoolsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarmPoolsResponseBody) SetMaxResults(v int32) *ListWarmPoolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWarmPoolsResponseBody) SetMessage(v string) *ListWarmPoolsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarmPoolsResponseBody) SetNextToken(v string) *ListWarmPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWarmPoolsResponseBody) SetRequestId(v string) *ListWarmPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarmPoolsResponseBody) SetWarmPools(v []*PublicWarmPoolConfig) *ListWarmPoolsResponseBody {
	s.WarmPools = v
	return s
}

func (s *ListWarmPoolsResponseBody) Validate() error {
	if s.WarmPools != nil {
		for _, item := range s.WarmPools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
