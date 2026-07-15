// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSessionIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationSessionIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationSessionIdsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationSessionIdsResponseBody) *DescribeApplicationSessionIdsResponse
	GetBody() *DescribeApplicationSessionIdsResponseBody
}

type DescribeApplicationSessionIdsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationSessionIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationSessionIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionIdsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationSessionIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationSessionIdsResponse) GetBody() *DescribeApplicationSessionIdsResponseBody {
	return s.Body
}

func (s *DescribeApplicationSessionIdsResponse) SetHeaders(v map[string]*string) *DescribeApplicationSessionIdsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationSessionIdsResponse) SetStatusCode(v int32) *DescribeApplicationSessionIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponse) SetBody(v *DescribeApplicationSessionIdsResponseBody) *DescribeApplicationSessionIdsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationSessionIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
