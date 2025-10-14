// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridMonitorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridMonitorTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridMonitorTaskResponseBody) *DeleteHybridMonitorTaskResponse
	GetBody() *DeleteHybridMonitorTaskResponseBody
}

type DeleteHybridMonitorTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridMonitorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridMonitorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridMonitorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridMonitorTaskResponse) GetBody() *DeleteHybridMonitorTaskResponseBody {
	return s.Body
}

func (s *DeleteHybridMonitorTaskResponse) SetHeaders(v map[string]*string) *DeleteHybridMonitorTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridMonitorTaskResponse) SetStatusCode(v int32) *DeleteHybridMonitorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridMonitorTaskResponse) SetBody(v *DeleteHybridMonitorTaskResponseBody) *DeleteHybridMonitorTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridMonitorTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
