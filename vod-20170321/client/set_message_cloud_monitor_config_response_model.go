// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMessageCloudMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetMessageCloudMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetMessageCloudMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetMessageCloudMonitorConfigResponseBody) *SetMessageCloudMonitorConfigResponse
	GetBody() *SetMessageCloudMonitorConfigResponseBody
}

type SetMessageCloudMonitorConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMessageCloudMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMessageCloudMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetMessageCloudMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *SetMessageCloudMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetMessageCloudMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetMessageCloudMonitorConfigResponse) GetBody() *SetMessageCloudMonitorConfigResponseBody {
	return s.Body
}

func (s *SetMessageCloudMonitorConfigResponse) SetHeaders(v map[string]*string) *SetMessageCloudMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *SetMessageCloudMonitorConfigResponse) SetStatusCode(v int32) *SetMessageCloudMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMessageCloudMonitorConfigResponse) SetBody(v *SetMessageCloudMonitorConfigResponseBody) *SetMessageCloudMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *SetMessageCloudMonitorConfigResponse) Validate() error {
	return dara.Validate(s)
}
