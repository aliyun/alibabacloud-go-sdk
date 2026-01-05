// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProductWithPortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateProductWithPortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateProductWithPortfolioResponse
	GetStatusCode() *int32
	SetBody(v *AssociateProductWithPortfolioResponseBody) *AssociateProductWithPortfolioResponse
	GetBody() *AssociateProductWithPortfolioResponseBody
}

type AssociateProductWithPortfolioResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateProductWithPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateProductWithPortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateProductWithPortfolioResponse) GoString() string {
	return s.String()
}

func (s *AssociateProductWithPortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateProductWithPortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateProductWithPortfolioResponse) GetBody() *AssociateProductWithPortfolioResponseBody {
	return s.Body
}

func (s *AssociateProductWithPortfolioResponse) SetHeaders(v map[string]*string) *AssociateProductWithPortfolioResponse {
	s.Headers = v
	return s
}

func (s *AssociateProductWithPortfolioResponse) SetStatusCode(v int32) *AssociateProductWithPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateProductWithPortfolioResponse) SetBody(v *AssociateProductWithPortfolioResponseBody) *AssociateProductWithPortfolioResponse {
	s.Body = v
	return s
}

func (s *AssociateProductWithPortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
