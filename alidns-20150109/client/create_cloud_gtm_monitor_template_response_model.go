// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmMonitorTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudGtmMonitorTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudGtmMonitorTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudGtmMonitorTemplateResponseBody) *CreateCloudGtmMonitorTemplateResponse
	GetBody() *CreateCloudGtmMonitorTemplateResponseBody
}

type CreateCloudGtmMonitorTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudGtmMonitorTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudGtmMonitorTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmMonitorTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmMonitorTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudGtmMonitorTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudGtmMonitorTemplateResponse) GetBody() *CreateCloudGtmMonitorTemplateResponseBody {
	return s.Body
}

func (s *CreateCloudGtmMonitorTemplateResponse) SetHeaders(v map[string]*string) *CreateCloudGtmMonitorTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudGtmMonitorTemplateResponse) SetStatusCode(v int32) *CreateCloudGtmMonitorTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateResponse) SetBody(v *CreateCloudGtmMonitorTemplateResponseBody) *CreateCloudGtmMonitorTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateCloudGtmMonitorTemplateResponse) Validate() error {
	return dara.Validate(s)
}
