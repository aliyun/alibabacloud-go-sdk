// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEdsAgentReportConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEdsAgentReportConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEdsAgentReportConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryEdsAgentReportConfigResponseBody) *QueryEdsAgentReportConfigResponse
	GetBody() *QueryEdsAgentReportConfigResponseBody
}

type QueryEdsAgentReportConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEdsAgentReportConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEdsAgentReportConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEdsAgentReportConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEdsAgentReportConfigResponse) GetBody() *QueryEdsAgentReportConfigResponseBody {
	return s.Body
}

func (s *QueryEdsAgentReportConfigResponse) SetHeaders(v map[string]*string) *QueryEdsAgentReportConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) SetStatusCode(v int32) *QueryEdsAgentReportConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) SetBody(v *QueryEdsAgentReportConfigResponseBody) *QueryEdsAgentReportConfigResponse {
	s.Body = v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
