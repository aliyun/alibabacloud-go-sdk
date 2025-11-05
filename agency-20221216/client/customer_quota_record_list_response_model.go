// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerQuotaRecordListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerQuotaRecordListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerQuotaRecordListResponse
	GetStatusCode() *int32
	SetBody(v *CustomerQuotaRecordListResponseBody) *CustomerQuotaRecordListResponse
	GetBody() *CustomerQuotaRecordListResponseBody
}

type CustomerQuotaRecordListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerQuotaRecordListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerQuotaRecordListResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerQuotaRecordListResponse) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerQuotaRecordListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerQuotaRecordListResponse) GetBody() *CustomerQuotaRecordListResponseBody {
	return s.Body
}

func (s *CustomerQuotaRecordListResponse) SetHeaders(v map[string]*string) *CustomerQuotaRecordListResponse {
	s.Headers = v
	return s
}

func (s *CustomerQuotaRecordListResponse) SetStatusCode(v int32) *CustomerQuotaRecordListResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerQuotaRecordListResponse) SetBody(v *CustomerQuotaRecordListResponseBody) *CustomerQuotaRecordListResponse {
	s.Body = v
	return s
}

func (s *CustomerQuotaRecordListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
