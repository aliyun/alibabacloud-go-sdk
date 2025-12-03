// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHighDefinationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHighDefinationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHighDefinationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHighDefinationMonitorResponseBody) *ModifyHighDefinationMonitorResponse
	GetBody() *ModifyHighDefinationMonitorResponseBody
}

type ModifyHighDefinationMonitorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHighDefinationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHighDefinationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHighDefinationMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyHighDefinationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHighDefinationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHighDefinationMonitorResponse) GetBody() *ModifyHighDefinationMonitorResponseBody {
	return s.Body
}

func (s *ModifyHighDefinationMonitorResponse) SetHeaders(v map[string]*string) *ModifyHighDefinationMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyHighDefinationMonitorResponse) SetStatusCode(v int32) *ModifyHighDefinationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHighDefinationMonitorResponse) SetBody(v *ModifyHighDefinationMonitorResponseBody) *ModifyHighDefinationMonitorResponse {
	s.Body = v
	return s
}

func (s *ModifyHighDefinationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
