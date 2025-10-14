// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLEvaluateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSQLEvaluateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSQLEvaluateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSQLEvaluateTaskResponseBody) *CreateSQLEvaluateTaskResponse
	GetBody() *CreateSQLEvaluateTaskResponseBody
}

type CreateSQLEvaluateTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSQLEvaluateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSQLEvaluateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLEvaluateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSQLEvaluateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSQLEvaluateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSQLEvaluateTaskResponse) GetBody() *CreateSQLEvaluateTaskResponseBody {
	return s.Body
}

func (s *CreateSQLEvaluateTaskResponse) SetHeaders(v map[string]*string) *CreateSQLEvaluateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSQLEvaluateTaskResponse) SetStatusCode(v int32) *CreateSQLEvaluateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSQLEvaluateTaskResponse) SetBody(v *CreateSQLEvaluateTaskResponseBody) *CreateSQLEvaluateTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSQLEvaluateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
