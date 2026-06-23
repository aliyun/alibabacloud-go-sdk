// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsMaxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSBpsMaxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSBpsMaxResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSBpsMaxResponseBody) *DescribeDDoSBpsMaxResponse
	GetBody() *DescribeDDoSBpsMaxResponseBody
}

type DescribeDDoSBpsMaxResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSBpsMaxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSBpsMaxResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsMaxResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsMaxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSBpsMaxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSBpsMaxResponse) GetBody() *DescribeDDoSBpsMaxResponseBody {
	return s.Body
}

func (s *DescribeDDoSBpsMaxResponse) SetHeaders(v map[string]*string) *DescribeDDoSBpsMaxResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSBpsMaxResponse) SetStatusCode(v int32) *DescribeDDoSBpsMaxResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSBpsMaxResponse) SetBody(v *DescribeDDoSBpsMaxResponseBody) *DescribeDDoSBpsMaxResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSBpsMaxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
