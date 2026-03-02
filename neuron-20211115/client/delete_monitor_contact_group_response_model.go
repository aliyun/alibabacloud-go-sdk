// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorContactGroupResponse
	GetStatusCode() *int32
}

type DeleteMonitorContactGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMonitorContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorContactGroupResponse) SetHeaders(v map[string]*string) *DeleteMonitorContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorContactGroupResponse) SetStatusCode(v int32) *DeleteMonitorContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorContactGroupResponse) Validate() error {
	return dara.Validate(s)
}
