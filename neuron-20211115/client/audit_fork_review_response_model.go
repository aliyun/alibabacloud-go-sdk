// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditForkReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuditForkReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuditForkReviewResponse
	GetStatusCode() *int32
	SetBody(v *AuditForkReviewResponseBody) *AuditForkReviewResponse
	GetBody() *AuditForkReviewResponseBody
}

type AuditForkReviewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditForkReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditForkReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s AuditForkReviewResponse) GoString() string {
	return s.String()
}

func (s *AuditForkReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuditForkReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuditForkReviewResponse) GetBody() *AuditForkReviewResponseBody {
	return s.Body
}

func (s *AuditForkReviewResponse) SetHeaders(v map[string]*string) *AuditForkReviewResponse {
	s.Headers = v
	return s
}

func (s *AuditForkReviewResponse) SetStatusCode(v int32) *AuditForkReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditForkReviewResponse) SetBody(v *AuditForkReviewResponseBody) *AuditForkReviewResponse {
	s.Body = v
	return s
}

func (s *AuditForkReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
