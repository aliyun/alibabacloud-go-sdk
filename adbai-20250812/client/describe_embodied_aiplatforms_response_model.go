// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmbodiedAIPlatformsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEmbodiedAIPlatformsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEmbodiedAIPlatformsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEmbodiedAIPlatformsResponseBody) *DescribeEmbodiedAIPlatformsResponse
	GetBody() *DescribeEmbodiedAIPlatformsResponseBody
}

type DescribeEmbodiedAIPlatformsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEmbodiedAIPlatformsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEmbodiedAIPlatformsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEmbodiedAIPlatformsResponse) GetBody() *DescribeEmbodiedAIPlatformsResponseBody {
	return s.Body
}

func (s *DescribeEmbodiedAIPlatformsResponse) SetHeaders(v map[string]*string) *DescribeEmbodiedAIPlatformsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponse) SetStatusCode(v int32) *DescribeEmbodiedAIPlatformsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponse) SetBody(v *DescribeEmbodiedAIPlatformsResponseBody) *DescribeEmbodiedAIPlatformsResponse {
	s.Body = v
	return s
}

func (s *DescribeEmbodiedAIPlatformsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
