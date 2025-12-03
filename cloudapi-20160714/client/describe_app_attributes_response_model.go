// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppAttributesResponseBody) *DescribeAppAttributesResponse
	GetBody() *DescribeAppAttributesResponseBody
}

type DescribeAppAttributesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppAttributesResponse) GetBody() *DescribeAppAttributesResponseBody {
	return s.Body
}

func (s *DescribeAppAttributesResponse) SetHeaders(v map[string]*string) *DescribeAppAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppAttributesResponse) SetStatusCode(v int32) *DescribeAppAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppAttributesResponse) SetBody(v *DescribeAppAttributesResponseBody) *DescribeAppAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
