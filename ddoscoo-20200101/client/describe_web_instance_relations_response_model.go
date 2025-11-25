// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebInstanceRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebInstanceRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebInstanceRelationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebInstanceRelationsResponseBody) *DescribeWebInstanceRelationsResponse
	GetBody() *DescribeWebInstanceRelationsResponseBody
}

type DescribeWebInstanceRelationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebInstanceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebInstanceRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebInstanceRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebInstanceRelationsResponse) GetBody() *DescribeWebInstanceRelationsResponseBody {
	return s.Body
}

func (s *DescribeWebInstanceRelationsResponse) SetHeaders(v map[string]*string) *DescribeWebInstanceRelationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebInstanceRelationsResponse) SetStatusCode(v int32) *DescribeWebInstanceRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponse) SetBody(v *DescribeWebInstanceRelationsResponseBody) *DescribeWebInstanceRelationsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebInstanceRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
