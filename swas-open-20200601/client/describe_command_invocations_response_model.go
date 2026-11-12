// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandInvocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommandInvocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommandInvocationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommandInvocationsResponseBody) *DescribeCommandInvocationsResponse
	GetBody() *DescribeCommandInvocationsResponseBody
}

type DescribeCommandInvocationsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommandInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommandInvocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommandInvocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommandInvocationsResponse) GetBody() *DescribeCommandInvocationsResponseBody {
	return s.Body
}

func (s *DescribeCommandInvocationsResponse) SetHeaders(v map[string]*string) *DescribeCommandInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommandInvocationsResponse) SetStatusCode(v int32) *DescribeCommandInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommandInvocationsResponse) SetBody(v *DescribeCommandInvocationsResponseBody) *DescribeCommandInvocationsResponse {
	s.Body = v
	return s
}

func (s *DescribeCommandInvocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
