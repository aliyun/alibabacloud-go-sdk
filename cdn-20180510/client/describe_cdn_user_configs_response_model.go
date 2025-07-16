// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserConfigsResponseBody) *DescribeCdnUserConfigsResponse
	GetBody() *DescribeCdnUserConfigsResponseBody
}

type DescribeCdnUserConfigsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserConfigsResponse) GetBody() *DescribeCdnUserConfigsResponseBody {
	return s.Body
}

func (s *DescribeCdnUserConfigsResponse) SetHeaders(v map[string]*string) *DescribeCdnUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserConfigsResponse) SetStatusCode(v int32) *DescribeCdnUserConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserConfigsResponse) SetBody(v *DescribeCdnUserConfigsResponseBody) *DescribeCdnUserConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserConfigsResponse) Validate() error {
	return dara.Validate(s)
}
