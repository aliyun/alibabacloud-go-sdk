// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkChannelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkChannelResponseBody) *DescribeNetworkChannelResponse
	GetBody() *DescribeNetworkChannelResponseBody
}

type DescribeNetworkChannelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkChannelResponse) GetBody() *DescribeNetworkChannelResponseBody {
	return s.Body
}

func (s *DescribeNetworkChannelResponse) SetHeaders(v map[string]*string) *DescribeNetworkChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkChannelResponse) SetStatusCode(v int32) *DescribeNetworkChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkChannelResponse) SetBody(v *DescribeNetworkChannelResponseBody) *DescribeNetworkChannelResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
