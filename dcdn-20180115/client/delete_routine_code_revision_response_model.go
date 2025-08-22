// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineCodeRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineCodeRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineCodeRevisionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineCodeRevisionResponseBody) *DeleteRoutineCodeRevisionResponse
	GetBody() *DeleteRoutineCodeRevisionResponseBody
}

type DeleteRoutineCodeRevisionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineCodeRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineCodeRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineCodeRevisionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineCodeRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineCodeRevisionResponse) GetBody() *DeleteRoutineCodeRevisionResponseBody {
	return s.Body
}

func (s *DeleteRoutineCodeRevisionResponse) SetHeaders(v map[string]*string) *DeleteRoutineCodeRevisionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineCodeRevisionResponse) SetStatusCode(v int32) *DeleteRoutineCodeRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineCodeRevisionResponse) SetBody(v *DeleteRoutineCodeRevisionResponseBody) *DeleteRoutineCodeRevisionResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineCodeRevisionResponse) Validate() error {
	return dara.Validate(s)
}
