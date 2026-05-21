// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaintenanceWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMaintenanceWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMaintenanceWindowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMaintenanceWindowResponseBody) *UpdateMaintenanceWindowResponse
	GetBody() *UpdateMaintenanceWindowResponseBody
}

type UpdateMaintenanceWindowResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMaintenanceWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMaintenanceWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaintenanceWindowResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaintenanceWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMaintenanceWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMaintenanceWindowResponse) GetBody() *UpdateMaintenanceWindowResponseBody {
	return s.Body
}

func (s *UpdateMaintenanceWindowResponse) SetHeaders(v map[string]*string) *UpdateMaintenanceWindowResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaintenanceWindowResponse) SetStatusCode(v int32) *UpdateMaintenanceWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMaintenanceWindowResponse) SetBody(v *UpdateMaintenanceWindowResponseBody) *UpdateMaintenanceWindowResponse {
	s.Body = v
	return s
}

func (s *UpdateMaintenanceWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
