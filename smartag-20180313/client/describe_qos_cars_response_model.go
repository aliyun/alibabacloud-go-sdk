// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosCarsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQosCarsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQosCarsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQosCarsResponseBody) *DescribeQosCarsResponse
	GetBody() *DescribeQosCarsResponseBody
}

type DescribeQosCarsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQosCarsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQosCarsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosCarsResponse) GoString() string {
	return s.String()
}

func (s *DescribeQosCarsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQosCarsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQosCarsResponse) GetBody() *DescribeQosCarsResponseBody {
	return s.Body
}

func (s *DescribeQosCarsResponse) SetHeaders(v map[string]*string) *DescribeQosCarsResponse {
	s.Headers = v
	return s
}

func (s *DescribeQosCarsResponse) SetStatusCode(v int32) *DescribeQosCarsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQosCarsResponse) SetBody(v *DescribeQosCarsResponseBody) *DescribeQosCarsResponse {
	s.Body = v
	return s
}

func (s *DescribeQosCarsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
