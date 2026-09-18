// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllRulesSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllRulesSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllRulesSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetAllRulesSummaryResponseBody) *GetAllRulesSummaryResponse
	GetBody() *GetAllRulesSummaryResponseBody
}

type GetAllRulesSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllRulesSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllRulesSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllRulesSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAllRulesSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllRulesSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllRulesSummaryResponse) GetBody() *GetAllRulesSummaryResponseBody {
	return s.Body
}

func (s *GetAllRulesSummaryResponse) SetHeaders(v map[string]*string) *GetAllRulesSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAllRulesSummaryResponse) SetStatusCode(v int32) *GetAllRulesSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllRulesSummaryResponse) SetBody(v *GetAllRulesSummaryResponseBody) *GetAllRulesSummaryResponse {
	s.Body = v
	return s
}

func (s *GetAllRulesSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
