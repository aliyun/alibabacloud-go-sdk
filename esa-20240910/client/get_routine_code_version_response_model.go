// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineCodeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineCodeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineCodeVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineCodeVersionResponseBody) *GetRoutineCodeVersionResponse
	GetBody() *GetRoutineCodeVersionResponseBody
}

type GetRoutineCodeVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineCodeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineCodeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineCodeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineCodeVersionResponse) GetBody() *GetRoutineCodeVersionResponseBody {
	return s.Body
}

func (s *GetRoutineCodeVersionResponse) SetHeaders(v map[string]*string) *GetRoutineCodeVersionResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineCodeVersionResponse) SetStatusCode(v int32) *GetRoutineCodeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineCodeVersionResponse) SetBody(v *GetRoutineCodeVersionResponseBody) *GetRoutineCodeVersionResponse {
	s.Body = v
	return s
}

func (s *GetRoutineCodeVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
