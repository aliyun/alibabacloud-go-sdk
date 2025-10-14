// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMonitorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMonitorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMonitorGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMonitorGroupResponseBody) *ModifyMonitorGroupResponse
	GetBody() *ModifyMonitorGroupResponseBody
}

type ModifyMonitorGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMonitorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMonitorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMonitorGroupResponse) GetBody() *ModifyMonitorGroupResponseBody {
	return s.Body
}

func (s *ModifyMonitorGroupResponse) SetHeaders(v map[string]*string) *ModifyMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyMonitorGroupResponse) SetStatusCode(v int32) *ModifyMonitorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMonitorGroupResponse) SetBody(v *ModifyMonitorGroupResponseBody) *ModifyMonitorGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyMonitorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
