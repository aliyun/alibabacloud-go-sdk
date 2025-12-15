// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyActiveOperationMaintainConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyActiveOperationMaintainConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyActiveOperationMaintainConfigResponseBody) *ModifyActiveOperationMaintainConfigResponse
	GetBody() *ModifyActiveOperationMaintainConfigResponseBody
}

type ModifyActiveOperationMaintainConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationMaintainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationMaintainConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyActiveOperationMaintainConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyActiveOperationMaintainConfigResponse) GetBody() *ModifyActiveOperationMaintainConfigResponseBody {
	return s.Body
}

func (s *ModifyActiveOperationMaintainConfigResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationMaintainConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationMaintainConfigResponse) SetStatusCode(v int32) *ModifyActiveOperationMaintainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigResponse) SetBody(v *ModifyActiveOperationMaintainConfigResponseBody) *ModifyActiveOperationMaintainConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyActiveOperationMaintainConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
