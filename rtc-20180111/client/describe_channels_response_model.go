// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelsResponseBody) *DescribeChannelsResponse
	GetBody() *DescribeChannelsResponseBody
}

type DescribeChannelsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelsResponse) GetBody() *DescribeChannelsResponseBody {
	return s.Body
}

func (s *DescribeChannelsResponse) SetHeaders(v map[string]*string) *DescribeChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelsResponse) SetStatusCode(v int32) *DescribeChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelsResponse) SetBody(v *DescribeChannelsResponseBody) *DescribeChannelsResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
