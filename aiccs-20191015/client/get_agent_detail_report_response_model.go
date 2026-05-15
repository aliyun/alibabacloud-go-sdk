// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDetailReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentDetailReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentDetailReportResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentDetailReportResponseBody) *GetAgentDetailReportResponse
	GetBody() *GetAgentDetailReportResponseBody
}

type GetAgentDetailReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentDetailReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentDetailReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentDetailReportResponse) GetBody() *GetAgentDetailReportResponseBody {
	return s.Body
}

func (s *GetAgentDetailReportResponse) SetHeaders(v map[string]*string) *GetAgentDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetAgentDetailReportResponse) SetStatusCode(v int32) *GetAgentDetailReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentDetailReportResponse) SetBody(v *GetAgentDetailReportResponseBody) *GetAgentDetailReportResponse {
	s.Body = v
	return s
}

func (s *GetAgentDetailReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
