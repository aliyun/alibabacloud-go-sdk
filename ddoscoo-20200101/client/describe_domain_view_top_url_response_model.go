// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainViewTopUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainViewTopUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainViewTopUrlResponseBody) *DescribeDomainViewTopUrlResponse
	GetBody() *DescribeDomainViewTopUrlResponseBody
}

type DescribeDomainViewTopUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainViewTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainViewTopUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainViewTopUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainViewTopUrlResponse) GetBody() *DescribeDomainViewTopUrlResponseBody {
	return s.Body
}

func (s *DescribeDomainViewTopUrlResponse) SetHeaders(v map[string]*string) *DescribeDomainViewTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewTopUrlResponse) SetStatusCode(v int32) *DescribeDomainViewTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponse) SetBody(v *DescribeDomainViewTopUrlResponseBody) *DescribeDomainViewTopUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainViewTopUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
