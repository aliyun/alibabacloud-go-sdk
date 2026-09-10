// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationTaskWriterWorkflowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBwmMigrationTaskWriterWorkflowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBwmMigrationTaskWriterWorkflowListResponse
	GetStatusCode() *int32
	SetBody(v *GetBwmMigrationTaskWriterWorkflowListResponseBody) *GetBwmMigrationTaskWriterWorkflowListResponse
	GetBody() *GetBwmMigrationTaskWriterWorkflowListResponseBody
}

type GetBwmMigrationTaskWriterWorkflowListResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBwmMigrationTaskWriterWorkflowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBwmMigrationTaskWriterWorkflowListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationTaskWriterWorkflowListResponse) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) GetBody() *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	return s.Body
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) SetHeaders(v map[string]*string) *GetBwmMigrationTaskWriterWorkflowListResponse {
	s.Headers = v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) SetStatusCode(v int32) *GetBwmMigrationTaskWriterWorkflowListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) SetBody(v *GetBwmMigrationTaskWriterWorkflowListResponseBody) *GetBwmMigrationTaskWriterWorkflowListResponse {
	s.Body = v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
