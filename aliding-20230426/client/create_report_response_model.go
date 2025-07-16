// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateReportResponseBody) *CreateReportResponse
	GetBody() *CreateReportResponseBody
}

type CreateReportResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReportResponse) GoString() string {
	return s.String()
}

func (s *CreateReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReportResponse) GetBody() *CreateReportResponseBody {
	return s.Body
}

func (s *CreateReportResponse) SetHeaders(v map[string]*string) *CreateReportResponse {
	s.Headers = v
	return s
}

func (s *CreateReportResponse) SetStatusCode(v int32) *CreateReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReportResponse) SetBody(v *CreateReportResponseBody) *CreateReportResponse {
	s.Body = v
	return s
}

func (s *CreateReportResponse) Validate() error {
	return dara.Validate(s)
}
