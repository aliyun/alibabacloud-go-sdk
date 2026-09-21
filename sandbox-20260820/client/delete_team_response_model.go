// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTeamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTeamResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTeamResponseBody) *DeleteTeamResponse
	GetBody() *DeleteTeamResponseBody
}

type DeleteTeamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTeamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTeamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamResponse) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTeamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTeamResponse) GetBody() *DeleteTeamResponseBody {
	return s.Body
}

func (s *DeleteTeamResponse) SetHeaders(v map[string]*string) *DeleteTeamResponse {
	s.Headers = v
	return s
}

func (s *DeleteTeamResponse) SetStatusCode(v int32) *DeleteTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTeamResponse) SetBody(v *DeleteTeamResponseBody) *DeleteTeamResponse {
	s.Body = v
	return s
}

func (s *DeleteTeamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
