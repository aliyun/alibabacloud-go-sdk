// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineResponseBody) *GetRoutineResponse
	GetBody() *GetRoutineResponseBody
}

type GetRoutineResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineResponse) GetBody() *GetRoutineResponseBody {
	return s.Body
}

func (s *GetRoutineResponse) SetHeaders(v map[string]*string) *GetRoutineResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineResponse) SetStatusCode(v int32) *GetRoutineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineResponse) SetBody(v *GetRoutineResponseBody) *GetRoutineResponse {
	s.Body = v
	return s
}

func (s *GetRoutineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
