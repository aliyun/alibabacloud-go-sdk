// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseOrderDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLicenseOrderDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLicenseOrderDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLicenseOrderDetailsResponseBody) *DescribeLicenseOrderDetailsResponse
	GetBody() *DescribeLicenseOrderDetailsResponseBody
}

type DescribeLicenseOrderDetailsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLicenseOrderDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLicenseOrderDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrderDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrderDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLicenseOrderDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLicenseOrderDetailsResponse) GetBody() *DescribeLicenseOrderDetailsResponseBody {
	return s.Body
}

func (s *DescribeLicenseOrderDetailsResponse) SetHeaders(v map[string]*string) *DescribeLicenseOrderDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLicenseOrderDetailsResponse) SetStatusCode(v int32) *DescribeLicenseOrderDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponse) SetBody(v *DescribeLicenseOrderDetailsResponseBody) *DescribeLicenseOrderDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeLicenseOrderDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
