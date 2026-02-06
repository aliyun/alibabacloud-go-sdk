// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentProfilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentProfilesResponseBody) *DeleteAgentProfilesResponse
	GetBody() *DeleteAgentProfilesResponseBody
}

type DeleteAgentProfilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentProfilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentProfilesResponse) GetBody() *DeleteAgentProfilesResponseBody {
	return s.Body
}

func (s *DeleteAgentProfilesResponse) SetHeaders(v map[string]*string) *DeleteAgentProfilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentProfilesResponse) SetStatusCode(v int32) *DeleteAgentProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentProfilesResponse) SetBody(v *DeleteAgentProfilesResponseBody) *DeleteAgentProfilesResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
