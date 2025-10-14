// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMonitorGroupResponseBody) *DeleteMonitorGroupResponse
	GetBody() *DeleteMonitorGroupResponseBody
}

type DeleteMonitorGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMonitorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorGroupResponse) GetBody() *DeleteMonitorGroupResponseBody {
	return s.Body
}

func (s *DeleteMonitorGroupResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupResponse) SetStatusCode(v int32) *DeleteMonitorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorGroupResponse) SetBody(v *DeleteMonitorGroupResponseBody) *DeleteMonitorGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteMonitorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
