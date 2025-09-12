// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainInfoForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDomainInfoForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDomainInfoForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetDomainInfoForPartnerResponseBody) *GetDomainInfoForPartnerResponse
	GetBody() *GetDomainInfoForPartnerResponseBody
}

type GetDomainInfoForPartnerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainInfoForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainInfoForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDomainInfoForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetDomainInfoForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDomainInfoForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDomainInfoForPartnerResponse) GetBody() *GetDomainInfoForPartnerResponseBody {
	return s.Body
}

func (s *GetDomainInfoForPartnerResponse) SetHeaders(v map[string]*string) *GetDomainInfoForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetDomainInfoForPartnerResponse) SetStatusCode(v int32) *GetDomainInfoForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainInfoForPartnerResponse) SetBody(v *GetDomainInfoForPartnerResponseBody) *GetDomainInfoForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetDomainInfoForPartnerResponse) Validate() error {
	return dara.Validate(s)
}
