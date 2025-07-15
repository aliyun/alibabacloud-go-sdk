// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterLayoutsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterLayoutsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterLayoutsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterLayoutsResponseBody) *DescribeCasterLayoutsResponse
	GetBody() *DescribeCasterLayoutsResponseBody
}

type DescribeCasterLayoutsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterLayoutsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterLayoutsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterLayoutsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterLayoutsResponse) GetBody() *DescribeCasterLayoutsResponseBody {
	return s.Body
}

func (s *DescribeCasterLayoutsResponse) SetHeaders(v map[string]*string) *DescribeCasterLayoutsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterLayoutsResponse) SetStatusCode(v int32) *DescribeCasterLayoutsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterLayoutsResponse) SetBody(v *DescribeCasterLayoutsResponseBody) *DescribeCasterLayoutsResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterLayoutsResponse) Validate() error {
	return dara.Validate(s)
}
