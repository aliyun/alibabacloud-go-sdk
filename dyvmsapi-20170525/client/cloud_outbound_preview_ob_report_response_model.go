// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudOutboundPreviewObReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudOutboundPreviewObReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudOutboundPreviewObReportResponse
	GetStatusCode() *int32
	SetBody(v *CloudOutboundPreviewObReportResponseBody) *CloudOutboundPreviewObReportResponse
	GetBody() *CloudOutboundPreviewObReportResponseBody
}

type CloudOutboundPreviewObReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudOutboundPreviewObReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudOutboundPreviewObReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundPreviewObReportResponse) GoString() string {
	return s.String()
}

func (s *CloudOutboundPreviewObReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudOutboundPreviewObReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudOutboundPreviewObReportResponse) GetBody() *CloudOutboundPreviewObReportResponseBody {
	return s.Body
}

func (s *CloudOutboundPreviewObReportResponse) SetHeaders(v map[string]*string) *CloudOutboundPreviewObReportResponse {
	s.Headers = v
	return s
}

func (s *CloudOutboundPreviewObReportResponse) SetStatusCode(v int32) *CloudOutboundPreviewObReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponse) SetBody(v *CloudOutboundPreviewObReportResponseBody) *CloudOutboundPreviewObReportResponse {
	s.Body = v
	return s
}

func (s *CloudOutboundPreviewObReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
