// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortfoliosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPortfoliosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPortfoliosResponse
	GetStatusCode() *int32
	SetBody(v *ListPortfoliosResponseBody) *ListPortfoliosResponse
	GetBody() *ListPortfoliosResponseBody
}

type ListPortfoliosResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortfoliosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPortfoliosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPortfoliosResponse) GoString() string {
	return s.String()
}

func (s *ListPortfoliosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPortfoliosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPortfoliosResponse) GetBody() *ListPortfoliosResponseBody {
	return s.Body
}

func (s *ListPortfoliosResponse) SetHeaders(v map[string]*string) *ListPortfoliosResponse {
	s.Headers = v
	return s
}

func (s *ListPortfoliosResponse) SetStatusCode(v int32) *ListPortfoliosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortfoliosResponse) SetBody(v *ListPortfoliosResponseBody) *ListPortfoliosResponse {
	s.Body = v
	return s
}

func (s *ListPortfoliosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
