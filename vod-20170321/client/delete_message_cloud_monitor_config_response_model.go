// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCloudMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageCloudMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageCloudMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessageCloudMonitorConfigResponseBody) *DeleteMessageCloudMonitorConfigResponse
	GetBody() *DeleteMessageCloudMonitorConfigResponseBody
}

type DeleteMessageCloudMonitorConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageCloudMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageCloudMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCloudMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageCloudMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageCloudMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageCloudMonitorConfigResponse) GetBody() *DeleteMessageCloudMonitorConfigResponseBody {
	return s.Body
}

func (s *DeleteMessageCloudMonitorConfigResponse) SetHeaders(v map[string]*string) *DeleteMessageCloudMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageCloudMonitorConfigResponse) SetStatusCode(v int32) *DeleteMessageCloudMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageCloudMonitorConfigResponse) SetBody(v *DeleteMessageCloudMonitorConfigResponseBody) *DeleteMessageCloudMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteMessageCloudMonitorConfigResponse) Validate() error {
	return dara.Validate(s)
}
