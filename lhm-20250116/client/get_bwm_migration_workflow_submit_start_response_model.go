// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationWorkflowSubmitStartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBwmMigrationWorkflowSubmitStartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBwmMigrationWorkflowSubmitStartResponse
	GetStatusCode() *int32
	SetBody(v *GetBwmMigrationWorkflowSubmitStartResponseBody) *GetBwmMigrationWorkflowSubmitStartResponse
	GetBody() *GetBwmMigrationWorkflowSubmitStartResponseBody
}

type GetBwmMigrationWorkflowSubmitStartResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBwmMigrationWorkflowSubmitStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBwmMigrationWorkflowSubmitStartResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationWorkflowSubmitStartResponse) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) GetBody() *GetBwmMigrationWorkflowSubmitStartResponseBody {
	return s.Body
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) SetHeaders(v map[string]*string) *GetBwmMigrationWorkflowSubmitStartResponse {
	s.Headers = v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) SetStatusCode(v int32) *GetBwmMigrationWorkflowSubmitStartResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) SetBody(v *GetBwmMigrationWorkflowSubmitStartResponseBody) *GetBwmMigrationWorkflowSubmitStartResponse {
	s.Body = v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
