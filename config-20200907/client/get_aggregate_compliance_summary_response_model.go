// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateComplianceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateComplianceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateComplianceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateComplianceSummaryResponseBody) *GetAggregateComplianceSummaryResponse
	GetBody() *GetAggregateComplianceSummaryResponseBody
}

type GetAggregateComplianceSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateComplianceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateComplianceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateComplianceSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateComplianceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateComplianceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateComplianceSummaryResponse) GetBody() *GetAggregateComplianceSummaryResponseBody {
	return s.Body
}

func (s *GetAggregateComplianceSummaryResponse) SetHeaders(v map[string]*string) *GetAggregateComplianceSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateComplianceSummaryResponse) SetStatusCode(v int32) *GetAggregateComplianceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateComplianceSummaryResponse) SetBody(v *GetAggregateComplianceSummaryResponseBody) *GetAggregateComplianceSummaryResponse {
	s.Body = v
	return s
}

func (s *GetAggregateComplianceSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
