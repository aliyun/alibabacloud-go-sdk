// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVulTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVulTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVulTargetResponseBody) *ModifyVulTargetResponse
	GetBody() *ModifyVulTargetResponseBody
}

type ModifyVulTargetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVulTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVulTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVulTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVulTargetResponse) GetBody() *ModifyVulTargetResponseBody {
	return s.Body
}

func (s *ModifyVulTargetResponse) SetHeaders(v map[string]*string) *ModifyVulTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyVulTargetResponse) SetStatusCode(v int32) *ModifyVulTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVulTargetResponse) SetBody(v *ModifyVulTargetResponseBody) *ModifyVulTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyVulTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
