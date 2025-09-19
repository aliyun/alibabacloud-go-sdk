// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomizeReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomizeReportResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomizeReportResponseBody) *DeleteCustomizeReportResponse
	GetBody() *DeleteCustomizeReportResponseBody
}

type DeleteCustomizeReportResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizeReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeReportResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomizeReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomizeReportResponse) GetBody() *DeleteCustomizeReportResponseBody {
	return s.Body
}

func (s *DeleteCustomizeReportResponse) SetHeaders(v map[string]*string) *DeleteCustomizeReportResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizeReportResponse) SetStatusCode(v int32) *DeleteCustomizeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizeReportResponse) SetBody(v *DeleteCustomizeReportResponseBody) *DeleteCustomizeReportResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomizeReportResponse) Validate() error {
	return dara.Validate(s)
}
