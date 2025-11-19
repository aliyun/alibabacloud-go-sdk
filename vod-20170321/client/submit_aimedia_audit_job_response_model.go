// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIMediaAuditJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIMediaAuditJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIMediaAuditJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIMediaAuditJobResponseBody) *SubmitAIMediaAuditJobResponse
	GetBody() *SubmitAIMediaAuditJobResponseBody
}

type SubmitAIMediaAuditJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIMediaAuditJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIMediaAuditJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIMediaAuditJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIMediaAuditJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIMediaAuditJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIMediaAuditJobResponse) GetBody() *SubmitAIMediaAuditJobResponseBody {
	return s.Body
}

func (s *SubmitAIMediaAuditJobResponse) SetHeaders(v map[string]*string) *SubmitAIMediaAuditJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIMediaAuditJobResponse) SetStatusCode(v int32) *SubmitAIMediaAuditJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIMediaAuditJobResponse) SetBody(v *SubmitAIMediaAuditJobResponseBody) *SubmitAIMediaAuditJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIMediaAuditJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
