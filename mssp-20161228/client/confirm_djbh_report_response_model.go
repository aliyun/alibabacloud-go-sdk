// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDjbhReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmDjbhReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmDjbhReportResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmDjbhReportResponseBody) *ConfirmDjbhReportResponse
	GetBody() *ConfirmDjbhReportResponseBody
}

type ConfirmDjbhReportResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmDjbhReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDjbhReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDjbhReportResponse) GoString() string {
	return s.String()
}

func (s *ConfirmDjbhReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmDjbhReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmDjbhReportResponse) GetBody() *ConfirmDjbhReportResponseBody {
	return s.Body
}

func (s *ConfirmDjbhReportResponse) SetHeaders(v map[string]*string) *ConfirmDjbhReportResponse {
	s.Headers = v
	return s
}

func (s *ConfirmDjbhReportResponse) SetStatusCode(v int32) *ConfirmDjbhReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmDjbhReportResponse) SetBody(v *ConfirmDjbhReportResponseBody) *ConfirmDjbhReportResponse {
	s.Body = v
	return s
}

func (s *ConfirmDjbhReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
