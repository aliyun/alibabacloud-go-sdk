// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorSLSGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridMonitorSLSGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridMonitorSLSGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridMonitorSLSGroupResponseBody) *CreateHybridMonitorSLSGroupResponse
	GetBody() *CreateHybridMonitorSLSGroupResponseBody
}

type CreateHybridMonitorSLSGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridMonitorSLSGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridMonitorSLSGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorSLSGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorSLSGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridMonitorSLSGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridMonitorSLSGroupResponse) GetBody() *CreateHybridMonitorSLSGroupResponseBody {
	return s.Body
}

func (s *CreateHybridMonitorSLSGroupResponse) SetHeaders(v map[string]*string) *CreateHybridMonitorSLSGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponse) SetStatusCode(v int32) *CreateHybridMonitorSLSGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponse) SetBody(v *CreateHybridMonitorSLSGroupResponseBody) *CreateHybridMonitorSLSGroupResponse {
	s.Body = v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponse) Validate() error {
	return dara.Validate(s)
}
