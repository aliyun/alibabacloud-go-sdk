// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportDefinitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReportDefinitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReportDefinitionsResponse
	GetStatusCode() *int32
	SetBody(v *ListReportDefinitionsResponseBody) *ListReportDefinitionsResponse
	GetBody() *ListReportDefinitionsResponseBody
}

type ListReportDefinitionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReportDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReportDefinitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReportDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReportDefinitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReportDefinitionsResponse) GetBody() *ListReportDefinitionsResponseBody {
	return s.Body
}

func (s *ListReportDefinitionsResponse) SetHeaders(v map[string]*string) *ListReportDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListReportDefinitionsResponse) SetStatusCode(v int32) *ListReportDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReportDefinitionsResponse) SetBody(v *ListReportDefinitionsResponseBody) *ListReportDefinitionsResponse {
	s.Body = v
	return s
}

func (s *ListReportDefinitionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
