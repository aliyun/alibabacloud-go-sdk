// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDocTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDocTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDocTypesResponseBody) *DescribeDocTypesResponse
	GetBody() *DescribeDocTypesResponseBody
}

type DescribeDocTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDocTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDocTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDocTypesResponse) GetBody() *DescribeDocTypesResponseBody {
	return s.Body
}

func (s *DescribeDocTypesResponse) SetHeaders(v map[string]*string) *DescribeDocTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocTypesResponse) SetStatusCode(v int32) *DescribeDocTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocTypesResponse) SetBody(v *DescribeDocTypesResponseBody) *DescribeDocTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeDocTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
