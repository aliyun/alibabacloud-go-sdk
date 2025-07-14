// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShareReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddShareReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddShareReportResponse
	GetStatusCode() *int32
	SetBody(v *AddShareReportResponseBody) *AddShareReportResponse
	GetBody() *AddShareReportResponseBody
}

type AddShareReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShareReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddShareReportResponse) String() string {
	return dara.Prettify(s)
}

func (s AddShareReportResponse) GoString() string {
	return s.String()
}

func (s *AddShareReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddShareReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddShareReportResponse) GetBody() *AddShareReportResponseBody {
	return s.Body
}

func (s *AddShareReportResponse) SetHeaders(v map[string]*string) *AddShareReportResponse {
	s.Headers = v
	return s
}

func (s *AddShareReportResponse) SetStatusCode(v int32) *AddShareReportResponse {
	s.StatusCode = &v
	return s
}

func (s *AddShareReportResponse) SetBody(v *AddShareReportResponseBody) *AddShareReportResponse {
	s.Body = v
	return s
}

func (s *AddShareReportResponse) Validate() error {
	return dara.Validate(s)
}
