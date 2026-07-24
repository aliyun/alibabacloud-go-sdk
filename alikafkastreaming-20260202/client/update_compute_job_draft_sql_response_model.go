// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobDraftSqlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeJobDraftSqlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeJobDraftSqlResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeJobDraftSqlResponseBody) *UpdateComputeJobDraftSqlResponse
	GetBody() *UpdateComputeJobDraftSqlResponseBody
}

type UpdateComputeJobDraftSqlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeJobDraftSqlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeJobDraftSqlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobDraftSqlResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobDraftSqlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeJobDraftSqlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeJobDraftSqlResponse) GetBody() *UpdateComputeJobDraftSqlResponseBody {
	return s.Body
}

func (s *UpdateComputeJobDraftSqlResponse) SetHeaders(v map[string]*string) *UpdateComputeJobDraftSqlResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeJobDraftSqlResponse) SetStatusCode(v int32) *UpdateComputeJobDraftSqlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeJobDraftSqlResponse) SetBody(v *UpdateComputeJobDraftSqlResponseBody) *UpdateComputeJobDraftSqlResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeJobDraftSqlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
