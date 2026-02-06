// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAgentProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAgentProfileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAgentProfileResponseBody) *ModifyAgentProfileResponse
	GetBody() *ModifyAgentProfileResponseBody
}

type ModifyAgentProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAgentProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAgentProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentProfileResponse) GoString() string {
	return s.String()
}

func (s *ModifyAgentProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAgentProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAgentProfileResponse) GetBody() *ModifyAgentProfileResponseBody {
	return s.Body
}

func (s *ModifyAgentProfileResponse) SetHeaders(v map[string]*string) *ModifyAgentProfileResponse {
	s.Headers = v
	return s
}

func (s *ModifyAgentProfileResponse) SetStatusCode(v int32) *ModifyAgentProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAgentProfileResponse) SetBody(v *ModifyAgentProfileResponseBody) *ModifyAgentProfileResponse {
	s.Body = v
	return s
}

func (s *ModifyAgentProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
