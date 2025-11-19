// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageAuditJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIImageAuditJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIImageAuditJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIImageAuditJobResponseBody) *SubmitAIImageAuditJobResponse
	GetBody() *SubmitAIImageAuditJobResponseBody
}

type SubmitAIImageAuditJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIImageAuditJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIImageAuditJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageAuditJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIImageAuditJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIImageAuditJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIImageAuditJobResponse) GetBody() *SubmitAIImageAuditJobResponseBody {
	return s.Body
}

func (s *SubmitAIImageAuditJobResponse) SetHeaders(v map[string]*string) *SubmitAIImageAuditJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIImageAuditJobResponse) SetStatusCode(v int32) *SubmitAIImageAuditJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIImageAuditJobResponse) SetBody(v *SubmitAIImageAuditJobResponseBody) *SubmitAIImageAuditJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIImageAuditJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
