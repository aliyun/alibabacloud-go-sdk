// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvadeEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvadeEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvadeEventDetailResponseBody) *DescribeInvadeEventDetailResponse
	GetBody() *DescribeInvadeEventDetailResponseBody
}

type DescribeInvadeEventDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvadeEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvadeEventDetailResponse) GetBody() *DescribeInvadeEventDetailResponseBody {
	return s.Body
}

func (s *DescribeInvadeEventDetailResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventDetailResponse) SetStatusCode(v int32) *DescribeInvadeEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventDetailResponse) SetBody(v *DescribeInvadeEventDetailResponseBody) *DescribeInvadeEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeInvadeEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
