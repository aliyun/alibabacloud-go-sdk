// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationServerlessConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationServerlessConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationServerlessConfResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationServerlessConfResponseBody) *DescribeApplicationServerlessConfResponse
	GetBody() *DescribeApplicationServerlessConfResponseBody
}

type DescribeApplicationServerlessConfResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationServerlessConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationServerlessConfResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationServerlessConfResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationServerlessConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationServerlessConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationServerlessConfResponse) GetBody() *DescribeApplicationServerlessConfResponseBody {
	return s.Body
}

func (s *DescribeApplicationServerlessConfResponse) SetHeaders(v map[string]*string) *DescribeApplicationServerlessConfResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationServerlessConfResponse) SetStatusCode(v int32) *DescribeApplicationServerlessConfResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationServerlessConfResponse) SetBody(v *DescribeApplicationServerlessConfResponseBody) *DescribeApplicationServerlessConfResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationServerlessConfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
