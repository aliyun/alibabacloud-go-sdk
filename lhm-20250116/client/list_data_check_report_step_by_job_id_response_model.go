// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportStepByJobIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckReportStepByJobIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckReportStepByJobIdResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckReportStepByJobIdResponseBody) *ListDataCheckReportStepByJobIdResponse
	GetBody() *ListDataCheckReportStepByJobIdResponseBody
}

type ListDataCheckReportStepByJobIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckReportStepByJobIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckReportStepByJobIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepByJobIdResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepByJobIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckReportStepByJobIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckReportStepByJobIdResponse) GetBody() *ListDataCheckReportStepByJobIdResponseBody {
	return s.Body
}

func (s *ListDataCheckReportStepByJobIdResponse) SetHeaders(v map[string]*string) *ListDataCheckReportStepByJobIdResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponse) SetStatusCode(v int32) *ListDataCheckReportStepByJobIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponse) SetBody(v *ListDataCheckReportStepByJobIdResponseBody) *ListDataCheckReportStepByJobIdResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
