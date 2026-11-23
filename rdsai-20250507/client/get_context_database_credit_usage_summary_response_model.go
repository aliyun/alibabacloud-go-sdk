// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseCreditUsageSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseCreditUsageSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseCreditUsageSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseCreditUsageSummaryResponseBody) *GetContextDatabaseCreditUsageSummaryResponse
	GetBody() *GetContextDatabaseCreditUsageSummaryResponseBody
}

type GetContextDatabaseCreditUsageSummaryResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseCreditUsageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) GetBody() *GetContextDatabaseCreditUsageSummaryResponseBody {
	return s.Body
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) SetHeaders(v map[string]*string) *GetContextDatabaseCreditUsageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) SetStatusCode(v int32) *GetContextDatabaseCreditUsageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) SetBody(v *GetContextDatabaseCreditUsageSummaryResponseBody) *GetContextDatabaseCreditUsageSummaryResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
