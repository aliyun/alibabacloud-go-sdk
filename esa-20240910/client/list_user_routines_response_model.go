// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRoutinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserRoutinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserRoutinesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserRoutinesResponseBody) *ListUserRoutinesResponse
	GetBody() *ListUserRoutinesResponseBody
}

type ListUserRoutinesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserRoutinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserRoutinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserRoutinesResponse) GoString() string {
	return s.String()
}

func (s *ListUserRoutinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserRoutinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserRoutinesResponse) GetBody() *ListUserRoutinesResponseBody {
	return s.Body
}

func (s *ListUserRoutinesResponse) SetHeaders(v map[string]*string) *ListUserRoutinesResponse {
	s.Headers = v
	return s
}

func (s *ListUserRoutinesResponse) SetStatusCode(v int32) *ListUserRoutinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserRoutinesResponse) SetBody(v *ListUserRoutinesResponseBody) *ListUserRoutinesResponse {
	s.Body = v
	return s
}

func (s *ListUserRoutinesResponse) Validate() error {
	return dara.Validate(s)
}
