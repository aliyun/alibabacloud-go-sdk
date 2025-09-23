// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainQpsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainQpsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainQpsListResponseBody) *DescribeDomainQpsListResponse
	GetBody() *DescribeDomainQpsListResponseBody
}

type DescribeDomainQpsListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQpsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainQpsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainQpsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainQpsListResponse) GetBody() *DescribeDomainQpsListResponseBody {
	return s.Body
}

func (s *DescribeDomainQpsListResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsListResponse) SetStatusCode(v int32) *DescribeDomainQpsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsListResponse) SetBody(v *DescribeDomainQpsListResponseBody) *DescribeDomainQpsListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainQpsListResponse) Validate() error {
	return dara.Validate(s)
}
