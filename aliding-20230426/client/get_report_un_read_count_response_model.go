// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportUnReadCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportUnReadCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportUnReadCountResponse
	GetStatusCode() *int32
	SetBody(v *GetReportUnReadCountResponseBody) *GetReportUnReadCountResponse
	GetBody() *GetReportUnReadCountResponseBody
}

type GetReportUnReadCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportUnReadCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportUnReadCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountResponse) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportUnReadCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportUnReadCountResponse) GetBody() *GetReportUnReadCountResponseBody {
	return s.Body
}

func (s *GetReportUnReadCountResponse) SetHeaders(v map[string]*string) *GetReportUnReadCountResponse {
	s.Headers = v
	return s
}

func (s *GetReportUnReadCountResponse) SetStatusCode(v int32) *GetReportUnReadCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportUnReadCountResponse) SetBody(v *GetReportUnReadCountResponseBody) *GetReportUnReadCountResponse {
	s.Body = v
	return s
}

func (s *GetReportUnReadCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
