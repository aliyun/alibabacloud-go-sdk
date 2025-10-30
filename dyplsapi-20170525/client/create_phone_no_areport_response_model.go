// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhoneNoAReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePhoneNoAReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePhoneNoAReportResponse
	GetStatusCode() *int32
	SetBody(v *CreatePhoneNoAReportResponseBody) *CreatePhoneNoAReportResponse
	GetBody() *CreatePhoneNoAReportResponseBody
}

type CreatePhoneNoAReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhoneNoAReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhoneNoAReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneNoAReportResponse) GoString() string {
	return s.String()
}

func (s *CreatePhoneNoAReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePhoneNoAReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePhoneNoAReportResponse) GetBody() *CreatePhoneNoAReportResponseBody {
	return s.Body
}

func (s *CreatePhoneNoAReportResponse) SetHeaders(v map[string]*string) *CreatePhoneNoAReportResponse {
	s.Headers = v
	return s
}

func (s *CreatePhoneNoAReportResponse) SetStatusCode(v int32) *CreatePhoneNoAReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhoneNoAReportResponse) SetBody(v *CreatePhoneNoAReportResponseBody) *CreatePhoneNoAReportResponse {
	s.Body = v
	return s
}

func (s *CreatePhoneNoAReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
