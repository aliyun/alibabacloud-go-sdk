// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmInstanceResponseBody) *DescribeDnsGtmInstanceResponse
	GetBody() *DescribeDnsGtmInstanceResponseBody
}

type DescribeDnsGtmInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmInstanceResponse) GetBody() *DescribeDnsGtmInstanceResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmInstanceResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceResponse) SetStatusCode(v int32) *DescribeDnsGtmInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponse) SetBody(v *DescribeDnsGtmInstanceResponseBody) *DescribeDnsGtmInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
