// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePortfolioResponse
	GetStatusCode() *int32
	SetBody(v *DeletePortfolioResponseBody) *DeletePortfolioResponse
	GetBody() *DeletePortfolioResponseBody
}

type DeletePortfolioResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePortfolioResponse) GoString() string {
	return s.String()
}

func (s *DeletePortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePortfolioResponse) GetBody() *DeletePortfolioResponseBody {
	return s.Body
}

func (s *DeletePortfolioResponse) SetHeaders(v map[string]*string) *DeletePortfolioResponse {
	s.Headers = v
	return s
}

func (s *DeletePortfolioResponse) SetStatusCode(v int32) *DeletePortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePortfolioResponse) SetBody(v *DeletePortfolioResponseBody) *DeletePortfolioResponse {
	s.Body = v
	return s
}

func (s *DeletePortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
