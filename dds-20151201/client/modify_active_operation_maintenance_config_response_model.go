// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintenanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyActiveOperationMaintenanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyActiveOperationMaintenanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyActiveOperationMaintenanceConfigResponseBody) *ModifyActiveOperationMaintenanceConfigResponse
	GetBody() *ModifyActiveOperationMaintenanceConfigResponseBody
}

type ModifyActiveOperationMaintenanceConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationMaintenanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationMaintenanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintenanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) GetBody() *ModifyActiveOperationMaintenanceConfigResponseBody {
	return s.Body
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationMaintenanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) SetStatusCode(v int32) *ModifyActiveOperationMaintenanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) SetBody(v *ModifyActiveOperationMaintenanceConfigResponseBody) *ModifyActiveOperationMaintenanceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
