// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoSyncRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoSyncRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoSyncRuleResponseBody) *CreateRepoSyncRuleResponse
	GetBody() *CreateRepoSyncRuleResponseBody
}

type CreateRepoSyncRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSyncRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSyncRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoSyncRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoSyncRuleResponse) GetBody() *CreateRepoSyncRuleResponseBody {
	return s.Body
}

func (s *CreateRepoSyncRuleResponse) SetHeaders(v map[string]*string) *CreateRepoSyncRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSyncRuleResponse) SetStatusCode(v int32) *CreateRepoSyncRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSyncRuleResponse) SetBody(v *CreateRepoSyncRuleResponseBody) *CreateRepoSyncRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRepoSyncRuleResponse) Validate() error {
	return dara.Validate(s)
}
