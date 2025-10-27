// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFixUsedCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFixUsedCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFixUsedCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFixUsedCountResponseBody) *DescribeFixUsedCountResponse
	GetBody() *DescribeFixUsedCountResponseBody
}

type DescribeFixUsedCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFixUsedCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFixUsedCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFixUsedCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeFixUsedCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFixUsedCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFixUsedCountResponse) GetBody() *DescribeFixUsedCountResponseBody {
	return s.Body
}

func (s *DescribeFixUsedCountResponse) SetHeaders(v map[string]*string) *DescribeFixUsedCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeFixUsedCountResponse) SetStatusCode(v int32) *DescribeFixUsedCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFixUsedCountResponse) SetBody(v *DescribeFixUsedCountResponseBody) *DescribeFixUsedCountResponse {
	s.Body = v
	return s
}

func (s *DescribeFixUsedCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
