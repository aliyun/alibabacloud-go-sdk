// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachTaskSessionReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachTaskSessionReportResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachTaskSessionReportResponseBody) *GetAICoachTaskSessionReportResponse
	GetBody() *GetAICoachTaskSessionReportResponseBody
}

type GetAICoachTaskSessionReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachTaskSessionReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachTaskSessionReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionReportResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachTaskSessionReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachTaskSessionReportResponse) GetBody() *GetAICoachTaskSessionReportResponseBody {
	return s.Body
}

func (s *GetAICoachTaskSessionReportResponse) SetHeaders(v map[string]*string) *GetAICoachTaskSessionReportResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachTaskSessionReportResponse) SetStatusCode(v int32) *GetAICoachTaskSessionReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponse) SetBody(v *GetAICoachTaskSessionReportResponseBody) *GetAICoachTaskSessionReportResponse {
	s.Body = v
	return s
}

func (s *GetAICoachTaskSessionReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
