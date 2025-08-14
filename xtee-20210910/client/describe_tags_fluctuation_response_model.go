// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsFluctuationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsFluctuationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsFluctuationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsFluctuationResponseBody) *DescribeTagsFluctuationResponse
	GetBody() *DescribeTagsFluctuationResponseBody
}

type DescribeTagsFluctuationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsFluctuationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsFluctuationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsFluctuationResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsFluctuationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsFluctuationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsFluctuationResponse) GetBody() *DescribeTagsFluctuationResponseBody {
	return s.Body
}

func (s *DescribeTagsFluctuationResponse) SetHeaders(v map[string]*string) *DescribeTagsFluctuationResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsFluctuationResponse) SetStatusCode(v int32) *DescribeTagsFluctuationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsFluctuationResponse) SetBody(v *DescribeTagsFluctuationResponseBody) *DescribeTagsFluctuationResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsFluctuationResponse) Validate() error {
	return dara.Validate(s)
}
