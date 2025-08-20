// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncTaskByRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoSyncTaskByRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoSyncTaskByRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoSyncTaskByRuleResponseBody) *CreateRepoSyncTaskByRuleResponse
	GetBody() *CreateRepoSyncTaskByRuleResponseBody
}

type CreateRepoSyncTaskByRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSyncTaskByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSyncTaskByRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncTaskByRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskByRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoSyncTaskByRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoSyncTaskByRuleResponse) GetBody() *CreateRepoSyncTaskByRuleResponseBody {
	return s.Body
}

func (s *CreateRepoSyncTaskByRuleResponse) SetHeaders(v map[string]*string) *CreateRepoSyncTaskByRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponse) SetStatusCode(v int32) *CreateRepoSyncTaskByRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponse) SetBody(v *CreateRepoSyncTaskByRuleResponseBody) *CreateRepoSyncTaskByRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponse) Validate() error {
	return dara.Validate(s)
}
