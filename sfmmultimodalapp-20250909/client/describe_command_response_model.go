// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommandResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommandResponseBody) *DescribeCommandResponse
	GetBody() *DescribeCommandResponseBody
}

type DescribeCommandResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommandResponse) GetBody() *DescribeCommandResponseBody {
	return s.Body
}

func (s *DescribeCommandResponse) SetHeaders(v map[string]*string) *DescribeCommandResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommandResponse) SetStatusCode(v int32) *DescribeCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommandResponse) SetBody(v *DescribeCommandResponseBody) *DescribeCommandResponse {
	s.Body = v
	return s
}

func (s *DescribeCommandResponse) Validate() error {
	return dara.Validate(s)
}
