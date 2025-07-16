// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSrcQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSrcQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSrcQpsDataResponseBody) *DescribeDomainSrcQpsDataResponse
	GetBody() *DescribeDomainSrcQpsDataResponseBody
}

type DescribeDomainSrcQpsDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSrcQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSrcQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSrcQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSrcQpsDataResponse) GetBody() *DescribeDomainSrcQpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainSrcQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcQpsDataResponse) SetStatusCode(v int32) *DescribeDomainSrcQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponse) SetBody(v *DescribeDomainSrcQpsDataResponseBody) *DescribeDomainSrcQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSrcQpsDataResponse) Validate() error {
	return dara.Validate(s)
}
