// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordDataResponseBody) *DescribeRecordDataResponse
	GetBody() *DescribeRecordDataResponseBody
}

type DescribeRecordDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordDataResponse) GetBody() *DescribeRecordDataResponseBody {
	return s.Body
}

func (s *DescribeRecordDataResponse) SetHeaders(v map[string]*string) *DescribeRecordDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordDataResponse) SetStatusCode(v int32) *DescribeRecordDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordDataResponse) SetBody(v *DescribeRecordDataResponseBody) *DescribeRecordDataResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
