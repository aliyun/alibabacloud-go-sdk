// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveUserTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveUserTagsResponseBody) *DescribeLiveUserTagsResponse
	GetBody() *DescribeLiveUserTagsResponseBody
}

type DescribeLiveUserTagsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveUserTagsResponse) GetBody() *DescribeLiveUserTagsResponseBody {
	return s.Body
}

func (s *DescribeLiveUserTagsResponse) SetHeaders(v map[string]*string) *DescribeLiveUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveUserTagsResponse) SetStatusCode(v int32) *DescribeLiveUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveUserTagsResponse) SetBody(v *DescribeLiveUserTagsResponseBody) *DescribeLiveUserTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveUserTagsResponse) Validate() error {
	return dara.Validate(s)
}
