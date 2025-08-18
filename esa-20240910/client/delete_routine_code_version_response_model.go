// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineCodeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineCodeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineCodeVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineCodeVersionResponseBody) *DeleteRoutineCodeVersionResponse
	GetBody() *DeleteRoutineCodeVersionResponseBody
}

type DeleteRoutineCodeVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineCodeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineCodeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineCodeVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineCodeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineCodeVersionResponse) GetBody() *DeleteRoutineCodeVersionResponseBody {
	return s.Body
}

func (s *DeleteRoutineCodeVersionResponse) SetHeaders(v map[string]*string) *DeleteRoutineCodeVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineCodeVersionResponse) SetStatusCode(v int32) *DeleteRoutineCodeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineCodeVersionResponse) SetBody(v *DeleteRoutineCodeVersionResponseBody) *DeleteRoutineCodeVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineCodeVersionResponse) Validate() error {
	return dara.Validate(s)
}
