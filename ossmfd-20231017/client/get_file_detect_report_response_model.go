// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileDetectReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileDetectReportResponse
	GetStatusCode() *int32
	SetBody(v *GetFileDetectReportResponseBody) *GetFileDetectReportResponse
	GetBody() *GetFileDetectReportResponseBody
}

type GetFileDetectReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileDetectReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileDetectReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectReportResponse) GoString() string {
	return s.String()
}

func (s *GetFileDetectReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileDetectReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileDetectReportResponse) GetBody() *GetFileDetectReportResponseBody {
	return s.Body
}

func (s *GetFileDetectReportResponse) SetHeaders(v map[string]*string) *GetFileDetectReportResponse {
	s.Headers = v
	return s
}

func (s *GetFileDetectReportResponse) SetStatusCode(v int32) *GetFileDetectReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDetectReportResponse) SetBody(v *GetFileDetectReportResponseBody) *GetFileDetectReportResponse {
	s.Body = v
	return s
}

func (s *GetFileDetectReportResponse) Validate() error {
	return dara.Validate(s)
}
