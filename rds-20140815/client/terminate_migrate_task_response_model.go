// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateMigrateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateMigrateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateMigrateTaskResponse
	GetStatusCode() *int32
	SetBody(v *TerminateMigrateTaskResponseBody) *TerminateMigrateTaskResponse
	GetBody() *TerminateMigrateTaskResponseBody
}

type TerminateMigrateTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateMigrateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateMigrateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateMigrateTaskResponse) GoString() string {
	return s.String()
}

func (s *TerminateMigrateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateMigrateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateMigrateTaskResponse) GetBody() *TerminateMigrateTaskResponseBody {
	return s.Body
}

func (s *TerminateMigrateTaskResponse) SetHeaders(v map[string]*string) *TerminateMigrateTaskResponse {
	s.Headers = v
	return s
}

func (s *TerminateMigrateTaskResponse) SetStatusCode(v int32) *TerminateMigrateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateMigrateTaskResponse) SetBody(v *TerminateMigrateTaskResponseBody) *TerminateMigrateTaskResponse {
	s.Body = v
	return s
}

func (s *TerminateMigrateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
