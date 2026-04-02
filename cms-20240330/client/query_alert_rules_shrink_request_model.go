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
}

type QueryAlertRulesShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// xxxxx-xxxx-xxxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123456
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s *QueryAlertRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
