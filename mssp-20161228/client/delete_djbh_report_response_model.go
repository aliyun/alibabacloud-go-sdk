// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDjbhReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDjbhReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDjbhReportResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDjbhReportResponseBody) *DeleteDjbhReportResponse
	GetBody() *DeleteDjbhReportResponseBody
}

type DeleteDjbhReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDjbhReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDjbhReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDjbhReportResponse) GoString() string {
	return s.String()
}

func (s *DeleteDjbhReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDjbhReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDjbhReportResponse) GetBody() *DeleteDjbhReportResponseBody {
	return s.Body
}

func (s *DeleteDjbhReportResponse) SetHeaders(v map[string]*string) *DeleteDjbhReportResponse {
	s.Headers = v
	return s
}

func (s *DeleteDjbhReportResponse) SetStatusCode(v int32) *DeleteDjbhReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDjbhReportResponse) SetBody(v *DeleteDjbhReportResponseBody) *DeleteDjbhReportResponse {
	s.Body = v
	return s
}

func (s *DeleteDjbhReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
