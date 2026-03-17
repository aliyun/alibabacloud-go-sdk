// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQosesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQosesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQosesResponseBody) *DescribeQosesResponse
	GetBody() *DescribeQosesResponseBody
}

type DescribeQosesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQosesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQosesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosesResponse) GoString() string {
	return s.String()
}

func (s *DescribeQosesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQosesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQosesResponse) GetBody() *DescribeQosesResponseBody {
	return s.Body
}

func (s *DescribeQosesResponse) SetHeaders(v map[string]*string) *DescribeQosesResponse {
	s.Headers = v
	return s
}

func (s *DescribeQosesResponse) SetStatusCode(v int32) *DescribeQosesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQosesResponse) SetBody(v *DescribeQosesResponseBody) *DescribeQosesResponse {
	s.Body = v
	return s
}

func (s *DescribeQosesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
