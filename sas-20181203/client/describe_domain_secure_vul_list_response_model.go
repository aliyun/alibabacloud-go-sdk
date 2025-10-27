// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureVulListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecureVulListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecureVulListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecureVulListResponseBody) *DescribeDomainSecureVulListResponse
	GetBody() *DescribeDomainSecureVulListResponseBody
}

type DescribeDomainSecureVulListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecureVulListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecureVulListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureVulListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecureVulListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecureVulListResponse) GetBody() *DescribeDomainSecureVulListResponseBody {
	return s.Body
}

func (s *DescribeDomainSecureVulListResponse) SetHeaders(v map[string]*string) *DescribeDomainSecureVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecureVulListResponse) SetStatusCode(v int32) *DescribeDomainSecureVulListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecureVulListResponse) SetBody(v *DescribeDomainSecureVulListResponseBody) *DescribeDomainSecureVulListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecureVulListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
