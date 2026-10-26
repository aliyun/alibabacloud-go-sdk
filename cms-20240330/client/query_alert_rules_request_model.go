// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *QueryAlertRulesInput) *QueryAlertRulesRequest
	GetBody() *QueryAlertRulesInput
	SetClientToken(v string) *QueryAlertRulesRequest
	GetClientToken() *string
	SetMaxResults(v int32) *QueryAlertRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryAlertRulesRequest
	GetNextToken() *string
	SetQueryJson(v string) *QueryAlertRulesRequest
	GetQueryJson() *string
}

type QueryAlertRulesRequest struct {
	// The request body for querying alert rules.
	Body *QueryAlertRulesInput `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// xxxxx-xxxx-xxxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The maximum number of data records to return in the current read operation.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token that specifies the starting position for the current read operation. If this parameter is left empty, the read operation starts from the beginning.
	//
	// example:
	//
	// 123456
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The query filter conditions in JSON string format. The field semantics are the same as those in the request body. Supported fields include filter (such as uuid, displayName, labels, status, and enabled, where sibling fields use the AND logical operator), pagination (pageNumber and pageSize), and workspace. You must specify either this parameter or the body parameter. If both are specified, the body parameter takes precedence.
	//
	// example:
	//
	// {"filter":{"displayName":{"contains":"my-rule"}},"pagination":{"pageNumber":1,"pageSize":20}}
	QueryJson *string `json:"queryJson,omitempty" xml:"queryJson,omitempty"`
}

func (s QueryAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesRequest) GetBody() *QueryAlertRulesInput {
	return s.Body
}

func (s *QueryAlertRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *QueryAlertRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryAlertRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryAlertRulesRequest) GetQueryJson() *string {
	return s.QueryJson
}

func (s *QueryAlertRulesRequest) SetBody(v *QueryAlertRulesInput) *QueryAlertRulesRequest {
	s.Body = v
	return s
}

func (s *QueryAlertRulesRequest) SetClientToken(v string) *QueryAlertRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryAlertRulesRequest) SetMaxResults(v int32) *QueryAlertRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAlertRulesRequest) SetNextToken(v string) *QueryAlertRulesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAlertRulesRequest) SetQueryJson(v string) *QueryAlertRulesRequest {
	s.QueryJson = &v
	return s
}

func (s *QueryAlertRulesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
