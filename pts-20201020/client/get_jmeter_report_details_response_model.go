// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterReportDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJMeterReportDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJMeterReportDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetJMeterReportDetailsResponseBody) *GetJMeterReportDetailsResponse
	GetBody() *GetJMeterReportDetailsResponseBody
}

type GetJMeterReportDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJMeterReportDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJMeterReportDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterReportDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterReportDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJMeterReportDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJMeterReportDetailsResponse) GetBody() *GetJMeterReportDetailsResponseBody {
	return s.Body
}

func (s *GetJMeterReportDetailsResponse) SetHeaders(v map[string]*string) *GetJMeterReportDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterReportDetailsResponse) SetStatusCode(v int32) *GetJMeterReportDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJMeterReportDetailsResponse) SetBody(v *GetJMeterReportDetailsResponseBody) *GetJMeterReportDetailsResponse {
	s.Body = v
	return s
}

func (s *GetJMeterReportDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
