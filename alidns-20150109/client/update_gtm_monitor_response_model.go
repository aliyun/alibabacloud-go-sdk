// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGtmMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGtmMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGtmMonitorResponseBody) *UpdateGtmMonitorResponse
	GetBody() *UpdateGtmMonitorResponseBody
}

type UpdateGtmMonitorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGtmMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGtmMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGtmMonitorResponse) GetBody() *UpdateGtmMonitorResponseBody {
	return s.Body
}

func (s *UpdateGtmMonitorResponse) SetHeaders(v map[string]*string) *UpdateGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmMonitorResponse) SetStatusCode(v int32) *UpdateGtmMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGtmMonitorResponse) SetBody(v *UpdateGtmMonitorResponseBody) *UpdateGtmMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateGtmMonitorResponse) Validate() error {
	return dara.Validate(s)
}
