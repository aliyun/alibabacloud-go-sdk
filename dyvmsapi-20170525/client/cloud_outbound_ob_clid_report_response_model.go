// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudOutboundObClidReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudOutboundObClidReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudOutboundObClidReportResponse
	GetStatusCode() *int32
	SetBody(v *CloudOutboundObClidReportResponseBody) *CloudOutboundObClidReportResponse
	GetBody() *CloudOutboundObClidReportResponseBody
}

type CloudOutboundObClidReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudOutboundObClidReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudOutboundObClidReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundObClidReportResponse) GoString() string {
	return s.String()
}

func (s *CloudOutboundObClidReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudOutboundObClidReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudOutboundObClidReportResponse) GetBody() *CloudOutboundObClidReportResponseBody {
	return s.Body
}

func (s *CloudOutboundObClidReportResponse) SetHeaders(v map[string]*string) *CloudOutboundObClidReportResponse {
	s.Headers = v
	return s
}

func (s *CloudOutboundObClidReportResponse) SetStatusCode(v int32) *CloudOutboundObClidReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudOutboundObClidReportResponse) SetBody(v *CloudOutboundObClidReportResponseBody) *CloudOutboundObClidReportResponse {
	s.Body = v
	return s
}

func (s *CloudOutboundObClidReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
