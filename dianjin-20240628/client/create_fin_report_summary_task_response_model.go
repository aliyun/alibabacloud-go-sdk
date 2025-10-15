// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFinReportSummaryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFinReportSummaryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFinReportSummaryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFinReportSummaryTaskResponseBody) *CreateFinReportSummaryTaskResponse
	GetBody() *CreateFinReportSummaryTaskResponseBody
}

type CreateFinReportSummaryTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFinReportSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFinReportSummaryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFinReportSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFinReportSummaryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFinReportSummaryTaskResponse) GetBody() *CreateFinReportSummaryTaskResponseBody {
	return s.Body
}

func (s *CreateFinReportSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateFinReportSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) SetStatusCode(v int32) *CreateFinReportSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) SetBody(v *CreateFinReportSummaryTaskResponseBody) *CreateFinReportSummaryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
