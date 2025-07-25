// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmMonitorTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudGtmMonitorTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudGtmMonitorTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudGtmMonitorTemplateResponseBody) *DeleteCloudGtmMonitorTemplateResponse
	GetBody() *DeleteCloudGtmMonitorTemplateResponseBody
}

type DeleteCloudGtmMonitorTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudGtmMonitorTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudGtmMonitorTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmMonitorTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmMonitorTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudGtmMonitorTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudGtmMonitorTemplateResponse) GetBody() *DeleteCloudGtmMonitorTemplateResponseBody {
	return s.Body
}

func (s *DeleteCloudGtmMonitorTemplateResponse) SetHeaders(v map[string]*string) *DeleteCloudGtmMonitorTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateResponse) SetStatusCode(v int32) *DeleteCloudGtmMonitorTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateResponse) SetBody(v *DeleteCloudGtmMonitorTemplateResponseBody) *DeleteCloudGtmMonitorTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateResponse) Validate() error {
	return dara.Validate(s)
}
