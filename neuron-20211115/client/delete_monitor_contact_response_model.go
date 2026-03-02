// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorContactResponse
	GetStatusCode() *int32
}

type DeleteMonitorContactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMonitorContactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorContactResponse) SetHeaders(v map[string]*string) *DeleteMonitorContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorContactResponse) SetStatusCode(v int32) *DeleteMonitorContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorContactResponse) Validate() error {
	return dara.Validate(s)
}
