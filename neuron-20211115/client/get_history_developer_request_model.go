// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *GetHistoryDeveloperRequest
	GetEnterpriseId() *int64
	SetMaxResults(v int32) *GetHistoryDeveloperRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetHistoryDeveloperRequest
	GetNextToken() *string
}

type GetHistoryDeveloperRequest struct {
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetHistoryDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryDeveloperRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryDeveloperRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *GetHistoryDeveloperRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetHistoryDeveloperRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetHistoryDeveloperRequest) SetEnterpriseId(v int64) *GetHistoryDeveloperRequest {
	s.EnterpriseId = &v
	return s
}

func (s *GetHistoryDeveloperRequest) SetMaxResults(v int32) *GetHistoryDeveloperRequest {
	s.MaxResults = &v
	return s
}

func (s *GetHistoryDeveloperRequest) SetNextToken(v string) *GetHistoryDeveloperRequest {
	s.NextToken = &v
	return s
}

func (s *GetHistoryDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
