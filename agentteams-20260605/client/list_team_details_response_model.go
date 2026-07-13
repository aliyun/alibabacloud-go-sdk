// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTeamDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTeamDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListTeamDetailsResponseBody) *ListTeamDetailsResponse
	GetBody() *ListTeamDetailsResponseBody
}

type ListTeamDetailsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTeamDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTeamDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTeamDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListTeamDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTeamDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTeamDetailsResponse) GetBody() *ListTeamDetailsResponseBody {
	return s.Body
}

func (s *ListTeamDetailsResponse) SetHeaders(v map[string]*string) *ListTeamDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListTeamDetailsResponse) SetStatusCode(v int32) *ListTeamDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTeamDetailsResponse) SetBody(v *ListTeamDetailsResponseBody) *ListTeamDetailsResponse {
	s.Body = v
	return s
}

func (s *ListTeamDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
