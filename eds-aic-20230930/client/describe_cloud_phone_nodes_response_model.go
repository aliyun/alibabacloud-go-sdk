// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudPhoneNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudPhoneNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudPhoneNodesResponseBody) *DescribeCloudPhoneNodesResponse
	GetBody() *DescribeCloudPhoneNodesResponseBody
}

type DescribeCloudPhoneNodesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudPhoneNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudPhoneNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudPhoneNodesResponse) GetBody() *DescribeCloudPhoneNodesResponseBody {
	return s.Body
}

func (s *DescribeCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *DescribeCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudPhoneNodesResponse) SetStatusCode(v int32) *DescribeCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponse) SetBody(v *DescribeCloudPhoneNodesResponseBody) *DescribeCloudPhoneNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudPhoneNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
