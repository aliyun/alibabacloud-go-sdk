// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnServiceResponseBody) *DescribeDcdnServiceResponse
	GetBody() *DescribeDcdnServiceResponseBody
}

type DescribeDcdnServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnServiceResponse) GetBody() *DescribeDcdnServiceResponseBody {
	return s.Body
}

func (s *DescribeDcdnServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnServiceResponse) SetStatusCode(v int32) *DescribeDcdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnServiceResponse) SetBody(v *DescribeDcdnServiceResponseBody) *DescribeDcdnServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnServiceResponse) Validate() error {
	return dara.Validate(s)
}
