// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentDetailReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineAgentDetailReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineAgentDetailReportResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineAgentDetailReportResponseBody) *GetHotlineAgentDetailReportResponse
	GetBody() *GetHotlineAgentDetailReportResponseBody
}

type GetHotlineAgentDetailReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineAgentDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineAgentDetailReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineAgentDetailReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineAgentDetailReportResponse) GetBody() *GetHotlineAgentDetailReportResponseBody {
	return s.Body
}

func (s *GetHotlineAgentDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) SetStatusCode(v int32) *GetHotlineAgentDetailReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) SetBody(v *GetHotlineAgentDetailReportResponseBody) *GetHotlineAgentDetailReportResponse {
	s.Body = v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
