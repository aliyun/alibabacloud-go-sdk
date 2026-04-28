// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDomainQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDomainQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetDomainQuotaResponseBody) *GetDomainQuotaResponse
	GetBody() *GetDomainQuotaResponseBody
}

type GetDomainQuotaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDomainQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetDomainQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDomainQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDomainQuotaResponse) GetBody() *GetDomainQuotaResponseBody {
	return s.Body
}

func (s *GetDomainQuotaResponse) SetHeaders(v map[string]*string) *GetDomainQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetDomainQuotaResponse) SetStatusCode(v int32) *GetDomainQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainQuotaResponse) SetBody(v *GetDomainQuotaResponseBody) *GetDomainQuotaResponse {
	s.Body = v
	return s
}

func (s *GetDomainQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
