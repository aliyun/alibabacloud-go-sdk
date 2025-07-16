// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainQpsDataResponseBody) *DescribeDomainQpsDataResponse
	GetBody() *DescribeDomainQpsDataResponseBody
}

type DescribeDomainQpsDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainQpsDataResponse) GetBody() *DescribeDomainQpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsDataResponse) SetStatusCode(v int32) *DescribeDomainQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsDataResponse) SetBody(v *DescribeDomainQpsDataResponseBody) *DescribeDomainQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainQpsDataResponse) Validate() error {
	return dara.Validate(s)
}
