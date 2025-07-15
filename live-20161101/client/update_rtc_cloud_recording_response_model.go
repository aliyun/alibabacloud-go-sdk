// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcCloudRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRtcCloudRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRtcCloudRecordingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRtcCloudRecordingResponseBody) *UpdateRtcCloudRecordingResponse
	GetBody() *UpdateRtcCloudRecordingResponseBody
}

type UpdateRtcCloudRecordingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRtcCloudRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRtcCloudRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingResponse) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRtcCloudRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRtcCloudRecordingResponse) GetBody() *UpdateRtcCloudRecordingResponseBody {
	return s.Body
}

func (s *UpdateRtcCloudRecordingResponse) SetHeaders(v map[string]*string) *UpdateRtcCloudRecordingResponse {
	s.Headers = v
	return s
}

func (s *UpdateRtcCloudRecordingResponse) SetStatusCode(v int32) *UpdateRtcCloudRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRtcCloudRecordingResponse) SetBody(v *UpdateRtcCloudRecordingResponseBody) *UpdateRtcCloudRecordingResponse {
	s.Body = v
	return s
}

func (s *UpdateRtcCloudRecordingResponse) Validate() error {
	return dara.Validate(s)
}
