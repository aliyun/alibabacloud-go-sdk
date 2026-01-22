// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateGovernanceReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateGovernanceReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateGovernanceReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateGovernanceReportResponseBody) *GenerateGovernanceReportResponse
	GetBody() *GenerateGovernanceReportResponseBody
}

type GenerateGovernanceReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateGovernanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateGovernanceReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateGovernanceReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateGovernanceReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateGovernanceReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateGovernanceReportResponse) GetBody() *GenerateGovernanceReportResponseBody {
	return s.Body
}

func (s *GenerateGovernanceReportResponse) SetHeaders(v map[string]*string) *GenerateGovernanceReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateGovernanceReportResponse) SetStatusCode(v int32) *GenerateGovernanceReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateGovernanceReportResponse) SetBody(v *GenerateGovernanceReportResponseBody) *GenerateGovernanceReportResponse {
	s.Body = v
	return s
}

func (s *GenerateGovernanceReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
