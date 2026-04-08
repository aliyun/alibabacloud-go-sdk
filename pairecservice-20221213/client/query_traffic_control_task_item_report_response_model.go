// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTaskItemReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTrafficControlTaskItemReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTrafficControlTaskItemReportResponse
	GetStatusCode() *int32
	SetBody(v *QueryTrafficControlTaskItemReportResponseBody) *QueryTrafficControlTaskItemReportResponse
	GetBody() *QueryTrafficControlTaskItemReportResponseBody
}

type QueryTrafficControlTaskItemReportResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTrafficControlTaskItemReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTrafficControlTaskItemReportResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskItemReportResponse) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskItemReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTrafficControlTaskItemReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTrafficControlTaskItemReportResponse) GetBody() *QueryTrafficControlTaskItemReportResponseBody {
	return s.Body
}

func (s *QueryTrafficControlTaskItemReportResponse) SetHeaders(v map[string]*string) *QueryTrafficControlTaskItemReportResponse {
	s.Headers = v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponse) SetStatusCode(v int32) *QueryTrafficControlTaskItemReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponse) SetBody(v *QueryTrafficControlTaskItemReportResponseBody) *QueryTrafficControlTaskItemReportResponse {
	s.Body = v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
