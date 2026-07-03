// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationSecurityDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationSecurityDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationSecurityDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationSecurityDomainsResponseBody) *ListNormalizationSecurityDomainsResponse
	GetBody() *ListNormalizationSecurityDomainsResponseBody
}

type ListNormalizationSecurityDomainsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationSecurityDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationSecurityDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSecurityDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationSecurityDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationSecurityDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationSecurityDomainsResponse) GetBody() *ListNormalizationSecurityDomainsResponseBody {
	return s.Body
}

func (s *ListNormalizationSecurityDomainsResponse) SetHeaders(v map[string]*string) *ListNormalizationSecurityDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationSecurityDomainsResponse) SetStatusCode(v int32) *ListNormalizationSecurityDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponse) SetBody(v *ListNormalizationSecurityDomainsResponseBody) *ListNormalizationSecurityDomainsResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationSecurityDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
