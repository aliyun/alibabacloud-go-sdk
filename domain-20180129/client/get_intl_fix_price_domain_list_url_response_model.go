// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntlFixPriceDomainListUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntlFixPriceDomainListUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntlFixPriceDomainListUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetIntlFixPriceDomainListUrlResponseBody) *GetIntlFixPriceDomainListUrlResponse
	GetBody() *GetIntlFixPriceDomainListUrlResponseBody
}

type GetIntlFixPriceDomainListUrlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntlFixPriceDomainListUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntlFixPriceDomainListUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntlFixPriceDomainListUrlResponse) GoString() string {
	return s.String()
}

func (s *GetIntlFixPriceDomainListUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntlFixPriceDomainListUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntlFixPriceDomainListUrlResponse) GetBody() *GetIntlFixPriceDomainListUrlResponseBody {
	return s.Body
}

func (s *GetIntlFixPriceDomainListUrlResponse) SetHeaders(v map[string]*string) *GetIntlFixPriceDomainListUrlResponse {
	s.Headers = v
	return s
}

func (s *GetIntlFixPriceDomainListUrlResponse) SetStatusCode(v int32) *GetIntlFixPriceDomainListUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntlFixPriceDomainListUrlResponse) SetBody(v *GetIntlFixPriceDomainListUrlResponseBody) *GetIntlFixPriceDomainListUrlResponse {
	s.Body = v
	return s
}

func (s *GetIntlFixPriceDomainListUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
