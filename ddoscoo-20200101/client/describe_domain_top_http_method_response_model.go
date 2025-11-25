// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopHttpMethodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopHttpMethodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopHttpMethodResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopHttpMethodResponseBody) *DescribeDomainTopHttpMethodResponse
	GetBody() *DescribeDomainTopHttpMethodResponseBody
}

type DescribeDomainTopHttpMethodResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopHttpMethodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopHttpMethodResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopHttpMethodResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopHttpMethodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopHttpMethodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopHttpMethodResponse) GetBody() *DescribeDomainTopHttpMethodResponseBody {
	return s.Body
}

func (s *DescribeDomainTopHttpMethodResponse) SetHeaders(v map[string]*string) *DescribeDomainTopHttpMethodResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopHttpMethodResponse) SetStatusCode(v int32) *DescribeDomainTopHttpMethodResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopHttpMethodResponse) SetBody(v *DescribeDomainTopHttpMethodResponseBody) *DescribeDomainTopHttpMethodResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopHttpMethodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
