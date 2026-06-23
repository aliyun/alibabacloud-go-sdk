// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResponseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportResponseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportResponseResponse
	GetStatusCode() *int32
	SetBody(v *GetReportResponseResponseBody) *GetReportResponseResponse
	GetBody() *GetReportResponseResponseBody
}

type GetReportResponseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportResponseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseResponse) GoString() string {
	return s.String()
}

func (s *GetReportResponseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportResponseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportResponseResponse) GetBody() *GetReportResponseResponseBody {
	return s.Body
}

func (s *GetReportResponseResponse) SetHeaders(v map[string]*string) *GetReportResponseResponse {
	s.Headers = v
	return s
}

func (s *GetReportResponseResponse) SetStatusCode(v int32) *GetReportResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportResponseResponse) SetBody(v *GetReportResponseResponseBody) *GetReportResponseResponse {
	s.Body = v
	return s
}

func (s *GetReportResponseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
