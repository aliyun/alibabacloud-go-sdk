// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInfoCheckExportRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInfoCheckExportRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInfoCheckExportRecordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInfoCheckExportRecordResponseBody) *DescribeInfoCheckExportRecordResponse
	GetBody() *DescribeInfoCheckExportRecordResponseBody
}

type DescribeInfoCheckExportRecordResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInfoCheckExportRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInfoCheckExportRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInfoCheckExportRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeInfoCheckExportRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInfoCheckExportRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInfoCheckExportRecordResponse) GetBody() *DescribeInfoCheckExportRecordResponseBody {
	return s.Body
}

func (s *DescribeInfoCheckExportRecordResponse) SetHeaders(v map[string]*string) *DescribeInfoCheckExportRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeInfoCheckExportRecordResponse) SetStatusCode(v int32) *DescribeInfoCheckExportRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInfoCheckExportRecordResponse) SetBody(v *DescribeInfoCheckExportRecordResponseBody) *DescribeInfoCheckExportRecordResponse {
	s.Body = v
	return s
}

func (s *DescribeInfoCheckExportRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
