// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckReportStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckReportStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckReportStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckReportStatusResponseBody) *GetDataCheckReportStatusResponse
	GetBody() *GetDataCheckReportStatusResponseBody
}

type GetDataCheckReportStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckReportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckReportStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckReportStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckReportStatusResponse) GetBody() *GetDataCheckReportStatusResponseBody {
	return s.Body
}

func (s *GetDataCheckReportStatusResponse) SetHeaders(v map[string]*string) *GetDataCheckReportStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckReportStatusResponse) SetStatusCode(v int32) *GetDataCheckReportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckReportStatusResponse) SetBody(v *GetDataCheckReportStatusResponseBody) *GetDataCheckReportStatusResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckReportStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
