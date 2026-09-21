// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTeamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTeamResponse
	GetStatusCode() *int32
	SetBody(v *CreateTeamResponseBody) *CreateTeamResponse
	GetBody() *CreateTeamResponseBody
}

type CreateTeamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTeamResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponse) GoString() string {
	return s.String()
}

func (s *CreateTeamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTeamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTeamResponse) GetBody() *CreateTeamResponseBody {
	return s.Body
}

func (s *CreateTeamResponse) SetHeaders(v map[string]*string) *CreateTeamResponse {
	s.Headers = v
	return s
}

func (s *CreateTeamResponse) SetStatusCode(v int32) *CreateTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTeamResponse) SetBody(v *CreateTeamResponseBody) *CreateTeamResponse {
	s.Body = v
	return s
}

func (s *CreateTeamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
