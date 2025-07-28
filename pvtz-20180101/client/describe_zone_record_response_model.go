// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZoneRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZoneRecordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZoneRecordResponseBody) *DescribeZoneRecordResponse
	GetBody() *DescribeZoneRecordResponseBody
}

type DescribeZoneRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZoneRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZoneRecordResponse) GetBody() *DescribeZoneRecordResponseBody {
	return s.Body
}

func (s *DescribeZoneRecordResponse) SetHeaders(v map[string]*string) *DescribeZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneRecordResponse) SetStatusCode(v int32) *DescribeZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneRecordResponse) SetBody(v *DescribeZoneRecordResponseBody) *DescribeZoneRecordResponse {
	s.Body = v
	return s
}

func (s *DescribeZoneRecordResponse) Validate() error {
	return dara.Validate(s)
}
