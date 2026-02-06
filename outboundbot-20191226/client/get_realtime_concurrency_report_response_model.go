// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeConcurrencyReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealtimeConcurrencyReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealtimeConcurrencyReportResponse
	GetStatusCode() *int32
	SetBody(v *GetRealtimeConcurrencyReportResponseBody) *GetRealtimeConcurrencyReportResponse
	GetBody() *GetRealtimeConcurrencyReportResponseBody
}

type GetRealtimeConcurrencyReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeConcurrencyReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeConcurrencyReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeConcurrencyReportResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeConcurrencyReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealtimeConcurrencyReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealtimeConcurrencyReportResponse) GetBody() *GetRealtimeConcurrencyReportResponseBody {
	return s.Body
}

func (s *GetRealtimeConcurrencyReportResponse) SetHeaders(v map[string]*string) *GetRealtimeConcurrencyReportResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeConcurrencyReportResponse) SetStatusCode(v int32) *GetRealtimeConcurrencyReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeConcurrencyReportResponse) SetBody(v *GetRealtimeConcurrencyReportResponseBody) *GetRealtimeConcurrencyReportResponse {
	s.Body = v
	return s
}

func (s *GetRealtimeConcurrencyReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
