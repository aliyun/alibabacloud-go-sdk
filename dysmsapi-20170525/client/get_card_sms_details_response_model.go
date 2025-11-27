// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCardSmsDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCardSmsDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetCardSmsDetailsResponseBody) *GetCardSmsDetailsResponse
	GetBody() *GetCardSmsDetailsResponseBody
}

type GetCardSmsDetailsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardSmsDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardSmsDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCardSmsDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCardSmsDetailsResponse) GetBody() *GetCardSmsDetailsResponseBody {
	return s.Body
}

func (s *GetCardSmsDetailsResponse) SetHeaders(v map[string]*string) *GetCardSmsDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetCardSmsDetailsResponse) SetStatusCode(v int32) *GetCardSmsDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardSmsDetailsResponse) SetBody(v *GetCardSmsDetailsResponseBody) *GetCardSmsDetailsResponse {
	s.Body = v
	return s
}

func (s *GetCardSmsDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
