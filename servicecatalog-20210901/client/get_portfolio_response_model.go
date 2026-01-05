// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPortfolioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPortfolioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPortfolioResponse
	GetStatusCode() *int32
	SetBody(v *GetPortfolioResponseBody) *GetPortfolioResponse
	GetBody() *GetPortfolioResponseBody
}

type GetPortfolioResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPortfolioResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPortfolioResponse) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPortfolioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPortfolioResponse) GetBody() *GetPortfolioResponseBody {
	return s.Body
}

func (s *GetPortfolioResponse) SetHeaders(v map[string]*string) *GetPortfolioResponse {
	s.Headers = v
	return s
}

func (s *GetPortfolioResponse) SetStatusCode(v int32) *GetPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPortfolioResponse) SetBody(v *GetPortfolioResponseBody) *GetPortfolioResponse {
	s.Body = v
	return s
}

func (s *GetPortfolioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
