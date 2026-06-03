// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForModifyingDomainDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForModifyingDomainDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForModifyingDomainDnsResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForModifyingDomainDnsResponseBody) *SaveBatchTaskForModifyingDomainDnsResponse
	GetBody() *SaveBatchTaskForModifyingDomainDnsResponseBody
}

type SaveBatchTaskForModifyingDomainDnsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForModifyingDomainDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForModifyingDomainDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForModifyingDomainDnsResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) GetBody() *SaveBatchTaskForModifyingDomainDnsResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForModifyingDomainDnsResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) SetStatusCode(v int32) *SaveBatchTaskForModifyingDomainDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) SetBody(v *SaveBatchTaskForModifyingDomainDnsResponseBody) *SaveBatchTaskForModifyingDomainDnsResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
