// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureRiskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecureRiskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecureRiskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecureRiskListResponseBody) *DescribeDomainSecureRiskListResponse
	GetBody() *DescribeDomainSecureRiskListResponseBody
}

type DescribeDomainSecureRiskListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecureRiskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecureRiskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureRiskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureRiskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecureRiskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecureRiskListResponse) GetBody() *DescribeDomainSecureRiskListResponseBody {
	return s.Body
}

func (s *DescribeDomainSecureRiskListResponse) SetHeaders(v map[string]*string) *DescribeDomainSecureRiskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecureRiskListResponse) SetStatusCode(v int32) *DescribeDomainSecureRiskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponse) SetBody(v *DescribeDomainSecureRiskListResponseBody) *DescribeDomainSecureRiskListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecureRiskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
