// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePresetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePresetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePresetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePresetsResponseBody) *DescribePresetsResponse
	GetBody() *DescribePresetsResponseBody
}

type DescribePresetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePresetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePresetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePresetsResponse) GoString() string {
	return s.String()
}

func (s *DescribePresetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePresetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePresetsResponse) GetBody() *DescribePresetsResponseBody {
	return s.Body
}

func (s *DescribePresetsResponse) SetHeaders(v map[string]*string) *DescribePresetsResponse {
	s.Headers = v
	return s
}

func (s *DescribePresetsResponse) SetStatusCode(v int32) *DescribePresetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePresetsResponse) SetBody(v *DescribePresetsResponseBody) *DescribePresetsResponse {
	s.Body = v
	return s
}

func (s *DescribePresetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
