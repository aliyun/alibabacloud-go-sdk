// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagCurrentDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagCurrentDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagCurrentDnsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagCurrentDnsResponseBody) *DescribeSagCurrentDnsResponse
	GetBody() *DescribeSagCurrentDnsResponseBody
}

type DescribeSagCurrentDnsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagCurrentDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagCurrentDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagCurrentDnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagCurrentDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagCurrentDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagCurrentDnsResponse) GetBody() *DescribeSagCurrentDnsResponseBody {
	return s.Body
}

func (s *DescribeSagCurrentDnsResponse) SetHeaders(v map[string]*string) *DescribeSagCurrentDnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagCurrentDnsResponse) SetStatusCode(v int32) *DescribeSagCurrentDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagCurrentDnsResponse) SetBody(v *DescribeSagCurrentDnsResponseBody) *DescribeSagCurrentDnsResponse {
	s.Body = v
	return s
}

func (s *DescribeSagCurrentDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
