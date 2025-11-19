// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodVerifyContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodVerifyContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodVerifyContentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodVerifyContentResponseBody) *DescribeVodVerifyContentResponse
	GetBody() *DescribeVodVerifyContentResponseBody
}

type DescribeVodVerifyContentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodVerifyContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodVerifyContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodVerifyContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodVerifyContentResponse) GetBody() *DescribeVodVerifyContentResponseBody {
	return s.Body
}

func (s *DescribeVodVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeVodVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodVerifyContentResponse) SetStatusCode(v int32) *DescribeVodVerifyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodVerifyContentResponse) SetBody(v *DescribeVodVerifyContentResponseBody) *DescribeVodVerifyContentResponse {
	s.Body = v
	return s
}

func (s *DescribeVodVerifyContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
