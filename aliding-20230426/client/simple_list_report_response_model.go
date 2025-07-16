// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleListReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SimpleListReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SimpleListReportResponse
	GetStatusCode() *int32
	SetBody(v *SimpleListReportResponseBody) *SimpleListReportResponse
	GetBody() *SimpleListReportResponseBody
}

type SimpleListReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SimpleListReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SimpleListReportResponse) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportResponse) GoString() string {
	return s.String()
}

func (s *SimpleListReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SimpleListReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SimpleListReportResponse) GetBody() *SimpleListReportResponseBody {
	return s.Body
}

func (s *SimpleListReportResponse) SetHeaders(v map[string]*string) *SimpleListReportResponse {
	s.Headers = v
	return s
}

func (s *SimpleListReportResponse) SetStatusCode(v int32) *SimpleListReportResponse {
	s.StatusCode = &v
	return s
}

func (s *SimpleListReportResponse) SetBody(v *SimpleListReportResponseBody) *SimpleListReportResponse {
	s.Body = v
	return s
}

func (s *SimpleListReportResponse) Validate() error {
	return dara.Validate(s)
}
