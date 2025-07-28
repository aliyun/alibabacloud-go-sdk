// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZoneRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZoneRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZoneRecordsResponseBody) *DescribeZoneRecordsResponse
	GetBody() *DescribeZoneRecordsResponseBody
}

type DescribeZoneRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZoneRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZoneRecordsResponse) GetBody() *DescribeZoneRecordsResponseBody {
	return s.Body
}

func (s *DescribeZoneRecordsResponse) SetHeaders(v map[string]*string) *DescribeZoneRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneRecordsResponse) SetStatusCode(v int32) *DescribeZoneRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneRecordsResponse) SetBody(v *DescribeZoneRecordsResponseBody) *DescribeZoneRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeZoneRecordsResponse) Validate() error {
	return dara.Validate(s)
}
