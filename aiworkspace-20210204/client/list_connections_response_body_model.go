// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnections(v []*Connection) *ListConnectionsResponseBody
	GetConnections() []*Connection
	SetMaxResults(v int32) *ListConnectionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListConnectionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListConnectionsResponseBody
	GetTotalCount() *int32
}

type ListConnectionsResponseBody struct {
	// The connection list.
	Connections []*Connection `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that indicates the start position from which to retrieve data on the next page.
	//
	// example:
	//
	// 15
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B2C51F93-1C07-5477-9705-5FDB****F19F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of connections that meet the filter conditions.
	//
	// example:
	//
	// 27
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBody) GetConnections() []*Connection {
	return s.Connections
}

func (s *ListConnectionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListConnectionsResponseBody) SetConnections(v []*Connection) *ListConnectionsResponseBody {
	s.Connections = v
	return s
}

func (s *ListConnectionsResponseBody) SetMaxResults(v int32) *ListConnectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionsResponseBody) SetNextToken(v string) *ListConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectionsResponseBody) SetRequestId(v string) *ListConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectionsResponseBody) SetTotalCount(v int32) *ListConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}
