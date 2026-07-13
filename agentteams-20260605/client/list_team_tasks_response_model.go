// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTeamTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTeamTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListTeamTasksResponseBody) *ListTeamTasksResponse
	GetBody() *ListTeamTasksResponseBody
}

type ListTeamTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTeamTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTeamTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTeamTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTeamTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTeamTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTeamTasksResponse) GetBody() *ListTeamTasksResponseBody {
	return s.Body
}

func (s *ListTeamTasksResponse) SetHeaders(v map[string]*string) *ListTeamTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTeamTasksResponse) SetStatusCode(v int32) *ListTeamTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTeamTasksResponse) SetBody(v *ListTeamTasksResponseBody) *ListTeamTasksResponse {
	s.Body = v
	return s
}

func (s *ListTeamTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
