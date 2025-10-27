// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulTargetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVulTargetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVulTargetConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVulTargetConfigResponseBody) *ModifyVulTargetConfigResponse
	GetBody() *ModifyVulTargetConfigResponseBody
}

type ModifyVulTargetConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVulTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVulTargetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVulTargetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVulTargetConfigResponse) GetBody() *ModifyVulTargetConfigResponseBody {
	return s.Body
}

func (s *ModifyVulTargetConfigResponse) SetHeaders(v map[string]*string) *ModifyVulTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyVulTargetConfigResponse) SetStatusCode(v int32) *ModifyVulTargetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVulTargetConfigResponse) SetBody(v *ModifyVulTargetConfigResponseBody) *ModifyVulTargetConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyVulTargetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
