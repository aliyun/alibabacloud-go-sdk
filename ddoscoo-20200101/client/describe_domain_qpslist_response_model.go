// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQPSListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainQPSListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainQPSListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainQPSListResponseBody) *DescribeDomainQPSListResponse
	GetBody() *DescribeDomainQPSListResponseBody
}

type DescribeDomainQPSListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQPSListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainQPSListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQPSListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainQPSListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainQPSListResponse) GetBody() *DescribeDomainQPSListResponseBody {
	return s.Body
}

func (s *DescribeDomainQPSListResponse) SetHeaders(v map[string]*string) *DescribeDomainQPSListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQPSListResponse) SetStatusCode(v int32) *DescribeDomainQPSListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQPSListResponse) SetBody(v *DescribeDomainQPSListResponseBody) *DescribeDomainQPSListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainQPSListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
