// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPendingMaintenanceActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPendingMaintenanceActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPendingMaintenanceActionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPendingMaintenanceActionResponseBody) *ModifyPendingMaintenanceActionResponse
	GetBody() *ModifyPendingMaintenanceActionResponseBody
}

type ModifyPendingMaintenanceActionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPendingMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPendingMaintenanceActionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPendingMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *ModifyPendingMaintenanceActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPendingMaintenanceActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPendingMaintenanceActionResponse) GetBody() *ModifyPendingMaintenanceActionResponseBody {
	return s.Body
}

func (s *ModifyPendingMaintenanceActionResponse) SetHeaders(v map[string]*string) *ModifyPendingMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *ModifyPendingMaintenanceActionResponse) SetStatusCode(v int32) *ModifyPendingMaintenanceActionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPendingMaintenanceActionResponse) SetBody(v *ModifyPendingMaintenanceActionResponseBody) *ModifyPendingMaintenanceActionResponse {
	s.Body = v
	return s
}

func (s *ModifyPendingMaintenanceActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
