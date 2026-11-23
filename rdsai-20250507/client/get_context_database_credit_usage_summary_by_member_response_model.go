// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseCreditUsageSummaryByMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseCreditUsageSummaryByMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseCreditUsageSummaryByMemberResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) *GetContextDatabaseCreditUsageSummaryByMemberResponse
	GetBody() *GetContextDatabaseCreditUsageSummaryByMemberResponseBody
}

type GetContextDatabaseCreditUsageSummaryByMemberResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseCreditUsageSummaryByMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseCreditUsageSummaryByMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseCreditUsageSummaryByMemberResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) GetBody() *GetContextDatabaseCreditUsageSummaryByMemberResponseBody {
	return s.Body
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) SetHeaders(v map[string]*string) *GetContextDatabaseCreditUsageSummaryByMemberResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) SetStatusCode(v int32) *GetContextDatabaseCreditUsageSummaryByMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) SetBody(v *GetContextDatabaseCreditUsageSummaryByMemberResponseBody) *GetContextDatabaseCreditUsageSummaryByMemberResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseCreditUsageSummaryByMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
