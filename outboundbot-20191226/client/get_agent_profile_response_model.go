// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentProfileResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentProfileResponseBody) *GetAgentProfileResponse
	GetBody() *GetAgentProfileResponseBody
}

type GetAgentProfileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileResponse) GoString() string {
	return s.String()
}

func (s *GetAgentProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentProfileResponse) GetBody() *GetAgentProfileResponseBody {
	return s.Body
}

func (s *GetAgentProfileResponse) SetHeaders(v map[string]*string) *GetAgentProfileResponse {
	s.Headers = v
	return s
}

func (s *GetAgentProfileResponse) SetStatusCode(v int32) *GetAgentProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentProfileResponse) SetBody(v *GetAgentProfileResponseBody) *GetAgentProfileResponse {
	s.Body = v
	return s
}

func (s *GetAgentProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
