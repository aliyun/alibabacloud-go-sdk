// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDnssecInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainDnssecInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainDnssecInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainDnssecInfoResponseBody) *DescribeDomainDnssecInfoResponse
	GetBody() *DescribeDomainDnssecInfoResponseBody
}

type DescribeDomainDnssecInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainDnssecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainDnssecInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDnssecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDnssecInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainDnssecInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainDnssecInfoResponse) GetBody() *DescribeDomainDnssecInfoResponseBody {
	return s.Body
}

func (s *DescribeDomainDnssecInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainDnssecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDnssecInfoResponse) SetStatusCode(v int32) *DescribeDomainDnssecInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDnssecInfoResponse) SetBody(v *DescribeDomainDnssecInfoResponseBody) *DescribeDomainDnssecInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainDnssecInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
