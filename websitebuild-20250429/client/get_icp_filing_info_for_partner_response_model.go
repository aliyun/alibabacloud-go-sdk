// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIcpFilingInfoForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIcpFilingInfoForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIcpFilingInfoForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetIcpFilingInfoForPartnerResponseBody) *GetIcpFilingInfoForPartnerResponse
	GetBody() *GetIcpFilingInfoForPartnerResponseBody
}

type GetIcpFilingInfoForPartnerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIcpFilingInfoForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIcpFilingInfoForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIcpFilingInfoForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetIcpFilingInfoForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIcpFilingInfoForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIcpFilingInfoForPartnerResponse) GetBody() *GetIcpFilingInfoForPartnerResponseBody {
	return s.Body
}

func (s *GetIcpFilingInfoForPartnerResponse) SetHeaders(v map[string]*string) *GetIcpFilingInfoForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponse) SetStatusCode(v int32) *GetIcpFilingInfoForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponse) SetBody(v *GetIcpFilingInfoForPartnerResponseBody) *GetIcpFilingInfoForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetIcpFilingInfoForPartnerResponse) Validate() error {
	return dara.Validate(s)
}
