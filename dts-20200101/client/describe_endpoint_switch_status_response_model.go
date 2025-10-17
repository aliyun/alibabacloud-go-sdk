// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointSwitchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEndpointSwitchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEndpointSwitchStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEndpointSwitchStatusResponseBody) *DescribeEndpointSwitchStatusResponse
	GetBody() *DescribeEndpointSwitchStatusResponseBody
}

type DescribeEndpointSwitchStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndpointSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndpointSwitchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEndpointSwitchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEndpointSwitchStatusResponse) GetBody() *DescribeEndpointSwitchStatusResponseBody {
	return s.Body
}

func (s *DescribeEndpointSwitchStatusResponse) SetHeaders(v map[string]*string) *DescribeEndpointSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointSwitchStatusResponse) SetStatusCode(v int32) *DescribeEndpointSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponse) SetBody(v *DescribeEndpointSwitchStatusResponseBody) *DescribeEndpointSwitchStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeEndpointSwitchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
