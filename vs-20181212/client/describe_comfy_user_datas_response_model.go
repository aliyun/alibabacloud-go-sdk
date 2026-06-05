// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDatasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyUserDatasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyUserDatasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyUserDatasResponseBody) *DescribeComfyUserDatasResponse
	GetBody() *DescribeComfyUserDatasResponseBody
}

type DescribeComfyUserDatasResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyUserDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyUserDatasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDatasResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDatasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyUserDatasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyUserDatasResponse) GetBody() *DescribeComfyUserDatasResponseBody {
	return s.Body
}

func (s *DescribeComfyUserDatasResponse) SetHeaders(v map[string]*string) *DescribeComfyUserDatasResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyUserDatasResponse) SetStatusCode(v int32) *DescribeComfyUserDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyUserDatasResponse) SetBody(v *DescribeComfyUserDatasResponseBody) *DescribeComfyUserDatasResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyUserDatasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
