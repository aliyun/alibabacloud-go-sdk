// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnsecServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnsecServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnsecServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnsecServiceResponseBody) *DescribeDcdnsecServiceResponse
	GetBody() *DescribeDcdnsecServiceResponseBody
}

type DescribeDcdnsecServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnsecServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnsecServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnsecServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnsecServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnsecServiceResponse) GetBody() *DescribeDcdnsecServiceResponseBody {
	return s.Body
}

func (s *DescribeDcdnsecServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnsecServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnsecServiceResponse) SetStatusCode(v int32) *DescribeDcdnsecServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnsecServiceResponse) SetBody(v *DescribeDcdnsecServiceResponseBody) *DescribeDcdnsecServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnsecServiceResponse) Validate() error {
	return dara.Validate(s)
}
