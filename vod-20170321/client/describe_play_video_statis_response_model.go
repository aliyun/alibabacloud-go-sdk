// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayVideoStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayVideoStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayVideoStatisResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayVideoStatisResponseBody) *DescribePlayVideoStatisResponse
	GetBody() *DescribePlayVideoStatisResponseBody
}

type DescribePlayVideoStatisResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayVideoStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayVideoStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayVideoStatisResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayVideoStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayVideoStatisResponse) GetBody() *DescribePlayVideoStatisResponseBody {
	return s.Body
}

func (s *DescribePlayVideoStatisResponse) SetHeaders(v map[string]*string) *DescribePlayVideoStatisResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayVideoStatisResponse) SetStatusCode(v int32) *DescribePlayVideoStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayVideoStatisResponse) SetBody(v *DescribePlayVideoStatisResponseBody) *DescribePlayVideoStatisResponse {
	s.Body = v
	return s
}

func (s *DescribePlayVideoStatisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
