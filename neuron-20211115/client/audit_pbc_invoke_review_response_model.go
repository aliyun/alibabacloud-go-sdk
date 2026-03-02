// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditPbcInvokeReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuditPbcInvokeReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuditPbcInvokeReviewResponse
	GetStatusCode() *int32
	SetBody(v *AuditPbcInvokeReviewResponseBody) *AuditPbcInvokeReviewResponse
	GetBody() *AuditPbcInvokeReviewResponseBody
}

type AuditPbcInvokeReviewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditPbcInvokeReviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditPbcInvokeReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s AuditPbcInvokeReviewResponse) GoString() string {
	return s.String()
}

func (s *AuditPbcInvokeReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuditPbcInvokeReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuditPbcInvokeReviewResponse) GetBody() *AuditPbcInvokeReviewResponseBody {
	return s.Body
}

func (s *AuditPbcInvokeReviewResponse) SetHeaders(v map[string]*string) *AuditPbcInvokeReviewResponse {
	s.Headers = v
	return s
}

func (s *AuditPbcInvokeReviewResponse) SetStatusCode(v int32) *AuditPbcInvokeReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditPbcInvokeReviewResponse) SetBody(v *AuditPbcInvokeReviewResponseBody) *AuditPbcInvokeReviewResponse {
	s.Body = v
	return s
}

func (s *AuditPbcInvokeReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
