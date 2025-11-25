// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticQpsRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticQpsRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticQpsRecordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticQpsRecordResponseBody) *DescribeElasticQpsRecordResponse
	GetBody() *DescribeElasticQpsRecordResponseBody
}

type DescribeElasticQpsRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticQpsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticQpsRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticQpsRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticQpsRecordResponse) GetBody() *DescribeElasticQpsRecordResponseBody {
	return s.Body
}

func (s *DescribeElasticQpsRecordResponse) SetHeaders(v map[string]*string) *DescribeElasticQpsRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticQpsRecordResponse) SetStatusCode(v int32) *DescribeElasticQpsRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticQpsRecordResponse) SetBody(v *DescribeElasticQpsRecordResponseBody) *DescribeElasticQpsRecordResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticQpsRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
