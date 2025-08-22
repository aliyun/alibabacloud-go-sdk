// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineConfEnvsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineConfEnvsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineConfEnvsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineConfEnvsResponseBody) *DeleteRoutineConfEnvsResponse
	GetBody() *DeleteRoutineConfEnvsResponseBody
}

type DeleteRoutineConfEnvsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineConfEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineConfEnvsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineConfEnvsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineConfEnvsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineConfEnvsResponse) GetBody() *DeleteRoutineConfEnvsResponseBody {
	return s.Body
}

func (s *DeleteRoutineConfEnvsResponse) SetHeaders(v map[string]*string) *DeleteRoutineConfEnvsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineConfEnvsResponse) SetStatusCode(v int32) *DeleteRoutineConfEnvsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineConfEnvsResponse) SetBody(v *DeleteRoutineConfEnvsResponseBody) *DeleteRoutineConfEnvsResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineConfEnvsResponse) Validate() error {
	return dara.Validate(s)
}
