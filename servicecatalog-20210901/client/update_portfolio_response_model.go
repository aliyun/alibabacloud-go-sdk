// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePortfolioResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePortfolioResponseBody) *UpdatePortfolioResponse
	GetBody() *UpdatePortfolioResponseBody
}

type UpdatePortfolioResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePortfolioResponse) GoString() string {
	return s.String()
}

func (s *UpdatePortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePortfolioResponse) GetBody() *UpdatePortfolioResponseBody {
	return s.Body
}

func (s *UpdatePortfolioResponse) SetHeaders(v map[string]*string) *UpdatePortfolioResponse {
	s.Headers = v
	return s
}

func (s *UpdatePortfolioResponse) SetStatusCode(v int32) *UpdatePortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePortfolioResponse) SetBody(v *UpdatePortfolioResponseBody) *UpdatePortfolioResponse {
	s.Body = v
	return s
}

func (s *UpdatePortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
