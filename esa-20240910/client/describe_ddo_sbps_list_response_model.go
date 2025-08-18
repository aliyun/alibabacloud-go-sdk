// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSBpsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSBpsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSBpsListResponseBody) *DescribeDDoSBpsListResponse
	GetBody() *DescribeDDoSBpsListResponseBody
}

type DescribeDDoSBpsListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSBpsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSBpsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSBpsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSBpsListResponse) GetBody() *DescribeDDoSBpsListResponseBody {
	return s.Body
}

func (s *DescribeDDoSBpsListResponse) SetHeaders(v map[string]*string) *DescribeDDoSBpsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSBpsListResponse) SetStatusCode(v int32) *DescribeDDoSBpsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSBpsListResponse) SetBody(v *DescribeDDoSBpsListResponseBody) *DescribeDDoSBpsListResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSBpsListResponse) Validate() error {
	return dara.Validate(s)
}
