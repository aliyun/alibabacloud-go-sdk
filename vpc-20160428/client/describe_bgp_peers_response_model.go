// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpPeersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBgpPeersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBgpPeersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBgpPeersResponseBody) *DescribeBgpPeersResponse
	GetBody() *DescribeBgpPeersResponseBody
}

type DescribeBgpPeersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBgpPeersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBgpPeersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPeersResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpPeersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBgpPeersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBgpPeersResponse) GetBody() *DescribeBgpPeersResponseBody {
	return s.Body
}

func (s *DescribeBgpPeersResponse) SetHeaders(v map[string]*string) *DescribeBgpPeersResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpPeersResponse) SetStatusCode(v int32) *DescribeBgpPeersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBgpPeersResponse) SetBody(v *DescribeBgpPeersResponseBody) *DescribeBgpPeersResponse {
	s.Body = v
	return s
}

func (s *DescribeBgpPeersResponse) Validate() error {
	return dara.Validate(s)
}
