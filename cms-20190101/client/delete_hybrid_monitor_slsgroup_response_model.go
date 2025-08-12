// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorSLSGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridMonitorSLSGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridMonitorSLSGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridMonitorSLSGroupResponseBody) *DeleteHybridMonitorSLSGroupResponse
	GetBody() *DeleteHybridMonitorSLSGroupResponseBody
}

type DeleteHybridMonitorSLSGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridMonitorSLSGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridMonitorSLSGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorSLSGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorSLSGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridMonitorSLSGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridMonitorSLSGroupResponse) GetBody() *DeleteHybridMonitorSLSGroupResponseBody {
	return s.Body
}

func (s *DeleteHybridMonitorSLSGroupResponse) SetHeaders(v map[string]*string) *DeleteHybridMonitorSLSGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponse) SetStatusCode(v int32) *DeleteHybridMonitorSLSGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponse) SetBody(v *DeleteHybridMonitorSLSGroupResponseBody) *DeleteHybridMonitorSLSGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponse) Validate() error {
	return dara.Validate(s)
}
