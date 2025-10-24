// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMirrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMirrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMirrorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMirrorResponseBody) *DescribeServiceMirrorResponse
	GetBody() *DescribeServiceMirrorResponseBody
}

type DescribeServiceMirrorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMirrorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMirrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMirrorResponse) GetBody() *DescribeServiceMirrorResponseBody {
	return s.Body
}

func (s *DescribeServiceMirrorResponse) SetHeaders(v map[string]*string) *DescribeServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMirrorResponse) SetStatusCode(v int32) *DescribeServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMirrorResponse) SetBody(v *DescribeServiceMirrorResponseBody) *DescribeServiceMirrorResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMirrorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
