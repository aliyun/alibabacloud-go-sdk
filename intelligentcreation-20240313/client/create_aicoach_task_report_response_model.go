// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAICoachTaskReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAICoachTaskReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateAICoachTaskReportResponseBody) *CreateAICoachTaskReportResponse
	GetBody() *CreateAICoachTaskReportResponseBody
}

type CreateAICoachTaskReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAICoachTaskReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAICoachTaskReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskReportResponse) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAICoachTaskReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAICoachTaskReportResponse) GetBody() *CreateAICoachTaskReportResponseBody {
	return s.Body
}

func (s *CreateAICoachTaskReportResponse) SetHeaders(v map[string]*string) *CreateAICoachTaskReportResponse {
	s.Headers = v
	return s
}

func (s *CreateAICoachTaskReportResponse) SetStatusCode(v int32) *CreateAICoachTaskReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAICoachTaskReportResponse) SetBody(v *CreateAICoachTaskReportResponseBody) *CreateAICoachTaskReportResponse {
	s.Body = v
	return s
}

func (s *CreateAICoachTaskReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
