// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAppointmentElectZookeeperLeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAppointmentElectZookeeperLeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAppointmentElectZookeeperLeaderResponse
	GetStatusCode() *int32
	SetBody(v *CancelAppointmentElectZookeeperLeaderResponseBody) *CancelAppointmentElectZookeeperLeaderResponse
	GetBody() *CancelAppointmentElectZookeeperLeaderResponseBody
}

type CancelAppointmentElectZookeeperLeaderResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAppointmentElectZookeeperLeaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAppointmentElectZookeeperLeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAppointmentElectZookeeperLeaderResponse) GoString() string {
	return s.String()
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) GetBody() *CancelAppointmentElectZookeeperLeaderResponseBody {
	return s.Body
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) SetHeaders(v map[string]*string) *CancelAppointmentElectZookeeperLeaderResponse {
	s.Headers = v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) SetStatusCode(v int32) *CancelAppointmentElectZookeeperLeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) SetBody(v *CancelAppointmentElectZookeeperLeaderResponseBody) *CancelAppointmentElectZookeeperLeaderResponse {
	s.Body = v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
