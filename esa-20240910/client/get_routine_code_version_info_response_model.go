// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineCodeVersionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineCodeVersionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineCodeVersionInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineCodeVersionInfoResponseBody) *GetRoutineCodeVersionInfoResponse
	GetBody() *GetRoutineCodeVersionInfoResponseBody
}

type GetRoutineCodeVersionInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineCodeVersionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineCodeVersionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineCodeVersionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineCodeVersionInfoResponse) GetBody() *GetRoutineCodeVersionInfoResponseBody {
	return s.Body
}

func (s *GetRoutineCodeVersionInfoResponse) SetHeaders(v map[string]*string) *GetRoutineCodeVersionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineCodeVersionInfoResponse) SetStatusCode(v int32) *GetRoutineCodeVersionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponse) SetBody(v *GetRoutineCodeVersionInfoResponseBody) *GetRoutineCodeVersionInfoResponse {
	s.Body = v
	return s
}

func (s *GetRoutineCodeVersionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
