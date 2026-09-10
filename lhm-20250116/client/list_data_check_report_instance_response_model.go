// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckReportInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckReportInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckReportInstanceResponseBody) *ListDataCheckReportInstanceResponse
	GetBody() *ListDataCheckReportInstanceResponseBody
}

type ListDataCheckReportInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckReportInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckReportInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckReportInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckReportInstanceResponse) GetBody() *ListDataCheckReportInstanceResponseBody {
	return s.Body
}

func (s *ListDataCheckReportInstanceResponse) SetHeaders(v map[string]*string) *ListDataCheckReportInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckReportInstanceResponse) SetStatusCode(v int32) *ListDataCheckReportInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckReportInstanceResponse) SetBody(v *ListDataCheckReportInstanceResponseBody) *ListDataCheckReportInstanceResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckReportInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
