// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighlightInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHighlightInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHighlightInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHighlightInfoResponseBody) *DescribeHighlightInfoResponse
	GetBody() *DescribeHighlightInfoResponseBody
}

type DescribeHighlightInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHighlightInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHighlightInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighlightInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeHighlightInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHighlightInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHighlightInfoResponse) GetBody() *DescribeHighlightInfoResponseBody {
	return s.Body
}

func (s *DescribeHighlightInfoResponse) SetHeaders(v map[string]*string) *DescribeHighlightInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeHighlightInfoResponse) SetStatusCode(v int32) *DescribeHighlightInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHighlightInfoResponse) SetBody(v *DescribeHighlightInfoResponseBody) *DescribeHighlightInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeHighlightInfoResponse) Validate() error {
	return dara.Validate(s)
}
