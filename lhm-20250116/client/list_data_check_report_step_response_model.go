// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportStepResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckReportStepResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckReportStepResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckReportStepResponseBody) *ListDataCheckReportStepResponse
	GetBody() *ListDataCheckReportStepResponseBody
}

type ListDataCheckReportStepResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckReportStepResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckReportStepResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckReportStepResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckReportStepResponse) GetBody() *ListDataCheckReportStepResponseBody {
	return s.Body
}

func (s *ListDataCheckReportStepResponse) SetHeaders(v map[string]*string) *ListDataCheckReportStepResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckReportStepResponse) SetStatusCode(v int32) *ListDataCheckReportStepResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckReportStepResponse) SetBody(v *ListDataCheckReportStepResponseBody) *ListDataCheckReportStepResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckReportStepResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
