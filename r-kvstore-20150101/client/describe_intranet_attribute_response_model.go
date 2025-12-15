// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntranetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIntranetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIntranetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIntranetAttributeResponseBody) *DescribeIntranetAttributeResponse
	GetBody() *DescribeIntranetAttributeResponseBody
}

type DescribeIntranetAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntranetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntranetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntranetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIntranetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIntranetAttributeResponse) GetBody() *DescribeIntranetAttributeResponseBody {
	return s.Body
}

func (s *DescribeIntranetAttributeResponse) SetHeaders(v map[string]*string) *DescribeIntranetAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntranetAttributeResponse) SetStatusCode(v int32) *DescribeIntranetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntranetAttributeResponse) SetBody(v *DescribeIntranetAttributeResponseBody) *DescribeIntranetAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeIntranetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
