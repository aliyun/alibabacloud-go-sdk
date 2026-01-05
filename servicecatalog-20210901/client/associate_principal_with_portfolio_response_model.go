// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePrincipalWithPortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociatePrincipalWithPortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociatePrincipalWithPortfolioResponse
	GetStatusCode() *int32
	SetBody(v *AssociatePrincipalWithPortfolioResponseBody) *AssociatePrincipalWithPortfolioResponse
	GetBody() *AssociatePrincipalWithPortfolioResponseBody
}

type AssociatePrincipalWithPortfolioResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociatePrincipalWithPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociatePrincipalWithPortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociatePrincipalWithPortfolioResponse) GoString() string {
	return s.String()
}

func (s *AssociatePrincipalWithPortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociatePrincipalWithPortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociatePrincipalWithPortfolioResponse) GetBody() *AssociatePrincipalWithPortfolioResponseBody {
	return s.Body
}

func (s *AssociatePrincipalWithPortfolioResponse) SetHeaders(v map[string]*string) *AssociatePrincipalWithPortfolioResponse {
	s.Headers = v
	return s
}

func (s *AssociatePrincipalWithPortfolioResponse) SetStatusCode(v int32) *AssociatePrincipalWithPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioResponse) SetBody(v *AssociatePrincipalWithPortfolioResponseBody) *AssociatePrincipalWithPortfolioResponse {
	s.Body = v
	return s
}

func (s *AssociatePrincipalWithPortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
