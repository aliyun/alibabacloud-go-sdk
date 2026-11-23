// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseCreditUsageRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseCreditUsageRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseCreditUsageRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseCreditUsageRecordsResponseBody) *ListContextDatabaseCreditUsageRecordsResponse
	GetBody() *ListContextDatabaseCreditUsageRecordsResponseBody
}

type ListContextDatabaseCreditUsageRecordsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseCreditUsageRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseCreditUsageRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseCreditUsageRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) GetBody() *ListContextDatabaseCreditUsageRecordsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseCreditUsageRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) SetStatusCode(v int32) *ListContextDatabaseCreditUsageRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) SetBody(v *ListContextDatabaseCreditUsageRecordsResponseBody) *ListContextDatabaseCreditUsageRecordsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseCreditUsageRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
