// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDdosServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDdosServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDdosServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDdosServiceResponseBody) *DescribeDcdnDdosServiceResponse
	GetBody() *DescribeDcdnDdosServiceResponseBody
}

type DescribeDcdnDdosServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDdosServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDdosServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDdosServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDdosServiceResponse) GetBody() *DescribeDcdnDdosServiceResponseBody {
	return s.Body
}

func (s *DescribeDcdnDdosServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnDdosServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDdosServiceResponse) SetStatusCode(v int32) *DescribeDcdnDdosServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDdosServiceResponse) SetBody(v *DescribeDcdnDdosServiceResponseBody) *DescribeDcdnDdosServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDdosServiceResponse) Validate() error {
	return dara.Validate(s)
}
