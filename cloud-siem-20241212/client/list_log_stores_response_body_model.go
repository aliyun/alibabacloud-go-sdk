// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStores(v []*string) *ListLogStoresResponseBody
	GetLogStores() []*string
	SetMaxResults(v int32) *ListLogStoresResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListLogStoresResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListLogStoresResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLogStoresResponseBody
	GetTotalCount() *int32
}

type ListLogStoresResponseBody struct {
	LogStores []*string `json:"LogStores,omitempty" xml:"LogStores,omitempty" type:"Repeated"`
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLogStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponseBody) GetLogStores() []*string {
	return s.LogStores
}

func (s *ListLogStoresResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLogStoresResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLogStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogStoresResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLogStoresResponseBody) SetLogStores(v []*string) *ListLogStoresResponseBody {
	s.LogStores = v
	return s
}

func (s *ListLogStoresResponseBody) SetMaxResults(v int32) *ListLogStoresResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLogStoresResponseBody) SetNextToken(v string) *ListLogStoresResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLogStoresResponseBody) SetRequestId(v string) *ListLogStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogStoresResponseBody) SetTotalCount(v int32) *ListLogStoresResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLogStoresResponseBody) Validate() error {
	return dara.Validate(s)
}
