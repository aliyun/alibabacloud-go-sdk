// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginRedirectApplicationForBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoginRedirectApplicationForBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoginRedirectApplicationForBrandResponse
	GetStatusCode() *int32
	SetBody(v *GetLoginRedirectApplicationForBrandResponseBody) *GetLoginRedirectApplicationForBrandResponse
	GetBody() *GetLoginRedirectApplicationForBrandResponseBody
}

type GetLoginRedirectApplicationForBrandResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginRedirectApplicationForBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginRedirectApplicationForBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoginRedirectApplicationForBrandResponse) GoString() string {
	return s.String()
}

func (s *GetLoginRedirectApplicationForBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoginRedirectApplicationForBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoginRedirectApplicationForBrandResponse) GetBody() *GetLoginRedirectApplicationForBrandResponseBody {
	return s.Body
}

func (s *GetLoginRedirectApplicationForBrandResponse) SetHeaders(v map[string]*string) *GetLoginRedirectApplicationForBrandResponse {
	s.Headers = v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponse) SetStatusCode(v int32) *GetLoginRedirectApplicationForBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponse) SetBody(v *GetLoginRedirectApplicationForBrandResponseBody) *GetLoginRedirectApplicationForBrandResponse {
	s.Body = v
	return s
}

func (s *GetLoginRedirectApplicationForBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
