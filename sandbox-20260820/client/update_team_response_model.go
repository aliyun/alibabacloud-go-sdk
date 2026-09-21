// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTeamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTeamResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTeamResponseBody) *UpdateTeamResponse
	GetBody() *UpdateTeamResponseBody
}

type UpdateTeamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTeamResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponse) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTeamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTeamResponse) GetBody() *UpdateTeamResponseBody {
	return s.Body
}

func (s *UpdateTeamResponse) SetHeaders(v map[string]*string) *UpdateTeamResponse {
	s.Headers = v
	return s
}

func (s *UpdateTeamResponse) SetStatusCode(v int32) *UpdateTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTeamResponse) SetBody(v *UpdateTeamResponseBody) *UpdateTeamResponse {
	s.Body = v
	return s
}

func (s *UpdateTeamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
