// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComplianceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComplianceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComplianceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetComplianceSummaryResponseBody) *GetComplianceSummaryResponse
	GetBody() *GetComplianceSummaryResponseBody
}

type GetComplianceSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComplianceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComplianceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComplianceSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetComplianceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComplianceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComplianceSummaryResponse) GetBody() *GetComplianceSummaryResponseBody {
	return s.Body
}

func (s *GetComplianceSummaryResponse) SetHeaders(v map[string]*string) *GetComplianceSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetComplianceSummaryResponse) SetStatusCode(v int32) *GetComplianceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComplianceSummaryResponse) SetBody(v *GetComplianceSummaryResponseBody) *GetComplianceSummaryResponse {
	s.Body = v
	return s
}

func (s *GetComplianceSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
