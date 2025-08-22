// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineResponseBody) *DeleteRoutineResponse
	GetBody() *DeleteRoutineResponseBody
}

type DeleteRoutineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineResponse) GetBody() *DeleteRoutineResponseBody {
	return s.Body
}

func (s *DeleteRoutineResponse) SetHeaders(v map[string]*string) *DeleteRoutineResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineResponse) SetStatusCode(v int32) *DeleteRoutineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineResponse) SetBody(v *DeleteRoutineResponseBody) *DeleteRoutineResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineResponse) Validate() error {
	return dara.Validate(s)
}
