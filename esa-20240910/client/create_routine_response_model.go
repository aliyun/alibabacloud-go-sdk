// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineResponseBody) *CreateRoutineResponse
	GetBody() *CreateRoutineResponseBody
}

type CreateRoutineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineResponse) GetBody() *CreateRoutineResponseBody {
	return s.Body
}

func (s *CreateRoutineResponse) SetHeaders(v map[string]*string) *CreateRoutineResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineResponse) SetStatusCode(v int32) *CreateRoutineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineResponse) SetBody(v *CreateRoutineResponseBody) *CreateRoutineResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineResponse) Validate() error {
	return dara.Validate(s)
}
