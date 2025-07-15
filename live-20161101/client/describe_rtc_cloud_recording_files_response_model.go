// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcCloudRecordingFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcCloudRecordingFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcCloudRecordingFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcCloudRecordingFilesResponseBody) *DescribeRtcCloudRecordingFilesResponse
	GetBody() *DescribeRtcCloudRecordingFilesResponseBody
}

type DescribeRtcCloudRecordingFilesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcCloudRecordingFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcCloudRecordingFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudRecordingFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudRecordingFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcCloudRecordingFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcCloudRecordingFilesResponse) GetBody() *DescribeRtcCloudRecordingFilesResponseBody {
	return s.Body
}

func (s *DescribeRtcCloudRecordingFilesResponse) SetHeaders(v map[string]*string) *DescribeRtcCloudRecordingFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponse) SetStatusCode(v int32) *DescribeRtcCloudRecordingFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponse) SetBody(v *DescribeRtcCloudRecordingFilesResponseBody) *DescribeRtcCloudRecordingFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponse) Validate() error {
	return dara.Validate(s)
}
