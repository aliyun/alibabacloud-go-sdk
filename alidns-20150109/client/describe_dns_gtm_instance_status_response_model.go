// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmInstanceStatusResponseBody) *DescribeDnsGtmInstanceStatusResponse
	GetBody() *DescribeDnsGtmInstanceStatusResponseBody
}

type DescribeDnsGtmInstanceStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmInstanceStatusResponse) GetBody() *DescribeDnsGtmInstanceStatusResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponse) SetStatusCode(v int32) *DescribeDnsGtmInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponse) SetBody(v *DescribeDnsGtmInstanceStatusResponseBody) *DescribeDnsGtmInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponse) Validate() error {
	return dara.Validate(s)
}
