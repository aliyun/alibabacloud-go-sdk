// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCloudMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageCloudMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageCloudMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageCloudMonitorConfigResponseBody) *GetMessageCloudMonitorConfigResponse
	GetBody() *GetMessageCloudMonitorConfigResponseBody
}

type GetMessageCloudMonitorConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageCloudMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageCloudMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageCloudMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageCloudMonitorConfigResponse) GetBody() *GetMessageCloudMonitorConfigResponseBody {
	return s.Body
}

func (s *GetMessageCloudMonitorConfigResponse) SetHeaders(v map[string]*string) *GetMessageCloudMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *GetMessageCloudMonitorConfigResponse) SetStatusCode(v int32) *GetMessageCloudMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageCloudMonitorConfigResponse) SetBody(v *GetMessageCloudMonitorConfigResponseBody) *GetMessageCloudMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *GetMessageCloudMonitorConfigResponse) Validate() error {
	return dara.Validate(s)
}
