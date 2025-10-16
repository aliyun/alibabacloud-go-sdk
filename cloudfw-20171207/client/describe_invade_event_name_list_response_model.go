// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvadeEventNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvadeEventNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvadeEventNameListResponseBody) *DescribeInvadeEventNameListResponse
	GetBody() *DescribeInvadeEventNameListResponseBody
}

type DescribeInvadeEventNameListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEventNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEventNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvadeEventNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvadeEventNameListResponse) GetBody() *DescribeInvadeEventNameListResponseBody {
	return s.Body
}

func (s *DescribeInvadeEventNameListResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventNameListResponse) SetStatusCode(v int32) *DescribeInvadeEventNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventNameListResponse) SetBody(v *DescribeInvadeEventNameListResponseBody) *DescribeInvadeEventNameListResponse {
	s.Body = v
	return s
}

func (s *DescribeInvadeEventNameListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
