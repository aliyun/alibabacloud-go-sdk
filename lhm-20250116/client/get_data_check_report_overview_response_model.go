// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckReportOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckReportOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckReportOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckReportOverviewResponseBody) *GetDataCheckReportOverviewResponse
	GetBody() *GetDataCheckReportOverviewResponseBody
}

type GetDataCheckReportOverviewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckReportOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckReportOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckReportOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckReportOverviewResponse) GetBody() *GetDataCheckReportOverviewResponseBody {
	return s.Body
}

func (s *GetDataCheckReportOverviewResponse) SetHeaders(v map[string]*string) *GetDataCheckReportOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckReportOverviewResponse) SetStatusCode(v int32) *GetDataCheckReportOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckReportOverviewResponse) SetBody(v *GetDataCheckReportOverviewResponseBody) *GetDataCheckReportOverviewResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckReportOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
