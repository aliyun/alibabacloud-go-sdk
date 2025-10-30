// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFixedNoAReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFixedNoAReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFixedNoAReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateFixedNoAReportResponseBody) *CreateFixedNoAReportResponse
	GetBody() *CreateFixedNoAReportResponseBody
}

type CreateFixedNoAReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFixedNoAReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFixedNoAReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFixedNoAReportResponse) GoString() string {
	return s.String()
}

func (s *CreateFixedNoAReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFixedNoAReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFixedNoAReportResponse) GetBody() *CreateFixedNoAReportResponseBody {
	return s.Body
}

func (s *CreateFixedNoAReportResponse) SetHeaders(v map[string]*string) *CreateFixedNoAReportResponse {
	s.Headers = v
	return s
}

func (s *CreateFixedNoAReportResponse) SetStatusCode(v int32) *CreateFixedNoAReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFixedNoAReportResponse) SetBody(v *CreateFixedNoAReportResponseBody) *CreateFixedNoAReportResponse {
	s.Body = v
	return s
}

func (s *CreateFixedNoAReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
