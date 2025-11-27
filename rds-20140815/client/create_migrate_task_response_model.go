// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMigrateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMigrateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMigrateTaskResponseBody) *CreateMigrateTaskResponse
	GetBody() *CreateMigrateTaskResponseBody
}

type CreateMigrateTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMigrateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMigrateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMigrateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMigrateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMigrateTaskResponse) GetBody() *CreateMigrateTaskResponseBody {
	return s.Body
}

func (s *CreateMigrateTaskResponse) SetHeaders(v map[string]*string) *CreateMigrateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMigrateTaskResponse) SetStatusCode(v int32) *CreateMigrateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMigrateTaskResponse) SetBody(v *CreateMigrateTaskResponseBody) *CreateMigrateTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMigrateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
