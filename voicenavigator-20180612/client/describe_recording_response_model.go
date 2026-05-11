// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordingResponseBody) *DescribeRecordingResponse
	GetBody() *DescribeRecordingResponseBody
}

type DescribeRecordingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordingResponse) GetBody() *DescribeRecordingResponseBody {
	return s.Body
}

func (s *DescribeRecordingResponse) SetHeaders(v map[string]*string) *DescribeRecordingResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordingResponse) SetStatusCode(v int32) *DescribeRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordingResponse) SetBody(v *DescribeRecordingResponseBody) *DescribeRecordingResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
