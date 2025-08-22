// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineCodeRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoutineCodeRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoutineCodeRevisionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoutineCodeRevisionResponseBody) *DescribeRoutineCodeRevisionResponse
	GetBody() *DescribeRoutineCodeRevisionResponseBody
}

type DescribeRoutineCodeRevisionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoutineCodeRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoutineCodeRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineCodeRevisionResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCodeRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoutineCodeRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoutineCodeRevisionResponse) GetBody() *DescribeRoutineCodeRevisionResponseBody {
	return s.Body
}

func (s *DescribeRoutineCodeRevisionResponse) SetHeaders(v map[string]*string) *DescribeRoutineCodeRevisionResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineCodeRevisionResponse) SetStatusCode(v int32) *DescribeRoutineCodeRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoutineCodeRevisionResponse) SetBody(v *DescribeRoutineCodeRevisionResponseBody) *DescribeRoutineCodeRevisionResponse {
	s.Body = v
	return s
}

func (s *DescribeRoutineCodeRevisionResponse) Validate() error {
	return dara.Validate(s)
}
