// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceItemReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGovernanceItemReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGovernanceItemReportResponse
	GetStatusCode() *int32
	SetBody(v *GetGovernanceItemReportResponseBody) *GetGovernanceItemReportResponse
	GetBody() *GetGovernanceItemReportResponseBody
}

type GetGovernanceItemReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGovernanceItemReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGovernanceItemReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGovernanceItemReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGovernanceItemReportResponse) GetBody() *GetGovernanceItemReportResponseBody {
	return s.Body
}

func (s *GetGovernanceItemReportResponse) SetHeaders(v map[string]*string) *GetGovernanceItemReportResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceItemReportResponse) SetStatusCode(v int32) *GetGovernanceItemReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGovernanceItemReportResponse) SetBody(v *GetGovernanceItemReportResponseBody) *GetGovernanceItemReportResponse {
	s.Body = v
	return s
}

func (s *GetGovernanceItemReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
