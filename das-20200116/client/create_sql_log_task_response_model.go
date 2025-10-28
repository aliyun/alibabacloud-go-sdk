// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlLogTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSqlLogTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSqlLogTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSqlLogTaskResponseBody) *CreateSqlLogTaskResponse
	GetBody() *CreateSqlLogTaskResponseBody
}

type CreateSqlLogTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlLogTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlLogTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlLogTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlLogTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSqlLogTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSqlLogTaskResponse) GetBody() *CreateSqlLogTaskResponseBody {
	return s.Body
}

func (s *CreateSqlLogTaskResponse) SetHeaders(v map[string]*string) *CreateSqlLogTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlLogTaskResponse) SetStatusCode(v int32) *CreateSqlLogTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlLogTaskResponse) SetBody(v *CreateSqlLogTaskResponseBody) *CreateSqlLogTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSqlLogTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
