// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourcesResponseBody) *DescribeDefenseResourcesResponse
	GetBody() *DescribeDefenseResourcesResponseBody
}

type DescribeDefenseResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourcesResponse) GetBody() *DescribeDefenseResourcesResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourcesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourcesResponse) SetStatusCode(v int32) *DescribeDefenseResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourcesResponse) SetBody(v *DescribeDefenseResourcesResponseBody) *DescribeDefenseResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
