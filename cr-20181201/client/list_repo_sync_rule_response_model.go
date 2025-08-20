// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoSyncRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoSyncRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoSyncRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoSyncRuleResponseBody) *ListRepoSyncRuleResponse
	GetBody() *ListRepoSyncRuleResponseBody
}

type ListRepoSyncRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoSyncRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoSyncRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncRuleResponse) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoSyncRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoSyncRuleResponse) GetBody() *ListRepoSyncRuleResponseBody {
	return s.Body
}

func (s *ListRepoSyncRuleResponse) SetHeaders(v map[string]*string) *ListRepoSyncRuleResponse {
	s.Headers = v
	return s
}

func (s *ListRepoSyncRuleResponse) SetStatusCode(v int32) *ListRepoSyncRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoSyncRuleResponse) SetBody(v *ListRepoSyncRuleResponseBody) *ListRepoSyncRuleResponse {
	s.Body = v
	return s
}

func (s *ListRepoSyncRuleResponse) Validate() error {
	return dara.Validate(s)
}
