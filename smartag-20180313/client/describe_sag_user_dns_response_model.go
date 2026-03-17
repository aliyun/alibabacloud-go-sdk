// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagUserDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagUserDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagUserDnsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagUserDnsResponseBody) *DescribeSagUserDnsResponse
	GetBody() *DescribeSagUserDnsResponseBody
}

type DescribeSagUserDnsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagUserDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagUserDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagUserDnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagUserDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagUserDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagUserDnsResponse) GetBody() *DescribeSagUserDnsResponseBody {
	return s.Body
}

func (s *DescribeSagUserDnsResponse) SetHeaders(v map[string]*string) *DescribeSagUserDnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagUserDnsResponse) SetStatusCode(v int32) *DescribeSagUserDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagUserDnsResponse) SetBody(v *DescribeSagUserDnsResponseBody) *DescribeSagUserDnsResponse {
	s.Body = v
	return s
}

func (s *DescribeSagUserDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
