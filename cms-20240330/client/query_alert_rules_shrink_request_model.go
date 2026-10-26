// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *QueryAlertRulesShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *QueryAlertRulesShrinkRequest
	GetClientToken() *string
	SetMaxResults(v int32) *QueryAlertRulesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryAlertRulesShrinkRequest
	GetNextToken() *string
	SetQueryJson(v string) *QueryAlertRulesShrinkRequest
	GetQueryJson() *string
}

type QueryAlertRulesShrinkRequest struct {
	// The request body for querying alert rules.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
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

func (s QueryAlertRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *QueryAlertRulesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *QueryAlertRulesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryAlertRulesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryAlertRulesShrinkRequest) GetQueryJson() *string {
	return s.QueryJson
}

func (s *QueryAlertRulesShrinkRequest) SetBodyShrink(v string) *QueryAlertRulesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *QueryAlertRulesShrinkRequest) SetClientToken(v string) *QueryAlertRulesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryAlertRulesShrinkRequest) SetMaxResults(v int32) *QueryAlertRulesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAlertRulesShrinkRequest) SetNextToken(v string) *QueryAlertRulesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAlertRulesShrinkRequest) SetQueryJson(v string) *QueryAlertRulesShrinkRequest {
	s.QueryJson = &v
	return s
}

func (s *QueryAlertRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
