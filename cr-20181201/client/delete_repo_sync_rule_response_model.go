// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoSyncRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepoSyncRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepoSyncRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepoSyncRuleResponseBody) *DeleteRepoSyncRuleResponse
	GetBody() *DeleteRepoSyncRuleResponseBody
}

type DeleteRepoSyncRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoSyncRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoSyncRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoSyncRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoSyncRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepoSyncRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepoSyncRuleResponse) GetBody() *DeleteRepoSyncRuleResponseBody {
	return s.Body
}

func (s *DeleteRepoSyncRuleResponse) SetHeaders(v map[string]*string) *DeleteRepoSyncRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoSyncRuleResponse) SetStatusCode(v int32) *DeleteRepoSyncRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoSyncRuleResponse) SetBody(v *DeleteRepoSyncRuleResponseBody) *DeleteRepoSyncRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRepoSyncRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
