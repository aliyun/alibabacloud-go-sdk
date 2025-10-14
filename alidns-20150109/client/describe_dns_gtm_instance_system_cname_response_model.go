// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceSystemCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceSystemCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmInstanceSystemCnameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmInstanceSystemCnameResponseBody) *DescribeDnsGtmInstanceSystemCnameResponse
	GetBody() *DescribeDnsGtmInstanceSystemCnameResponseBody
}

type DescribeDnsGtmInstanceSystemCnameResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmInstanceSystemCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmInstanceSystemCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) GetBody() *DescribeDnsGtmInstanceSystemCnameResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceSystemCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) SetStatusCode(v int32) *DescribeDnsGtmInstanceSystemCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) SetBody(v *DescribeDnsGtmInstanceSystemCnameResponseBody) *DescribeDnsGtmInstanceSystemCnameResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
