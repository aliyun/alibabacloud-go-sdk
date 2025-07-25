// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGtmMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGtmMonitorResponse
	GetStatusCode() *int32
	SetBody(v *AddGtmMonitorResponseBody) *AddGtmMonitorResponse
	GetBody() *AddGtmMonitorResponseBody
}

type AddGtmMonitorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGtmMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGtmMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGtmMonitorResponse) GetBody() *AddGtmMonitorResponseBody {
	return s.Body
}

func (s *AddGtmMonitorResponse) SetHeaders(v map[string]*string) *AddGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *AddGtmMonitorResponse) SetStatusCode(v int32) *AddGtmMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGtmMonitorResponse) SetBody(v *AddGtmMonitorResponseBody) *AddGtmMonitorResponse {
	s.Body = v
	return s
}

func (s *AddGtmMonitorResponse) Validate() error {
	return dara.Validate(s)
}
