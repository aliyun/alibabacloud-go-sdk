// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordingFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppRecordingFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppRecordingFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppRecordingFilesResponseBody) *DescribeAppRecordingFilesResponse
	GetBody() *DescribeAppRecordingFilesResponseBody
}

type DescribeAppRecordingFilesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppRecordingFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppRecordingFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordingFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordingFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppRecordingFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppRecordingFilesResponse) GetBody() *DescribeAppRecordingFilesResponseBody {
	return s.Body
}

func (s *DescribeAppRecordingFilesResponse) SetHeaders(v map[string]*string) *DescribeAppRecordingFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppRecordingFilesResponse) SetStatusCode(v int32) *DescribeAppRecordingFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppRecordingFilesResponse) SetBody(v *DescribeAppRecordingFilesResponseBody) *DescribeAppRecordingFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppRecordingFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
