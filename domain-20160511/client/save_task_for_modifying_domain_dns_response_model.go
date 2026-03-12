// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForModifyingDomainDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForModifyingDomainDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForModifyingDomainDnsResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForModifyingDomainDnsResponseBody) *SaveTaskForModifyingDomainDnsResponse
	GetBody() *SaveTaskForModifyingDomainDnsResponseBody
}

type SaveTaskForModifyingDomainDnsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForModifyingDomainDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForModifyingDomainDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForModifyingDomainDnsResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForModifyingDomainDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForModifyingDomainDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForModifyingDomainDnsResponse) GetBody() *SaveTaskForModifyingDomainDnsResponseBody {
	return s.Body
}

func (s *SaveTaskForModifyingDomainDnsResponse) SetHeaders(v map[string]*string) *SaveTaskForModifyingDomainDnsResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForModifyingDomainDnsResponse) SetStatusCode(v int32) *SaveTaskForModifyingDomainDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsResponse) SetBody(v *SaveTaskForModifyingDomainDnsResponseBody) *SaveTaskForModifyingDomainDnsResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForModifyingDomainDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
