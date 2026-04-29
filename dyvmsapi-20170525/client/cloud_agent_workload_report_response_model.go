// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentWorkloadReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAgentWorkloadReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAgentWorkloadReportResponse
	GetStatusCode() *int32
	SetBody(v *CloudAgentWorkloadReportResponseBody) *CloudAgentWorkloadReportResponse
	GetBody() *CloudAgentWorkloadReportResponseBody
}

type CloudAgentWorkloadReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAgentWorkloadReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAgentWorkloadReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentWorkloadReportResponse) GoString() string {
	return s.String()
}

func (s *CloudAgentWorkloadReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAgentWorkloadReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAgentWorkloadReportResponse) GetBody() *CloudAgentWorkloadReportResponseBody {
	return s.Body
}

func (s *CloudAgentWorkloadReportResponse) SetHeaders(v map[string]*string) *CloudAgentWorkloadReportResponse {
	s.Headers = v
	return s
}

func (s *CloudAgentWorkloadReportResponse) SetStatusCode(v int32) *CloudAgentWorkloadReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAgentWorkloadReportResponse) SetBody(v *CloudAgentWorkloadReportResponseBody) *CloudAgentWorkloadReportResponse {
	s.Body = v
	return s
}

func (s *CloudAgentWorkloadReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
