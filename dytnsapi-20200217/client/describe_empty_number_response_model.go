// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmptyNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEmptyNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEmptyNumberResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEmptyNumberResponseBody) *DescribeEmptyNumberResponse
	GetBody() *DescribeEmptyNumberResponseBody
}

type DescribeEmptyNumberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEmptyNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEmptyNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmptyNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEmptyNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEmptyNumberResponse) GetBody() *DescribeEmptyNumberResponseBody {
	return s.Body
}

func (s *DescribeEmptyNumberResponse) SetHeaders(v map[string]*string) *DescribeEmptyNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmptyNumberResponse) SetStatusCode(v int32) *DescribeEmptyNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmptyNumberResponse) SetBody(v *DescribeEmptyNumberResponseBody) *DescribeEmptyNumberResponse {
	s.Body = v
	return s
}

func (s *DescribeEmptyNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
