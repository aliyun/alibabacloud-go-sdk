// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlPreviewResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *FetchSqlPreviewResultsRequest
	GetQueryId() *string
	SetSessionClusterName(v string) *FetchSqlPreviewResultsRequest
	GetSessionClusterName() *string
	SetSessionId(v string) *FetchSqlPreviewResultsRequest
	GetSessionId() *string
}

type FetchSqlPreviewResultsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 00000000-0000-0000-0000-00000022****
	QueryId *string `json:"queryId,omitempty" xml:"queryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// session-test
	SessionClusterName *string `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 00000000-0000-0000-0000-00000033****
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s FetchSqlPreviewResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlPreviewResultsRequest) GoString() string {
	return s.String()
}

func (s *FetchSqlPreviewResultsRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *FetchSqlPreviewResultsRequest) GetSessionClusterName() *string {
	return s.SessionClusterName
}

func (s *FetchSqlPreviewResultsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FetchSqlPreviewResultsRequest) SetQueryId(v string) *FetchSqlPreviewResultsRequest {
	s.QueryId = &v
	return s
}

func (s *FetchSqlPreviewResultsRequest) SetSessionClusterName(v string) *FetchSqlPreviewResultsRequest {
	s.SessionClusterName = &v
	return s
}

func (s *FetchSqlPreviewResultsRequest) SetSessionId(v string) *FetchSqlPreviewResultsRequest {
	s.SessionId = &v
	return s
}

func (s *FetchSqlPreviewResultsRequest) Validate() error {
	return dara.Validate(s)
}
