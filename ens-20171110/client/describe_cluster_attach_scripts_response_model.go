// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttachScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAttachScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAttachScriptsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterAttachScriptsResponseBody) *DescribeClusterAttachScriptsResponse
	GetBody() *DescribeClusterAttachScriptsResponseBody
}

type DescribeClusterAttachScriptsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAttachScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAttachScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttachScriptsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAttachScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAttachScriptsResponse) GetBody() *DescribeClusterAttachScriptsResponseBody {
	return s.Body
}

func (s *DescribeClusterAttachScriptsResponse) SetHeaders(v map[string]*string) *DescribeClusterAttachScriptsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAttachScriptsResponse) SetStatusCode(v int32) *DescribeClusterAttachScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAttachScriptsResponse) SetBody(v *DescribeClusterAttachScriptsResponseBody) *DescribeClusterAttachScriptsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAttachScriptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
