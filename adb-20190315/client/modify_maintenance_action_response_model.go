// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaintenanceActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaintenanceActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaintenanceActionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaintenanceActionResponseBody) *ModifyMaintenanceActionResponse
	GetBody() *ModifyMaintenanceActionResponseBody
}

type ModifyMaintenanceActionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaintenanceActionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaintenanceActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaintenanceActionResponse) GetBody() *ModifyMaintenanceActionResponseBody {
	return s.Body
}

func (s *ModifyMaintenanceActionResponse) SetHeaders(v map[string]*string) *ModifyMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaintenanceActionResponse) SetStatusCode(v int32) *ModifyMaintenanceActionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaintenanceActionResponse) SetBody(v *ModifyMaintenanceActionResponseBody) *ModifyMaintenanceActionResponse {
	s.Body = v
	return s
}

func (s *ModifyMaintenanceActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
