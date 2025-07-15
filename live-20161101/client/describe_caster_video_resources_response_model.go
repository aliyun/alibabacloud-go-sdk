// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterVideoResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterVideoResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterVideoResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterVideoResourcesResponseBody) *DescribeCasterVideoResourcesResponse
	GetBody() *DescribeCasterVideoResourcesResponseBody
}

type DescribeCasterVideoResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterVideoResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterVideoResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterVideoResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterVideoResourcesResponse) GetBody() *DescribeCasterVideoResourcesResponseBody {
	return s.Body
}

func (s *DescribeCasterVideoResourcesResponse) SetHeaders(v map[string]*string) *DescribeCasterVideoResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterVideoResourcesResponse) SetStatusCode(v int32) *DescribeCasterVideoResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponse) SetBody(v *DescribeCasterVideoResourcesResponseBody) *DescribeCasterVideoResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterVideoResourcesResponse) Validate() error {
	return dara.Validate(s)
}
