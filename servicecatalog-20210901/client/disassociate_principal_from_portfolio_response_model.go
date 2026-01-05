// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociatePrincipalFromPortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociatePrincipalFromPortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociatePrincipalFromPortfolioResponse
	GetStatusCode() *int32
	SetBody(v *DisassociatePrincipalFromPortfolioResponseBody) *DisassociatePrincipalFromPortfolioResponse
	GetBody() *DisassociatePrincipalFromPortfolioResponseBody
}

type DisassociatePrincipalFromPortfolioResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociatePrincipalFromPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociatePrincipalFromPortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociatePrincipalFromPortfolioResponse) GoString() string {
	return s.String()
}

func (s *DisassociatePrincipalFromPortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociatePrincipalFromPortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociatePrincipalFromPortfolioResponse) GetBody() *DisassociatePrincipalFromPortfolioResponseBody {
	return s.Body
}

func (s *DisassociatePrincipalFromPortfolioResponse) SetHeaders(v map[string]*string) *DisassociatePrincipalFromPortfolioResponse {
	s.Headers = v
	return s
}

func (s *DisassociatePrincipalFromPortfolioResponse) SetStatusCode(v int32) *DisassociatePrincipalFromPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioResponse) SetBody(v *DisassociatePrincipalFromPortfolioResponseBody) *DisassociatePrincipalFromPortfolioResponse {
	s.Body = v
	return s
}

func (s *DisassociatePrincipalFromPortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
