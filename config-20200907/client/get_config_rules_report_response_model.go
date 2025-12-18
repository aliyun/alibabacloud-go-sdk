// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRulesReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigRulesReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigRulesReportResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigRulesReportResponseBody) *GetConfigRulesReportResponse
	GetBody() *GetConfigRulesReportResponseBody
}

type GetConfigRulesReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigRulesReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigRulesReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigRulesReportResponse) GetBody() *GetConfigRulesReportResponseBody {
	return s.Body
}

func (s *GetConfigRulesReportResponse) SetHeaders(v map[string]*string) *GetConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRulesReportResponse) SetStatusCode(v int32) *GetConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRulesReportResponse) SetBody(v *GetConfigRulesReportResponseBody) *GetConfigRulesReportResponse {
	s.Body = v
	return s
}

func (s *GetConfigRulesReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
