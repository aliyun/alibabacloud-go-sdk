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
}

type QueryAlertRulesRequest struct {
	Body *QueryAlertRulesInput `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryAlertRulesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
