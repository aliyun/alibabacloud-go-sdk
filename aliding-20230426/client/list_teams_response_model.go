// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTeamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTeamsResponse
	GetStatusCode() *int32
	SetBody(v *ListTeamsResponseBody) *ListTeamsResponse
	GetBody() *ListTeamsResponseBody
}

type ListTeamsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTeamsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListTeamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTeamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTeamsResponse) GetBody() *ListTeamsResponseBody {
	return s.Body
}

func (s *ListTeamsResponse) SetHeaders(v map[string]*string) *ListTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListTeamsResponse) SetStatusCode(v int32) *ListTeamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTeamsResponse) SetBody(v *ListTeamsResponseBody) *ListTeamsResponse {
	s.Body = v
	return s
}

func (s *ListTeamsResponse) Validate() error {
	return dara.Validate(s)
}
