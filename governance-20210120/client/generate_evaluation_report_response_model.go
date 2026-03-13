// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateEvaluationReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateEvaluationReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateEvaluationReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateEvaluationReportResponseBody) *GenerateEvaluationReportResponse
	GetBody() *GenerateEvaluationReportResponseBody
}

type GenerateEvaluationReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateEvaluationReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateEvaluationReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateEvaluationReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateEvaluationReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateEvaluationReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateEvaluationReportResponse) GetBody() *GenerateEvaluationReportResponseBody {
	return s.Body
}

func (s *GenerateEvaluationReportResponse) SetHeaders(v map[string]*string) *GenerateEvaluationReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateEvaluationReportResponse) SetStatusCode(v int32) *GenerateEvaluationReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateEvaluationReportResponse) SetBody(v *GenerateEvaluationReportResponseBody) *GenerateEvaluationReportResponse {
	s.Body = v
	return s
}

func (s *GenerateEvaluationReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
