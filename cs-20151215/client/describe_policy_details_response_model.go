// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyDetailsResponseBody) *DescribePolicyDetailsResponse
	GetBody() *DescribePolicyDetailsResponseBody
}

type DescribePolicyDetailsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyDetailsResponse) GetBody() *DescribePolicyDetailsResponseBody {
	return s.Body
}

func (s *DescribePolicyDetailsResponse) SetHeaders(v map[string]*string) *DescribePolicyDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyDetailsResponse) SetStatusCode(v int32) *DescribePolicyDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyDetailsResponse) SetBody(v *DescribePolicyDetailsResponseBody) *DescribePolicyDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
