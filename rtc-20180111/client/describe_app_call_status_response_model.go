// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppCallStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppCallStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppCallStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppCallStatusResponseBody) *DescribeAppCallStatusResponse
	GetBody() *DescribeAppCallStatusResponseBody
}

type DescribeAppCallStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppCallStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppCallStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppCallStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppCallStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppCallStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppCallStatusResponse) GetBody() *DescribeAppCallStatusResponseBody {
	return s.Body
}

func (s *DescribeAppCallStatusResponse) SetHeaders(v map[string]*string) *DescribeAppCallStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppCallStatusResponse) SetStatusCode(v int32) *DescribeAppCallStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppCallStatusResponse) SetBody(v *DescribeAppCallStatusResponseBody) *DescribeAppCallStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAppCallStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
