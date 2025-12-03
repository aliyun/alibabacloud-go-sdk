// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainExtensionsResponseBody) *DescribeDomainExtensionsResponse
	GetBody() *DescribeDomainExtensionsResponseBody
}

type DescribeDomainExtensionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainExtensionsResponse) GetBody() *DescribeDomainExtensionsResponseBody {
	return s.Body
}

func (s *DescribeDomainExtensionsResponse) SetHeaders(v map[string]*string) *DescribeDomainExtensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainExtensionsResponse) SetStatusCode(v int32) *DescribeDomainExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainExtensionsResponse) SetBody(v *DescribeDomainExtensionsResponseBody) *DescribeDomainExtensionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
