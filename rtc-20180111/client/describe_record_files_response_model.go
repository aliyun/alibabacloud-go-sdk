// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordFilesResponseBody) *DescribeRecordFilesResponse
	GetBody() *DescribeRecordFilesResponseBody
}

type DescribeRecordFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordFilesResponse) GetBody() *DescribeRecordFilesResponseBody {
	return s.Body
}

func (s *DescribeRecordFilesResponse) SetHeaders(v map[string]*string) *DescribeRecordFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordFilesResponse) SetStatusCode(v int32) *DescribeRecordFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordFilesResponse) SetBody(v *DescribeRecordFilesResponseBody) *DescribeRecordFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
