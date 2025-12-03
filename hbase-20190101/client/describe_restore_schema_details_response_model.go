// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreSchemaDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreSchemaDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreSchemaDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreSchemaDetailsResponseBody) *DescribeRestoreSchemaDetailsResponse
	GetBody() *DescribeRestoreSchemaDetailsResponseBody
}

type DescribeRestoreSchemaDetailsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreSchemaDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreSchemaDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreSchemaDetailsResponse) GetBody() *DescribeRestoreSchemaDetailsResponseBody {
	return s.Body
}

func (s *DescribeRestoreSchemaDetailsResponse) SetHeaders(v map[string]*string) *DescribeRestoreSchemaDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponse) SetStatusCode(v int32) *DescribeRestoreSchemaDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponse) SetBody(v *DescribeRestoreSchemaDetailsResponseBody) *DescribeRestoreSchemaDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
