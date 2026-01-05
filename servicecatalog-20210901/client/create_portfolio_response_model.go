// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePortfolioResponse
	GetStatusCode() *int32
	SetBody(v *CreatePortfolioResponseBody) *CreatePortfolioResponse
	GetBody() *CreatePortfolioResponseBody
}

type CreatePortfolioResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePortfolioResponse) GoString() string {
	return s.String()
}

func (s *CreatePortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePortfolioResponse) GetBody() *CreatePortfolioResponseBody {
	return s.Body
}

func (s *CreatePortfolioResponse) SetHeaders(v map[string]*string) *CreatePortfolioResponse {
	s.Headers = v
	return s
}

func (s *CreatePortfolioResponse) SetStatusCode(v int32) *CreatePortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortfolioResponse) SetBody(v *CreatePortfolioResponseBody) *CreatePortfolioResponse {
	s.Body = v
	return s
}

func (s *CreatePortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
