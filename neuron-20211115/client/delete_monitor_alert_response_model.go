// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorAlertResponse
	GetStatusCode() *int32
}

type DeleteMonitorAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMonitorAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorAlertResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorAlertResponse) SetHeaders(v map[string]*string) *DeleteMonitorAlertResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorAlertResponse) SetStatusCode(v int32) *DeleteMonitorAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorAlertResponse) Validate() error {
	return dara.Validate(s)
}
