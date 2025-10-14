// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddrAttributeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmAddrAttributeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmAddrAttributeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmAddrAttributeInfoResponseBody) *DescribeDnsGtmAddrAttributeInfoResponse
	GetBody() *DescribeDnsGtmAddrAttributeInfoResponseBody
}

type DescribeDnsGtmAddrAttributeInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmAddrAttributeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) GetBody() *DescribeDnsGtmAddrAttributeInfoResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAddrAttributeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) SetStatusCode(v int32) *DescribeDnsGtmAddrAttributeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) SetBody(v *DescribeDnsGtmAddrAttributeInfoResponseBody) *DescribeDnsGtmAddrAttributeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
