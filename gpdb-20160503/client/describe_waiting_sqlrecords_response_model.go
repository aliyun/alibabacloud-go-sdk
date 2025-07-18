// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWaitingSQLRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWaitingSQLRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWaitingSQLRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWaitingSQLRecordsResponseBody) *DescribeWaitingSQLRecordsResponse
	GetBody() *DescribeWaitingSQLRecordsResponseBody
}

type DescribeWaitingSQLRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWaitingSQLRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWaitingSQLRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWaitingSQLRecordsResponse) GetBody() *DescribeWaitingSQLRecordsResponseBody {
	return s.Body
}

func (s *DescribeWaitingSQLRecordsResponse) SetHeaders(v map[string]*string) *DescribeWaitingSQLRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) SetStatusCode(v int32) *DescribeWaitingSQLRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) SetBody(v *DescribeWaitingSQLRecordsResponseBody) *DescribeWaitingSQLRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponse) Validate() error {
	return dara.Validate(s)
}
