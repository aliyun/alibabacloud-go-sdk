// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceReportStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGovernanceReportStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGovernanceReportStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetGovernanceReportStatusResponseBody) *GetGovernanceReportStatusResponse
	GetBody() *GetGovernanceReportStatusResponseBody
}

type GetGovernanceReportStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGovernanceReportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGovernanceReportStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceReportStatusResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGovernanceReportStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGovernanceReportStatusResponse) GetBody() *GetGovernanceReportStatusResponseBody {
	return s.Body
}

func (s *GetGovernanceReportStatusResponse) SetHeaders(v map[string]*string) *GetGovernanceReportStatusResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceReportStatusResponse) SetStatusCode(v int32) *GetGovernanceReportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGovernanceReportStatusResponse) SetBody(v *GetGovernanceReportStatusResponseBody) *GetGovernanceReportStatusResponse {
	s.Body = v
	return s
}

func (s *GetGovernanceReportStatusResponse) Validate() error {
	return dara.Validate(s)
}
