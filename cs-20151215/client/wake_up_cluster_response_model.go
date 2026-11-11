// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeUpClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WakeUpClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WakeUpClusterResponse
	GetStatusCode() *int32
	SetBody(v *WakeUpClusterResponseBody) *WakeUpClusterResponse
	GetBody() *WakeUpClusterResponseBody
}

type WakeUpClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WakeUpClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WakeUpClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s WakeUpClusterResponse) GoString() string {
	return s.String()
}

func (s *WakeUpClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WakeUpClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WakeUpClusterResponse) GetBody() *WakeUpClusterResponseBody {
	return s.Body
}

func (s *WakeUpClusterResponse) SetHeaders(v map[string]*string) *WakeUpClusterResponse {
	s.Headers = v
	return s
}

func (s *WakeUpClusterResponse) SetStatusCode(v int32) *WakeUpClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *WakeUpClusterResponse) SetBody(v *WakeUpClusterResponseBody) *WakeUpClusterResponse {
	s.Body = v
	return s
}

func (s *WakeUpClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
