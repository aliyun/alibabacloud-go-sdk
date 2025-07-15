// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpNetworksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBgpNetworksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBgpNetworksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBgpNetworksResponseBody) *DescribeBgpNetworksResponse
	GetBody() *DescribeBgpNetworksResponseBody
}

type DescribeBgpNetworksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBgpNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBgpNetworksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpNetworksResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpNetworksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBgpNetworksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBgpNetworksResponse) GetBody() *DescribeBgpNetworksResponseBody {
	return s.Body
}

func (s *DescribeBgpNetworksResponse) SetHeaders(v map[string]*string) *DescribeBgpNetworksResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpNetworksResponse) SetStatusCode(v int32) *DescribeBgpNetworksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBgpNetworksResponse) SetBody(v *DescribeBgpNetworksResponseBody) *DescribeBgpNetworksResponse {
	s.Body = v
	return s
}

func (s *DescribeBgpNetworksResponse) Validate() error {
	return dara.Validate(s)
}
