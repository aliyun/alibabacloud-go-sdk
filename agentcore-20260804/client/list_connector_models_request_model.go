// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorKeyId(v string) *ListConnectorModelsRequest
	GetConnectorKeyId() *string
	SetMaxResults(v int32) *ListConnectorModelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListConnectorModelsRequest
	GetNextToken() *string
}

type ListConnectorModelsRequest struct {
	// The ID of a service account key. This parameter is required when multiple keys exist. You can leave this parameter empty if only one key exists.
	//
	// example:
	//
	// ckey-xxxx
	ConnectorKeyId *string `json:"connectorKeyId,omitempty" xml:"connectorKeyId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// dGVzdA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListConnectorModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorModelsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorModelsRequest) GetConnectorKeyId() *string {
	return s.ConnectorKeyId
}

func (s *ListConnectorModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectorModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectorModelsRequest) SetConnectorKeyId(v string) *ListConnectorModelsRequest {
	s.ConnectorKeyId = &v
	return s
}

func (s *ListConnectorModelsRequest) SetMaxResults(v int32) *ListConnectorModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectorModelsRequest) SetNextToken(v string) *ListConnectorModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectorModelsRequest) Validate() error {
	return dara.Validate(s)
}
