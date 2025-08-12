// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridMonitorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridMonitorTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridMonitorTaskResponseBody) *CreateHybridMonitorTaskResponse
	GetBody() *CreateHybridMonitorTaskResponseBody
}

type CreateHybridMonitorTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridMonitorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridMonitorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridMonitorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridMonitorTaskResponse) GetBody() *CreateHybridMonitorTaskResponseBody {
	return s.Body
}

func (s *CreateHybridMonitorTaskResponse) SetHeaders(v map[string]*string) *CreateHybridMonitorTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridMonitorTaskResponse) SetStatusCode(v int32) *CreateHybridMonitorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridMonitorTaskResponse) SetBody(v *CreateHybridMonitorTaskResponseBody) *CreateHybridMonitorTaskResponse {
	s.Body = v
	return s
}

func (s *CreateHybridMonitorTaskResponse) Validate() error {
	return dara.Validate(s)
}
