// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomizeReportStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomizeReportStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomizeReportStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomizeReportStatusResponseBody) *UpdateCustomizeReportStatusResponse
	GetBody() *UpdateCustomizeReportStatusResponseBody
}

type UpdateCustomizeReportStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomizeReportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomizeReportStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomizeReportStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomizeReportStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomizeReportStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomizeReportStatusResponse) GetBody() *UpdateCustomizeReportStatusResponseBody {
	return s.Body
}

func (s *UpdateCustomizeReportStatusResponse) SetHeaders(v map[string]*string) *UpdateCustomizeReportStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomizeReportStatusResponse) SetStatusCode(v int32) *UpdateCustomizeReportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomizeReportStatusResponse) SetBody(v *UpdateCustomizeReportStatusResponseBody) *UpdateCustomizeReportStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomizeReportStatusResponse) Validate() error {
	return dara.Validate(s)
}
