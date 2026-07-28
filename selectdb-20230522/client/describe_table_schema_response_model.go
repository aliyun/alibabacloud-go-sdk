// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTableSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTableSchemaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTableSchemaResponseBody) *DescribeTableSchemaResponse
	GetBody() *DescribeTableSchemaResponseBody
}

type DescribeTableSchemaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTableSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTableSchemaResponse) GetBody() *DescribeTableSchemaResponseBody {
	return s.Body
}

func (s *DescribeTableSchemaResponse) SetHeaders(v map[string]*string) *DescribeTableSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableSchemaResponse) SetStatusCode(v int32) *DescribeTableSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableSchemaResponse) SetBody(v *DescribeTableSchemaResponseBody) *DescribeTableSchemaResponse {
	s.Body = v
	return s
}

func (s *DescribeTableSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
