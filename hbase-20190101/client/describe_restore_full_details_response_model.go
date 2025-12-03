// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreFullDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreFullDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreFullDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreFullDetailsResponseBody) *DescribeRestoreFullDetailsResponse
	GetBody() *DescribeRestoreFullDetailsResponseBody
}

type DescribeRestoreFullDetailsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreFullDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreFullDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreFullDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreFullDetailsResponse) GetBody() *DescribeRestoreFullDetailsResponseBody {
	return s.Body
}

func (s *DescribeRestoreFullDetailsResponse) SetHeaders(v map[string]*string) *DescribeRestoreFullDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreFullDetailsResponse) SetStatusCode(v int32) *DescribeRestoreFullDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponse) SetBody(v *DescribeRestoreFullDetailsResponseBody) *DescribeRestoreFullDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreFullDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
