// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeRecordFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeRecordFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeRecordFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeRecordFileUrlResponseBody) *ClinkDescribeRecordFileUrlResponse
	GetBody() *ClinkDescribeRecordFileUrlResponseBody
}

type ClinkDescribeRecordFileUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeRecordFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeRecordFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeRecordFileUrlResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeRecordFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeRecordFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeRecordFileUrlResponse) GetBody() *ClinkDescribeRecordFileUrlResponseBody {
	return s.Body
}

func (s *ClinkDescribeRecordFileUrlResponse) SetHeaders(v map[string]*string) *ClinkDescribeRecordFileUrlResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponse) SetStatusCode(v int32) *ClinkDescribeRecordFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponse) SetBody(v *ClinkDescribeRecordFileUrlResponseBody) *ClinkDescribeRecordFileUrlResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
