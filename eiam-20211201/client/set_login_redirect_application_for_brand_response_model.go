// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoginRedirectApplicationForBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoginRedirectApplicationForBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoginRedirectApplicationForBrandResponse
	GetStatusCode() *int32
	SetBody(v *SetLoginRedirectApplicationForBrandResponseBody) *SetLoginRedirectApplicationForBrandResponse
	GetBody() *SetLoginRedirectApplicationForBrandResponseBody
}

type SetLoginRedirectApplicationForBrandResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoginRedirectApplicationForBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoginRedirectApplicationForBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoginRedirectApplicationForBrandResponse) GoString() string {
	return s.String()
}

func (s *SetLoginRedirectApplicationForBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoginRedirectApplicationForBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoginRedirectApplicationForBrandResponse) GetBody() *SetLoginRedirectApplicationForBrandResponseBody {
	return s.Body
}

func (s *SetLoginRedirectApplicationForBrandResponse) SetHeaders(v map[string]*string) *SetLoginRedirectApplicationForBrandResponse {
	s.Headers = v
	return s
}

func (s *SetLoginRedirectApplicationForBrandResponse) SetStatusCode(v int32) *SetLoginRedirectApplicationForBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoginRedirectApplicationForBrandResponse) SetBody(v *SetLoginRedirectApplicationForBrandResponseBody) *SetLoginRedirectApplicationForBrandResponse {
	s.Body = v
	return s
}

func (s *SetLoginRedirectApplicationForBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
