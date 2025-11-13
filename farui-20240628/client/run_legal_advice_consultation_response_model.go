// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLegalAdviceConsultationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunLegalAdviceConsultationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunLegalAdviceConsultationResponse
	GetStatusCode() *int32
	SetBody(v *RunLegalAdviceConsultationResponseBody) *RunLegalAdviceConsultationResponse
	GetBody() *RunLegalAdviceConsultationResponseBody
}

type RunLegalAdviceConsultationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLegalAdviceConsultationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLegalAdviceConsultationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationResponse) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunLegalAdviceConsultationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunLegalAdviceConsultationResponse) GetBody() *RunLegalAdviceConsultationResponseBody {
	return s.Body
}

func (s *RunLegalAdviceConsultationResponse) SetHeaders(v map[string]*string) *RunLegalAdviceConsultationResponse {
	s.Headers = v
	return s
}

func (s *RunLegalAdviceConsultationResponse) SetStatusCode(v int32) *RunLegalAdviceConsultationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLegalAdviceConsultationResponse) SetBody(v *RunLegalAdviceConsultationResponseBody) *RunLegalAdviceConsultationResponse {
	s.Body = v
	return s
}

func (s *RunLegalAdviceConsultationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
