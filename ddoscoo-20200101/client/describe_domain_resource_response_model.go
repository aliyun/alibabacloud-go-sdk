// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainResourceResponseBody) *DescribeDomainResourceResponse
	GetBody() *DescribeDomainResourceResponseBody
}

type DescribeDomainResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainResourceResponse) GetBody() *DescribeDomainResourceResponseBody {
	return s.Body
}

func (s *DescribeDomainResourceResponse) SetHeaders(v map[string]*string) *DescribeDomainResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResourceResponse) SetStatusCode(v int32) *DescribeDomainResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResourceResponse) SetBody(v *DescribeDomainResourceResponseBody) *DescribeDomainResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
