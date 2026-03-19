// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse
	GetBody() *DescribeOpEntitiesResponseBody
}

type DescribeOpEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpEntitiesResponse) GetBody() *DescribeOpEntitiesResponseBody {
	return s.Body
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetStatusCode(v int32) *DescribeOpEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

func (s *DescribeOpEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
