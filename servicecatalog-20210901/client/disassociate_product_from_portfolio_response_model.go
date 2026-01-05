// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateProductFromPortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateProductFromPortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateProductFromPortfolioResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateProductFromPortfolioResponseBody) *DisassociateProductFromPortfolioResponse
	GetBody() *DisassociateProductFromPortfolioResponseBody
}

type DisassociateProductFromPortfolioResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateProductFromPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateProductFromPortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateProductFromPortfolioResponse) GoString() string {
	return s.String()
}

func (s *DisassociateProductFromPortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateProductFromPortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateProductFromPortfolioResponse) GetBody() *DisassociateProductFromPortfolioResponseBody {
	return s.Body
}

func (s *DisassociateProductFromPortfolioResponse) SetHeaders(v map[string]*string) *DisassociateProductFromPortfolioResponse {
	s.Headers = v
	return s
}

func (s *DisassociateProductFromPortfolioResponse) SetStatusCode(v int32) *DisassociateProductFromPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateProductFromPortfolioResponse) SetBody(v *DisassociateProductFromPortfolioResponseBody) *DisassociateProductFromPortfolioResponse {
	s.Body = v
	return s
}

func (s *DisassociateProductFromPortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
