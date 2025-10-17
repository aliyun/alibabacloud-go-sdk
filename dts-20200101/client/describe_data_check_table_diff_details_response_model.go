// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckTableDiffDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataCheckTableDiffDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataCheckTableDiffDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataCheckTableDiffDetailsResponseBody) *DescribeDataCheckTableDiffDetailsResponse
	GetBody() *DescribeDataCheckTableDiffDetailsResponseBody
}

type DescribeDataCheckTableDiffDetailsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataCheckTableDiffDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataCheckTableDiffDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDiffDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDiffDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataCheckTableDiffDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataCheckTableDiffDetailsResponse) GetBody() *DescribeDataCheckTableDiffDetailsResponseBody {
	return s.Body
}

func (s *DescribeDataCheckTableDiffDetailsResponse) SetHeaders(v map[string]*string) *DescribeDataCheckTableDiffDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponse) SetStatusCode(v int32) *DescribeDataCheckTableDiffDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponse) SetBody(v *DescribeDataCheckTableDiffDetailsResponseBody) *DescribeDataCheckTableDiffDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
