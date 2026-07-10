// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTeamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTeamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTeamResponse
	GetStatusCode() *int32
	SetBody(v *GetTeamResponseBody) *GetTeamResponse
	GetBody() *GetTeamResponseBody
}

type GetTeamResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTeamResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponse) GoString() string {
	return s.String()
}

func (s *GetTeamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTeamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTeamResponse) GetBody() *GetTeamResponseBody {
	return s.Body
}

func (s *GetTeamResponse) SetHeaders(v map[string]*string) *GetTeamResponse {
	s.Headers = v
	return s
}

func (s *GetTeamResponse) SetStatusCode(v int32) *GetTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTeamResponse) SetBody(v *GetTeamResponseBody) *GetTeamResponse {
	s.Body = v
	return s
}

func (s *GetTeamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
