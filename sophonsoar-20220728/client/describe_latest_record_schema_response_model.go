// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLatestRecordSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLatestRecordSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLatestRecordSchemaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLatestRecordSchemaResponseBody) *DescribeLatestRecordSchemaResponse
	GetBody() *DescribeLatestRecordSchemaResponseBody
}

type DescribeLatestRecordSchemaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLatestRecordSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLatestRecordSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLatestRecordSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLatestRecordSchemaResponse) GetBody() *DescribeLatestRecordSchemaResponseBody {
	return s.Body
}

func (s *DescribeLatestRecordSchemaResponse) SetHeaders(v map[string]*string) *DescribeLatestRecordSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) SetStatusCode(v int32) *DescribeLatestRecordSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) SetBody(v *DescribeLatestRecordSchemaResponseBody) *DescribeLatestRecordSchemaResponse {
	s.Body = v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
