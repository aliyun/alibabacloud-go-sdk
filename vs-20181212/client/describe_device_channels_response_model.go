// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceChannelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceChannelsResponseBody) *DescribeDeviceChannelsResponse
	GetBody() *DescribeDeviceChannelsResponseBody
}

type DescribeDeviceChannelsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceChannelsResponse) GetBody() *DescribeDeviceChannelsResponseBody {
	return s.Body
}

func (s *DescribeDeviceChannelsResponse) SetHeaders(v map[string]*string) *DescribeDeviceChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceChannelsResponse) SetStatusCode(v int32) *DescribeDeviceChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceChannelsResponse) SetBody(v *DescribeDeviceChannelsResponseBody) *DescribeDeviceChannelsResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
